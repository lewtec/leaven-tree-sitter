package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// LeavenTranspiler turns tree-sitter C into Go via clang -emit-llvm + go tool leaven.
// Output is raw leaven Go (no post-processing). This is the only codegen backend.
type LeavenTranspiler struct {
	TreeSitterPath string
	KeepTemp       bool
	// Clang is the clang binary (default: CC or clang on PATH; pin 14 via mise).
	Clang string
}

// TranspileCore emits grammar/core.go from tree-sitter lib/src/lib.c.
func (t *LeavenTranspiler) TranspileCore(outputDir string) error {
	libC := filepath.Join(t.TreeSitterPath, "lib/src/lib.c")
	if _, err := os.Stat(libC); err != nil {
		return fmt.Errorf("tree-sitter lib.c: %w", err)
	}
	tmpDir, err := os.MkdirTemp("", "leaven-core-*")
	if err != nil {
		return err
	}
	if !t.KeepTemp {
		defer os.RemoveAll(tmpDir)
	} else {
		slog.Info("keeping temp dir", "path", tmpDir)
	}

	llPath := filepath.Join(tmpDir, "core.ll")
	includes := []string{
		filepath.Join(t.TreeSitterPath, "lib/include"),
		filepath.Join(t.TreeSitterPath, "lib/src"),
	}
	if err := t.emitLLVM(libC, llPath, includes); err != nil {
		return err
	}
	if err := t.runLeaven(llPath, "grammar"); err != nil {
		return err
	}
	srcGo := strings.TrimSuffix(llPath, ".ll") + ".go"
	data, err := os.ReadFile(srcGo)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	dest := filepath.Join(outputDir, "core.go")
	if err := os.WriteFile(dest, data, 0644); err != nil {
		return err
	}
	slog.Info("wrote leaven core", "path", dest, "bytes", len(data))
	return nil
}

// TranspileGrammar writes grammar/<name>/grammar.go from unit root grammarPath.
func (t *LeavenTranspiler) TranspileGrammar(grammarPath, grammarName, grammarRoot string) error {
	if grammarName == "" {
		grammarName = normalizeGrammarName(grammarPath)
	}
	parserC := filepath.Join(grammarPath, "src", "parser.c")
	if _, err := os.Stat(parserC); err != nil {
		return fmt.Errorf("parser.c: %w", err)
	}
	scannerC := filepath.Join(grammarPath, "src", "scanner.c")
	hasScanner := false
	if _, err := os.Stat(scannerC); err == nil {
		hasScanner = true
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("scanner.c: %w", err)
	}

	tmpDir, err := os.MkdirTemp("", "leaven-"+grammarName+"-*")
	if err != nil {
		return err
	}
	if !t.KeepTemp {
		defer os.RemoveAll(tmpDir)
	} else {
		slog.Info("keeping temp dir", "path", tmpDir)
	}

	srcC := parserC
	if hasScanner {
		combined := filepath.Join(tmpDir, "combined.c")
		// scanner.c first (typed definitions), then parser.c with external-scanner
		// prototypes stripped so void*/const prototypes do not conflict with
		// scanner Payload*/char* definitions (idris, purescript, …).
		if err := combineScannerAndParser(scannerC, parserC, combined); err != nil {
			return err
		}
		srcC = combined
	}

	llPath := filepath.Join(tmpDir, grammarName+".ll")
	includes := []string{
		filepath.Join(grammarPath, "src"),
		grammarPath,
		filepath.Dir(grammarPath),
		filepath.Join(t.TreeSitterPath, "lib/include"),
		filepath.Join(t.TreeSitterPath, "lib/src"),
	}
	if err := t.emitLLVM(srcC, llPath, includes); err != nil {
		return err
	}
	pkg := "grammar_" + grammarName
	if err := t.runLeaven(llPath, pkg); err != nil {
		return err
	}
	srcGo := strings.TrimSuffix(llPath, ".ll") + ".go"
	data, err := os.ReadFile(srcGo)
	if err != nil {
		return err
	}

	outDir := filepath.Join(grammarRoot, grammarName)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}
	dest := filepath.Join(outDir, "grammar.go")
	if err := os.WriteFile(dest, data, 0644); err != nil {
		return err
	}
	if hasScanner {
		if err := GenerateAPIWrapperWithScanner(grammarRoot, grammarName); err != nil {
			return err
		}
	} else {
		if err := GenerateAPIWrapper(grammarRoot, grammarName); err != nil {
			return err
		}
	}
	slog.Info("wrote leaven grammar", "grammar", grammarName, "path", dest, "bytes", len(data))
	return nil
}

