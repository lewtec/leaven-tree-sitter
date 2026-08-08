package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// GrammarUnit is one transpilation target: a directory containing src/parser.c
// plus the Go package / registry name for that language.
type GrammarUnit struct {
	// Name is the normalized language id (Go package suffix, registry key).
	Name string
	// Path is the grammar unit root (parent of src/), containing src/parser.c.
	Path string
	// ParserC is the absolute or relative path to src/parser.c.
	ParserC string
	// Priority is lower for preferred units; unwanted paths (schema, dialects, …) score higher.
	Priority int
}

var (
	// tree_sitter_<name>(void) language entry in generated parser.c
	langEntryRe = regexp.MustCompile(`(?m)(?:TS_PUBLIC\s+)?const\s+TSLanguage\s*\*\s*tree_sitter_(\w+)\s*\(\s*void\s*\)`)
)

// discoverGrammarUnits finds every …/src/parser.c under third-party/tree-sitter-*,
// assigns a language name (prefer C entry symbol tree_sitter_X, else unit folder),
// sorts unwanted locations last, and keeps the first unit per normalized name.
// No size or denylist filters — oversized/unsupported units fail at transpile time.
func discoverGrammarUnits(thirdPartyGlob string) ([]GrammarUnit, error) {
	repos, err := filepath.Glob(thirdPartyGlob)
	if err != nil {
		return nil, err
	}

	var candidates []GrammarUnit
	for _, repo := range repos {
		info, err := os.Stat(repo)
		if err != nil || !info.IsDir() {
			continue
		}
		err = filepath.WalkDir(repo, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				// Skip heavy / irrelevant trees
				base := d.Name()
				switch base {
				case "node_modules", ".git", "target", "build", "prebuilds", ".build":
					return filepath.SkipDir
				}
				return nil
			}
			if d.Name() != "parser.c" {
				return nil
			}
			// Require …/src/parser.c (drops examples/parser.c and similar)
			if filepath.Base(filepath.Dir(path)) != "src" {
				return nil
			}
			unitPath := filepath.Dir(filepath.Dir(path)) // parent of src/
			name := languageNameForParser(path, unitPath)
			if name == "" {
				return nil
			}
			candidates = append(candidates, GrammarUnit{
				Name:     name,
				Path:     unitPath,
				ParserC:  path,
				Priority: grammarPriority(path, repo),
			})
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	// Preferred (low priority) first; stable by path for determinism
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].Priority != candidates[j].Priority {
			return candidates[i].Priority < candidates[j].Priority
		}
		if candidates[i].Name != candidates[j].Name {
			return candidates[i].Name < candidates[j].Name
		}
		return candidates[i].Path < candidates[j].Path
	})

	// First match wins per language name
	seen := make(map[string]struct{}, len(candidates))
	out := make([]GrammarUnit, 0, len(candidates))
	for _, u := range candidates {
		if _, ok := seen[u.Name]; ok {
			continue
		}
		seen[u.Name] = struct{}{}
		out = append(out, u)
	}

	// Stable output order by name for logs / summaries
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

// languageNameForParser prefers the tree_sitter_<id> symbol from parser.c;
// falls back to the unit directory name (tree-sitter- prefix stripped).
// Missing parser.c is a quiet fallback; other read errors are logged then fall back.
func languageNameForParser(parserC, unitPath string) string {
	data, err := os.ReadFile(parserC)
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Warn("failed to read parser.c for language name", "path", parserC, "error", err)
		}
		return normalizeGrammarName(filepath.Base(unitPath))
	}
	if m := langEntryRe.FindSubmatch(data); len(m) == 2 {
		return string(m[1])
	}
	return normalizeGrammarName(filepath.Base(unitPath))
}

// normalizeGrammarName turns a folder or JSON-ish name into a Go-safe id.
func normalizeGrammarName(name string) string {
	name = filepath.Base(name)
	const prefix = "tree-sitter-"
	if strings.HasPrefix(name, prefix) {
		name = name[len(prefix):]
	}
	return strings.ReplaceAll(name, "-", "_")
}

// grammarPriority ranks candidates; lower is preferred. Unwanted locations sort last
// so first-wins on name keeps product grammars over schema/dialect/example copies.
func grammarPriority(parserC, repoRoot string) int {
	rel, err := filepath.Rel(repoRoot, parserC)
	if err != nil {
		rel = parserC
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	score := 0

	for _, p := range parts {
		switch p {
		case "examples", "test", "tests", "corpus":
			score += 100
		case "schema":
			score += 50
		case "dialects", "dialect":
			score += 40
		}
	}

	// Prefer shallower paths under the submodule (classic <repo>/src/parser.c)
	// over monorepo subdirs; monorepos still beat schema/dialects via the switch above.
	// Depth in path components excluding parser.c itself.
	if depth := len(parts) - 1; depth > 0 {
		score += depth
	}

	return score
}
