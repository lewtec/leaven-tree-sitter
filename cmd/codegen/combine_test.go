package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStripExternalScannerPrototypes(t *testing.T) {
	in := `
void *tree_sitter_foo_external_scanner_create(void);
void tree_sitter_foo_external_scanner_destroy(void *);
bool tree_sitter_foo_external_scanner_scan(void *, TSLexer *, const bool *);
unsigned tree_sitter_foo_external_scanner_serialize(void *, char *);
void tree_sitter_foo_external_scanner_deserialize(void *, const char *, unsigned);

static const TSLanguage *tree_sitter_foo(void) {
  static TSLanguage language = {
    .external_scanner = {
      tree_sitter_foo_external_scanner_create,
      tree_sitter_foo_external_scanner_destroy,
    },
  };
  return &language;
}
`
	out := string(stripExternalScannerPrototypes([]byte(in)))
	for _, bad := range []string{
		"external_scanner_create(void);",
		"external_scanner_destroy(void *);",
		"external_scanner_scan(void *",
		"external_scanner_serialize(void *",
		"external_scanner_deserialize(void *",
	} {
		if strings.Contains(out, bad) {
			t.Errorf("prototype still present: %q\n%s", bad, out)
		}
	}
	if !strings.Contains(out, "tree_sitter_foo_external_scanner_create,") {
		t.Fatalf("language table reference stripped:\n%s", out)
	}
}

func TestCombineScannerAndParserNoConflict(t *testing.T) {
	dir := t.TempDir()
	scanner := filepath.Join(dir, "scanner.c")
	parser := filepath.Join(dir, "parser.c")
	combined := filepath.Join(dir, "combined.c")
	// Minimal repro shape (idris-style Payload* vs void*).
	mustWrite(t, scanner, `
typedef struct TSLexer TSLexer;
typedef struct { int depth; } Payload;
void *tree_sitter_foo_external_scanner_create(void) { return (void *)0; }
_Bool tree_sitter_foo_external_scanner_scan(Payload *p, TSLexer *l, const _Bool *s) {
  (void)p;(void)l;(void)s; return 0;
}
unsigned tree_sitter_foo_external_scanner_serialize(Payload *p, char *b) {
  (void)p;(void)b; return 0;
}
void tree_sitter_foo_external_scanner_deserialize(Payload *p, char *b, unsigned n) {
  (void)p;(void)b;(void)n;
}
void tree_sitter_foo_external_scanner_destroy(Payload *p) { (void)p; }
`)
	mustWrite(t, parser, `
typedef struct TSLexer TSLexer;
void *tree_sitter_foo_external_scanner_create(void);
void tree_sitter_foo_external_scanner_destroy(void *);
_Bool tree_sitter_foo_external_scanner_scan(void *, TSLexer *, const _Bool *);
unsigned tree_sitter_foo_external_scanner_serialize(void *, char *);
void tree_sitter_foo_external_scanner_deserialize(void *, const char *, unsigned);
void *hooks[] = {
  (void *)tree_sitter_foo_external_scanner_create,
  (void *)tree_sitter_foo_external_scanner_destroy,
};
`)
	if err := combineScannerAndParser(scanner, parser, combined); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(combined)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "destroy(void *);") {
		t.Fatalf("parser prototype not stripped:\n%s", data)
	}
}

func mustWrite(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}