func (t *LeavenTranspiler) clangBin() string {
	if t.Clang != "" {
		return t.Clang
	}
	if cc := os.Getenv("CC"); cc != "" {
		return cc
	}
	return "clang"
}

func (t *LeavenTranspiler) emitLLVM(srcC, llPath string, includes []string) error {
	args := []string{
		"-S", "-emit-llvm",
		"-fno-discard-value-names",
		"-std=gnu11",
		"-O0",
		"-o", llPath,
	}
	for _, inc := range includes {
		abs, err := filepath.Abs(inc)
		if err != nil {
			return err
		}
		args = append(args, "-I", abs)
	}
	srcAbs, err := filepath.Abs(srcC)
	if err != nil {
		return err
	}
	args = append(args, srcAbs)

	cmd := exec.Command(t.clangBin(), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	slog.Info("clang emit-llvm", "cc", t.clangBin(), "src", srcC)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("clang -emit-llvm: %w", err)
	}
	return nil
}

// runLeaven invokes `go tool leaven` (module tool; no global install).
// pkg is the generated Go package name (-package).
func (t *LeavenTranspiler) runLeaven(llPath, pkg string) error {
	abs, err := filepath.Abs(llPath)
	if err != nil {
		return err
	}
	if pkg == "" {
		pkg = "main"
	}
	cmd := exec.Command("go", "tool", "leaven", "-package", pkg, abs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if root, err := moduleRoot(); err == nil {
		cmd.Dir = root
	}
	slog.Info("go tool leaven", "ll", abs, "package", pkg)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go tool leaven: %w", err)
	}
	return nil
}

func moduleRoot() (string, error) {
	cmd := exec.Command("go", "env", "GOMOD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	mod := strings.TrimSpace(string(out))
	if mod == "" || mod == os.DevNull {
		return "", fmt.Errorf("not in a module")
	}
	return filepath.Dir(mod), nil
}

// combineScannerAndParser writes scanner.c then parser.c into one TU.
// Parser forward decls of tree_sitter_*_external_scanner_* are stripped so
// they do not conflict with scanner definitions (Payload* vs void*, char* vs
// const char*).
func combineScannerAndParser(scannerPath, parserPath, output string) (err error) {
	scanner, err := os.ReadFile(scannerPath)
	if err != nil {
		return err
	}
	parser, err := os.ReadFile(parserPath)
	if err != nil {
		return err
	}
	parser = stripExternalScannerPrototypes(parser)

	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); err == nil {
			err = cerr
		}
	}()
	if _, err = out.Write(scanner); err != nil {
		return err
	}
	if _, err = out.WriteString("\n\n"); err != nil {
		return err
	}
	if _, err = out.Write(parser); err != nil {
		return err
	}
	if _, err = out.WriteString("\n"); err != nil {
		return err
	}
	return nil
}

// externalScannerProto matches a single-line C prototype for the external
// scanner hooks that generated parser.c emits (definition ends with '{', not ';').
var externalScannerProto = regexp.MustCompile(
	`(?m)^[ \t]*(?:extern[ \t]+)?` +
		`[a-zA-Z_][\w\s\*]*\b` +
		`tree_sitter_\w+_external_scanner_(?:create|destroy|scan|serialize|deserialize)` +
		`\s*\([^;]*\)\s*;[ \t]*\r?\n?`,
)

// stripExternalScannerPrototypes removes parser.c forward decls of external
// scanner API symbols. Leaves language-table references and any definitions.
func stripExternalScannerPrototypes(src []byte) []byte {
	return externalScannerProto.ReplaceAll(src, nil)
}
