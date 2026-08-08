package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/lewtec/leaven-tree-sitter/grammar"
)

var structPaddingPattern = regexp.MustCompile(`&struct \{ _ \[`)

const goldenSuffix = ".golden.json"

// Layout:
//
//	testdata/<language>/<file.ext>
//	testdata/<language>/<file.ext>.golden.json
//
// Golden is grammar.ParseOutput JSON (language, file, root ParseNode tree).
// Test plan: grammar.List() languages × source files under testdata/<lang>/.
// Missing testdata skips with a GitHub Actions warning (does not fail CI).
// fixtureSources seeds sources when UPDATE_GOLDENS=1.

func TestGeneratedCoreHasNoStructPadding(t *testing.T) {
	root := repoRoot(t)
	corePath := filepath.Join(root, "grammar", "core.go")
	content, err := os.ReadFile(corePath)
	if err != nil {
		t.Fatalf("read core %s: %v", corePath, err)
	}
	if structPaddingPattern.Match(content) {
		t.Fatalf("found struct padding pattern in %s", corePath)
	}
}

func TestLanguageFixtures(t *testing.T) {
	root := repoRoot(t)
	testdataRoot := filepath.Join(root, "testdata")
	update := os.Getenv("UPDATE_GOLDENS") == "1"

	langs := grammar.List()
	if len(langs) == 0 {
		t.Skip("no languages registered; cook PR adds grammar packages")
	}

	// fixtureSources entries for packages that do not build yet are ignored
	// until those grammars compile under leaven and are blank-imported again.
	for language := range fixtureSources {
		if _, ok := grammar.Get(language); !ok {
			t.Logf("fixtureSources %q skipped (not registered / package does not build)", language)
		}
	}

	fixtureCount := 0
	var parseErrors []string
	var missingFixtures []string

	for _, language := range langs {
		t.Run(language, func(t *testing.T) {
			lang, ok := grammar.Get(language)
			if !ok {
				t.Fatalf("language %q not registered", language)
			}

			langDir := filepath.Join(testdataRoot, language)

			// Materialise curated source when refreshing goldens.
			if update {
				if snip, ok := fixtureSources[language]; ok {
					if err := os.MkdirAll(langDir, 0o755); err != nil {
						t.Fatalf("mkdir %s: %v", langDir, err)
					}
					srcPath := filepath.Join(langDir, snip.File)
					if err := os.WriteFile(srcPath, []byte(snip.Source), 0o644); err != nil {
						t.Fatalf("write source %s: %v", srcPath, err)
					}
				}
			}

			sources, err := listFixtureSources(langDir)
			if err != nil && !os.IsNotExist(err) {
				t.Fatalf("list fixtures in %s: %v", langDir, err)
			}
			if len(sources) == 0 {
				missingFixtures = append(missingFixtures, language)
				skipMissingFixture(t, language)
				return
			}

			for _, fileName := range sources {
				fileName := fileName
				srcPath := filepath.Join(langDir, fileName)
				goldenPath := srcPath + goldenSuffix
				fixtureCount++

				t.Run(fileName, func(t *testing.T) {
					source, err := os.ReadFile(srcPath)
					if err != nil {
						t.Fatalf("read fixture: %v", err)
					}

					parser := grammar.NewParser()
					if !parser.SetLanguage(lang) {
						t.Fatalf("SetLanguage %q failed", language)
					}
					tree := parser.ParseBytes(source)
					rootNode := tree.RootNode()
					if rootNode.IsNull() {
						t.Fatal("null root")
					}
					if rootNode.HasError() {
						msg := language + "/" + fileName + ":\n" + rootNode.PrintTree()
						if update {
							// Still record failure when updating — goldens for broken trees are useless.
							parseErrors = append(parseErrors, msg)
						}
						t.Fatalf("parse errors:\n%s", rootNode.PrintTree())
					}

					parseRoot := grammar.BuildParseNode(rootNode, source, "")
					if parseRoot == nil {
						t.Fatal("BuildParseNode returned nil")
					}

					gotOut := grammar.ParseOutput{
						Language: language,
						File:     fileName,
						Root:     parseRoot,
					}
					gotJSON, err := json.MarshalIndent(gotOut, "", "  ")
					if err != nil {
						t.Fatalf("marshal golden candidate: %v", err)
					}
					gotJSON = append(gotJSON, '\n')

					want, err := os.ReadFile(goldenPath)
					if err != nil {
						if os.IsNotExist(err) && update {
							if err := os.WriteFile(goldenPath, gotJSON, 0o644); err != nil {
								t.Fatalf("write golden %s: %v", goldenPath, err)
							}
							t.Logf("wrote %s", goldenPath)
							return
						}
						t.Fatalf("missing golden %s (set UPDATE_GOLDENS=1 to create)", goldenPath)
					}

					if d := cmp.Diff(string(want), string(gotJSON)); d != "" {
						if update {
							if err := os.WriteFile(goldenPath, gotJSON, 0o644); err != nil {
								t.Fatalf("update golden %s: %v", goldenPath, err)
							}
							t.Logf("updated %s", goldenPath)
							return
						}
						t.Fatalf("golden mismatch for %s (-want +got):\n%s\n(set UPDATE_GOLDENS=1 to rewrite)", goldenPath, d)
					}
				})
			}
		})
	}

	// Emit GHA annotations from the parent test. Prefer GHA_ANNOTATIONS_FILE
	// (CI re-cats after go test) because go test discards stdout on success.
	if len(missingFixtures) > 0 {
		emitMissingFixtureHints(missingFixtures)
		t.Logf("skipped %d languages with no testdata fixtures (see GitHub Actions warnings)", len(missingFixtures))
	}

	if fixtureCount == 0 {
		t.Fatal("no language fixtures exercised")
	}
	if len(parseErrors) > 0 && update {
		t.Fatalf("%d fixtures still parse with errors — fix fixture sources:\n%s",
			len(parseErrors), strings.Join(parseErrors, "\n---\n"))
	}
}

