package grammar

import (
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// stubLang is a typed nil Language; the registry only needs a present entry.
var stubLang Language

func TestList(t *testing.T) {
	Register("zz_registry_test_b", stubLang)
	Register("aa_registry_test_a", stubLang)
	Register("mm_registry_test_m", stubLang)

	names := List()
	if !sort.StringsAreSorted(names) {
		t.Fatalf("List() not sorted: %v", names)
	}
	if got, want := strings.Join(List(), ","), strings.Join(names, ","); got != want {
		t.Fatalf("List() not stable across calls: %q vs %q", got, want)
	}

	var idxA, idxM, idxB = -1, -1, -1
	for i, n := range names {
		switch n {
		case "aa_registry_test_a":
			idxA = i
		case "mm_registry_test_m":
			idxM = i
		case "zz_registry_test_b":
			idxB = i
		}
	}
	if idxA < 0 || idxM < 0 || idxB < 0 {
		t.Fatalf("expected test langs in List: %v", names)
	}
	if !(idxA < idxM && idxM < idxB) {
		t.Fatalf("test langs not in sorted order: a=%d m=%d b=%d", idxA, idxM, idxB)
	}
}

func TestSupportedLanguages(t *testing.T) {
	Register("yy_supported_test", stubLang)
	Register("bb_supported_test", stubLang)

	s := SupportedLanguages()
	want := strings.Join(List(), ", ")
	if s != want {
		t.Fatalf("SupportedLanguages() = %q, want %q", s, want)
	}
	if !strings.Contains(s, "bb_supported_test") || !strings.Contains(s, "yy_supported_test") {
		t.Fatalf("SupportedLanguages() missing test names: %q", s)
	}
	if i, j := strings.Index(s, "bb_supported_test"), strings.Index(s, "yy_supported_test"); i > j {
		t.Fatalf("SupportedLanguages not sorted: %q", s)
	}
}

// registerStub temporarily installs stubLang under name and restores the prior
// registry entry (or deletes the key) in t.Cleanup so tests do not clobber
// production language registrations from blank-imports.
func registerStub(t *testing.T, name string) {
	t.Helper()
	mu.Lock()
	prev, had := registry[name]
	registry[name] = stubLang
	mu.Unlock()
	t.Cleanup(func() {
		mu.Lock()
		defer mu.Unlock()
		if had {
			registry[name] = prev
		} else {
			delete(registry, name)
		}
	})
}

func TestGetByExtension(t *testing.T) {
	// Names referenced by mapping / direct-match extensions. Use registerStub
	// so real grammar entries are restored after the test.
	for _, name := range []string{
		"javascript", "typescript", "tsx", "python", "ruby", "rust", "go",
		"c", "cpp", "lua", "json", "c_sharp", "kotlin", "yaml", "markdown",
		"bash", "toml", "html", "css", "xml", "nix", "terraform", "hcl",
		"vue", "svelte", "zig", "julia",
	} {
		registerStub(t, name)
	}

	cases := []struct {
		file string
		want string // expected language name; empty means not found
	}{
		// Existing mappings
		{"app.js", "javascript"},
		{"main.ts", "typescript"},
		{"script.py", "python"},
		{"lib.rb", "ruby"},
		{"main.rs", "rust"},
		{"main.go", "go"},
		{"file.c", "c"},
		{"file.cpp", "cpp"},
		{"file.h", "c"},
		{"file.hpp", "cpp"},
		{"mod.lua", "lua"},
		{"data.json", "json"},

		// Expanded
		{"App.cs", "c_sharp"},
		{"Main.kt", "kotlin"},
		{"build.kts", "kotlin"},
		{"config.yaml", "yaml"},
		{"config.yml", "yaml"},
		{"README.md", "markdown"},
		{"run.sh", "bash"},
		{"Cargo.toml", "toml"},
		{"index.html", "html"},
		{"style.css", "css"},
		{"doc.xml", "xml"},
		{"flake.nix", "nix"},
		{"main.tf", "terraform"},
		{"module.hcl", "hcl"},
		{"App.vue", "vue"},
		{"Widget.svelte", "svelte"},
		{"main.zig", "zig"},
		{"plot.jl", "julia"},
		{"lib.jsx", "javascript"},
		{"app.tsx", "tsx"},
		{"util.cc", "cpp"},
		{"module.mjs", "javascript"},

		// Case-insensitive extension
		{"Foo.JS", "javascript"},
		{"Bar.MD", "markdown"},

		// Misses
		{"file.unknownext", ""},
		{"Makefile", ""},
		{"", ""},
	}

	for _, tc := range cases {
		_, ok := GetByExtension(tc.file)
		if tc.want == "" {
			if ok {
				t.Errorf("GetByExtension(%q) ok=true, want false", tc.file)
			}
			continue
		}
		if !ok {
			t.Errorf("GetByExtension(%q) ok=false, want language %q", tc.file, tc.want)
			continue
		}
		// Verify resolution path matches expected name (stubs share identity).
		if resolved := resolveExtension(tc.file); resolved != tc.want {
			t.Errorf("GetByExtension(%q) resolves to %q, want %q", tc.file, resolved, tc.want)
		}
	}
}

// resolveExtension mirrors GetByExtension name resolution for assertions.
func resolveExtension(filename string) string {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	if ext == "" {
		return ""
	}
	if _, ok := Get(ext); ok {
		return ext
	}
	if name, ok := extensionToLanguage[ext]; ok {
		if _, ok := Get(name); ok {
			return name
		}
	}
	return ""
}
