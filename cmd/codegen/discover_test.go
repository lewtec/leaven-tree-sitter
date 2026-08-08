package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverGrammarUnits_integration(t *testing.T) {
	root := findRepoRoot(t)
	if _, err := os.Stat(filepath.Join(root, "third-party")); err != nil {
		t.Skip("third-party not placed (JIT C); run mise run grammars:sync")
	}
	old, _ := os.Getwd()
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(old) })

	units, err := discoverGrammarUnits("third-party/tree-sitter-*")
	if err != nil {
		t.Fatal(err)
	}
	if len(units) == 0 {
		t.Skip("no grammar units under third-party/")
	}
	if len(units) < 20 {
		// Phase 1 places php only. Full inventory checks need a full sync.
		t.Logf("discovered %d unit(s); skip monorepo inventory asserts", len(units))
		return
	}

	byName := map[string]GrammarUnit{}
	for _, u := range units {
		if _, dup := byName[u.Name]; dup {
			t.Fatalf("duplicate name after first-wins: %s", u.Name)
		}
		byName[u.Name] = u
	}

	// Monorepo languages should appear with C symbol names
	for _, want := range []string{"markdown", "markdown_inline", "typescript", "tsx"} {
		if _, ok := byName[want]; !ok {
			t.Errorf("missing monorepo grammar %q", want)
		}
	}

	// Product json from C symbol tree_sitter_json (not folder "json" under yaml/schema)
	jsonUnit, ok := byName["json"]
	if !ok {
		t.Fatal("missing json")
	}
	if containsPathPart(jsonUnit.ParserC, "schema") {
		t.Errorf("json must be tree-sitter-json product grammar, got %s", jsonUnit.ParserC)
	}
	if !containsPathPart(jsonUnit.ParserC, "tree-sitter-json") {
		t.Errorf("json unit path unexpected: %s", jsonUnit.ParserC)
	}

	// Schema parsers use distinct C entry names (tree_sitter_json_schema), so they do not
	// collide with product json — first-wins is on the C-derived id, not the folder alone.
	if _, ok := byName["json_schema"]; !ok {
		t.Error("expected json_schema from yaml schema/json (C symbol tree_sitter_json_schema)")
	}
	if _, ok := byName["core_schema"]; !ok {
		t.Error("expected core_schema from yaml schema/core")
	}
	if _, ok := byName["legacy_schema"]; !ok {
		t.Error("expected legacy_schema from yaml schema/legacy")
	}
	if _, ok := byName["terraform"]; !ok {
		t.Error("expected terraform from hcl dialects (declared upstream language)")
	}

	// No examples/parser.c
	for _, u := range units {
		if containsPathPart(u.ParserC, "examples") {
			t.Errorf("examples parser should be excluded: %s", u.ParserC)
		}
	}

	// No denylist / size gate: oversized grammars (e.g. fsharp) are scheduled.
	if _, ok := byName["fsharp"]; !ok {
		t.Error("expected fsharp to be discovered (no skip filters)")
	}
	if _, ok := byName["fsharp_signature"]; !ok {
		t.Error("expected fsharp_signature to be discovered (no skip filters)")
	}
}

func TestNormalizeGrammarName(t *testing.T) {
	cases := map[string]string{
		"tree-sitter-bash": "bash",
		"c-sharp":          "c_sharp",
		"typescript":       "typescript",
		"markdown_inline":  "markdown_inline",
	}
	for in, want := range cases {
		if got := normalizeGrammarName(in); got != want {
			t.Errorf("normalizeGrammarName(%q)=%q want %q", in, got, want)
		}
	}
}

func TestLanguageNameForParser(t *testing.T) {
	dir := t.TempDir()
	unitPath := filepath.Join(dir, "tree-sitter-example")
	if err := os.MkdirAll(filepath.Join(unitPath, "src"), 0o755); err != nil {
		t.Fatal(err)
	}
	parserC := filepath.Join(unitPath, "src", "parser.c")

	// C entry symbol wins over folder name.
	body := "const TSLanguage *tree_sitter_example_lang(void) {\n\treturn NULL;\n}\n"
	if err := os.WriteFile(parserC, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := languageNameForParser(parserC, unitPath); got != "example_lang" {
		t.Errorf("symbol path: got %q want example_lang", got)
	}

	// Missing file: quiet fallback to unit folder.
	missing := filepath.Join(unitPath, "src", "nope.c")
	if got := languageNameForParser(missing, unitPath); got != "example" {
		t.Errorf("IsNotExist fallback: got %q want example", got)
	}

	// Unreadable parser.c: still falls back (slog.Warn is side effect).
	if err := os.WriteFile(parserC, []byte(body), 0o000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chmod(parserC, 0o644) })
	// Root may ignore mode bits; only assert when read truly fails.
	if _, err := os.ReadFile(parserC); err != nil {
		if got := languageNameForParser(parserC, unitPath); got != "example" {
			t.Errorf("read-error fallback: got %q want example", got)
		}
	}
}

func containsPathPart(path, part string) bool {
	for _, p := range splitPath(path) {
		if p == part {
			return true
		}
	}
	return false
}

func splitPath(path string) []string {
	var out []string
	for path != "" && path != "." && path != string(filepath.Separator) {
		out = append(out, filepath.Base(path))
		parent := filepath.Dir(path)
		if parent == path {
			break
		}
		path = parent
	}
	return out
}

func findRepoRoot(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir := wd
	for i := 0; i < 6; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "workspaced.cue")); err == nil {
				return dir
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatalf("could not find repo root from %s", wd)
	return ""
}
