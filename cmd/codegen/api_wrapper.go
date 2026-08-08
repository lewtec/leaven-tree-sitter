package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// ensureCoreAPI seeds outputDir/api.go from the hand-maintained grammar/api.go
// when missing. Returns a clear error if the live file cannot be found.
func ensureCoreAPI(outputDir string) error {
	coreAPIPath := filepath.Join(outputDir, "api.go")
	_, err := os.Stat(coreAPIPath)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}

	src, err := findLiveCoreAPI(coreAPIPath)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("read live core API %s: %w", src, err)
	}
	if err := os.WriteFile(coreAPIPath, data, 0644); err != nil {
		return err
	}
	return nil
}

// findLiveCoreAPI walks up from outputDir and cwd looking for grammar/api.go
// that is not the missing destination path.
func findLiveCoreAPI(dest string) (string, error) {
	destAbs, err := filepath.Abs(dest)
	if err != nil {
		destAbs = dest
	}

	seen := map[string]struct{}{}
	try := func(p string) (string, bool) {
		abs, err := filepath.Abs(p)
		if err != nil {
			abs = p
		}
		if abs == destAbs {
			return "", false
		}
		if _, ok := seen[abs]; ok {
			return "", false
		}
		seen[abs] = struct{}{}
		st, err := os.Stat(abs)
		if err != nil || st.IsDir() {
			return "", false
		}
		return abs, true
	}

	walk := func(start string) string {
		for d := start; d != "" && d != string(filepath.Separator) && d != "."; {
			if abs, ok := try(filepath.Join(d, "grammar", "api.go")); ok {
				return abs
			}
			parent := filepath.Dir(d)
			if parent == d {
				break
			}
			d = parent
		}
		return ""
	}

	starts := []string{}
	if abs, err := filepath.Abs(filepath.Dir(dest)); err == nil {
		starts = append(starts, abs)
	} else {
		starts = append(starts, filepath.Dir(dest))
	}
	if wd, err := os.Getwd(); err == nil {
		starts = append(starts, wd)
	}

	for _, start := range starts {
		if found := walk(start); found != "" {
			return found, nil
		}
	}

	return "", fmt.Errorf("core api.go missing at %s and hand-maintained grammar/api.go not found relative to outputDir or cwd; restore grammar/api.go before running codegen", dest)
}

// GenerateAPIWrapper writes the language package entry that registers with the
// hand-written grammar API. Leaven emits tree_sitter_<name>() with no TLS arg.
func GenerateAPIWrapper(outputDir, grammarName string) error {
	if err := ensureCoreAPI(outputDir); err != nil {
		return err
	}

	grammarAPI := fmt.Sprintf(`package grammar_%s

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for %s (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_%s()))
}

func init() {
	grammar.Register(%q, Language())
}
`, grammarName, grammarName, grammarName, grammarName)

	return os.WriteFile(filepath.Join(outputDir, grammarName, "api.go"), []byte(grammarAPI), 0644)
}

// GenerateAPIWrapperWithScanner is the same wrapper for grammars with scanners.
func GenerateAPIWrapperWithScanner(outputDir, grammarName string) error {
	return GenerateAPIWrapper(outputDir, grammarName)
}
