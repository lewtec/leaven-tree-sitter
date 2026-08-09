package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
)

const (
	rootModulePath    = "github.com/lewtec/leaven-tree-sitter"
	grammarModulePath = rootModulePath + "/grammar"
	moduleGoVersion   = "1.25.0"
	localPseudoVer    = "v0.0.0"
	// leavenModuleVer must match the root go.mod require.
	leavenModulePath = "github.com/lewtec/leaven"
	leavenModuleVer  = "v0.0.0-20260809213144-953c116a246c"
)

// ensureGrammarModules writes grammar/go.mod, grammar/<lang>/go.mod (with local
// replace directives), go.work, and wires root go.mod require+replace entries.
func ensureGrammarModules(outputDir string) error {
	abs, err := filepath.Abs(outputDir)
	if err != nil {
		return err
	}
	grammarDir := filepath.Join(abs, "grammar")
	if err := os.MkdirAll(grammarDir, 0755); err != nil {
		return err
	}
	if err := writeCoreGoMod(grammarDir); err != nil {
		return err
	}

	langs, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}
	for _, lang := range langs {
		if err := writeLangGoMod(grammarDir, lang); err != nil {
			return err
		}
	}
	if err := writeGoWork(abs, langs); err != nil {
		return err
	}
	if err := updateRootGoMod(abs, langs); err != nil {
		return fmt.Errorf("update root go.mod: %w", err)
	}
	slog.Info("wrote grammar modules", "core", 1, "langs", len(langs), "go.work", true)
	return nil
}

// tidyGrammarModules runs go work sync and go mod tidy for root, core, and each
// language module so go.sum files exist.
func tidyGrammarModules(outputDir string) error {
	abs, err := filepath.Abs(outputDir)
	if err != nil {
		return err
	}
	grammarDir := filepath.Join(abs, "grammar")
	langs, err := listGrammarLangs(grammarDir)
	if err != nil {
		return err
	}

	if err := runGo(abs, "work", "sync"); err != nil {
		return fmt.Errorf("go work sync: %w", err)
	}
	if err := runGo(abs, "mod", "tidy"); err != nil {
		return fmt.Errorf("go mod tidy (root): %w", err)
	}
	if err := runGo(grammarDir, "mod", "tidy"); err != nil {
		return fmt.Errorf("go mod tidy (grammar): %w", err)
	}
	for _, lang := range langs {
		dir := filepath.Join(grammarDir, lang)
		if err := runGo(dir, "mod", "tidy"); err != nil {
			return fmt.Errorf("go mod tidy (grammar/%s): %w", lang, err)
		}
		slog.Info("tidied grammar module", "lang", lang)
	}
	return nil
}

func runGo(dir string, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	return cmd.Run()
}

// listGrammarLangs lists language dirs that have leaven grammar.go.
func listGrammarLangs(grammarDir string) ([]string, error) {
	entries, err := os.ReadDir(grammarDir)
	if err != nil {
		return nil, err
	}
	var langs []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		if _, err := os.Stat(filepath.Join(grammarDir, name, "grammar.go")); err != nil {
			continue
		}
		langs = append(langs, name)
	}
	sort.Strings(langs)
	return langs, nil
}

func writeCoreGoMod(grammarDir string) error {
	content := fmt.Sprintf(`module %s

go %s

require %s %s
`, grammarModulePath, moduleGoVersion, leavenModulePath, leavenModuleVer)
	return os.WriteFile(filepath.Join(grammarDir, "go.mod"), []byte(content), 0644)
}

func writeLangGoMod(grammarDir, lang string) error {
	content := fmt.Sprintf(`module %s/%s

go %s

require (
	%s %s
	%s %s
)

replace %s => ../
`, grammarModulePath, lang, moduleGoVersion,
		grammarModulePath, localPseudoVer,
		leavenModulePath, leavenModuleVer,
		grammarModulePath)
	return os.WriteFile(filepath.Join(grammarDir, lang, "go.mod"), []byte(content), 0644)
}

func writeGoWork(outputDir string, langs []string) error {
	var b strings.Builder
	b.WriteString("go " + moduleGoVersion + "\n\n")
	b.WriteString("use (\n")
	b.WriteString("\t.\n")
	b.WriteString("\t./grammar\n")
	for _, lang := range langs {
		fmt.Fprintf(&b, "\t./grammar/%s\n", lang)
	}
	b.WriteString(")\n")
	return os.WriteFile(filepath.Join(outputDir, "go.work"), []byte(b.String()), 0644)
}

// updateRootGoMod adds require + replace for the core grammar module and each
// language module so cmd/parse can depend on them from the monorepo root.
func updateRootGoMod(outputDir string, langs []string) error {
	path := filepath.Join(outputDir, "go.mod")
	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		data = []byte("module " + rootModulePath + "\n\ngo " + moduleGoVersion + "\n")
	}
	f, err := modfile.Parse(path, data, nil)
	if err != nil {
		return err
	}

	type modRef struct {
		path string
		dir  string
	}
	mods := []modRef{{grammarModulePath, "./grammar"}}
	for _, lang := range langs {
		mods = append(mods, modRef{grammarModulePath + "/" + lang, "./grammar/" + lang})
	}
	// Drop every local grammar require/replace we manage, including stale
	// languages removed since the last run (prefix match, not only current set).
	isManaged := func(path string) bool {
		if path == grammarModulePath {
			return true
		}
		return strings.HasPrefix(path, grammarModulePath+"/")
	}
	for {
		var dropPath string
		for _, r := range f.Require {
			if isManaged(r.Mod.Path) {
				dropPath = r.Mod.Path
				break
			}
		}
		if dropPath == "" {
			break
		}
		if err := f.DropRequire(dropPath); err != nil {
			return fmt.Errorf("drop require %s: %w", dropPath, err)
		}
	}
	for {
		var dropPath string
		for _, r := range f.Replace {
			if isManaged(r.Old.Path) {
				dropPath = r.Old.Path
				break
			}
		}
		if dropPath == "" {
			break
		}
		if err := f.DropReplace(dropPath, ""); err != nil {
			return fmt.Errorf("drop replace %s: %w", dropPath, err)
		}
	}

	for _, m := range mods {
		f.AddNewRequire(m.path, localPseudoVer, false)
		if err := f.AddReplace(m.path, "", m.dir, ""); err != nil {
			return err
		}
	}

	f.Cleanup()
	out, err := f.Format()
	if err != nil {
		return err
	}
	return os.WriteFile(path, out, 0644)
}