// listFixtureSources returns non-golden regular files under langDir.
func listFixtureSources(langDir string) ([]string, error) {
	entries, err := os.ReadDir(langDir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, goldenSuffix) {
			continue
		}
		// provenance / notes (not source)
		if name == "SOURCE.txt" || strings.HasPrefix(name, ".") {
			continue
		}
		files = append(files, name)
	}
	slices.Sort(files)
	return files, nil
}

func skipMissingFixture(t *testing.T, language string) {
	t.Helper()
	t.Skipf("no parse fixture under testdata/%s", language)
}

// emitMissingFixtureHints surfaces missing fixtures to GitHub Actions without failing the job.
//
// go test discards the test binary's stdout on success unless -v is set, so workflow
// commands written only to os.Stdout never become Annotations. Prefer, in order:
//  1. GHA_ANNOTATIONS_FILE — CI cats this after go test (always reaches the step log)
//  2. GITHUB_STEP_SUMMARY — job summary markdown
//  3. stdout — visible with go test -v
func emitMissingFixtureHints(languages []string) {
	var ann strings.Builder
	var summary strings.Builder
	summary.WriteString("## Missing parse fixtures\n\n")
	summary.WriteString("Registered languages with no files under `testdata/<lang>/` (skipped, not failed):\n\n")
	for _, language := range languages {
		msg := fmt.Sprintf("no parse fixture under testdata/%s — add a source file + golden (optional: fixtureSources for UPDATE_GOLDENS=1)", language)
		file := "testdata/" + language
		line := fmt.Sprintf("::warning title=%s,file=%s::%s\n",
			ghaEscape("Missing parse fixture"),
			ghaEscape(file),
			ghaEscape(msg),
		)
		ann.WriteString(line)
		fmt.Fprintf(&summary, "- `%s`\n", language)
		fmt.Fprint(os.Stdout, line)
	}
	if path := os.Getenv("GHA_ANNOTATIONS_FILE"); path != "" {
		_ = os.WriteFile(path, []byte(ann.String()), 0o644)
	}
	if path := os.Getenv("GITHUB_STEP_SUMMARY"); path != "" {
		if f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644); err == nil {
			_, _ = f.WriteString(summary.String())
			_ = f.Close()
		}
	}
}

func ghaEscape(s string) string {
	s = strings.ReplaceAll(s, "%", "%25")
	s = strings.ReplaceAll(s, "\r", "%0D")
	s = strings.ReplaceAll(s, "\n", "%0A")
	return s
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("could not locate repository root from %s", dir)
		}
		dir = parent
	}
}
