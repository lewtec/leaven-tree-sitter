package grammar

import (
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var (
	registry = make(map[string]Language)
	mu       sync.RWMutex
)

// extensionToLanguage maps file extensions (no leading dot, lowercase) to
// registered language names. Only unambiguous mappings for languages this repo
// ships. Extensions that equal a registered name are handled by Get first.
var extensionToLanguage = map[string]string{
	// Existing public mappings (keep behavior)
	"js":   "javascript",
	"ts":   "typescript",
	"py":   "python",
	"rb":   "ruby",
	"rs":   "rust",
	"go":   "go",
	"c":    "c",
	"cpp":  "cpp",
	"h":    "c",
	"hpp":  "cpp",
	"lua":  "lua",
	"json": "json",

	// JavaScript / TypeScript family
	"mjs":    "javascript",
	"cjs":    "javascript",
	"jsx":    "javascript",
	"mts":    "typescript",
	"cts":    "typescript",
	"tsx":    "tsx",
	"vue":    "vue",
	"svelte": "svelte",
	"astro":  "astro",

	// Web markup / style
	"css":  "css",
	"html": "html",
	"htm":  "html",
	"xml":  "xml",
	"dtd":  "dtd",

	// C / C++ extras
	"cc":  "cpp",
	"cxx": "cpp",
	"hh":  "cpp",
	"hxx": "cpp",

	// Systems
	"zig":  "zig",
	"ha":   "hare",
	"cu":   "cuda",
	"cuh":  "cuda",
	"glsl": "glsl",
	"vert": "glsl",
	"frag": "glsl",

	// JVM / .NET / mobile
	"java":  "java",
	"kt":    "kotlin",
	"kts":   "kotlin",
	"scala": "scala",
	"cs":    "c_sharp",
	"dart":  "dart",

	// Scripting / shells
	"pyi":  "python",
	"sh":   "bash",
	"bash": "bash",
	"zsh":  "bash",
	"tcl":  "tcl",
	"pp":   "puppet",

	// Functional
	"hs":   "haskell",
	"lhs":  "haskell",
	"ml":   "ocaml",
	"mli":  "ocaml_interface",
	"ex":   "elixir",
	"exs":  "elixir",
	"jl":   "julia",
	"lisp": "commonlisp",
	"asd":  "commonlisp",

	// Data / config
	"yaml":      "yaml",
	"yml":       "yaml",
	"toml":      "toml",
	"csv":       "csv",
	"tsv":       "tsv",
	"psv":       "psv",
	"md":        "markdown",
	"markdown":  "markdown",
	"ron":       "ron",
	"po":        "po",
	"pot":       "po",
	"beancount": "beancount",
	"bean":      "beancount",

	// Infra
	"nix":   "nix",
	"tf":    "terraform",
	"hcl":   "hcl",
	"bicep": "bicep",
	"cmake": "cmake",
	"sol":   "solidity",
	"cairo": "cairo",

	// Build / other
	"mk":    "make",
	"make":  "make",
	"ino":   "arduino",
	"diff":  "diff",
	"patch": "diff",
	"scm":   "query", // tree-sitter query files
}

// Register associates name with lang in the process-wide registry.
// Re-registering the same name overwrites the previous entry silently.
// Names are case-sensitive map keys (exact match for Get) and should match
// the values used by GetByExtension's extension map. An empty name is stored
// as a normal key if passed.
func Register(name string, lang Language) {
	mu.Lock()
	defer mu.Unlock()
	registry[name] = lang
}

// Get returns the language registered under name, or false if none.
// Lookup is case-sensitive (exact key match on the name passed to Register).
func Get(name string) (Language, bool) {
	mu.RLock()
	defer mu.RUnlock()
	lang, ok := registry[name]
	return lang, ok
}

// List returns all registered grammar names in sorted order.
func List() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// GetByExtension tries to find a grammar based on file extension.
// Matching is case-insensitive on the extension.
func GetByExtension(filename string) (Language, bool) {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filename), "."))
	if ext == "" {
		return nil, false
	}

	// Direct match against registered language name
	if lang, ok := Get(ext); ok {
		return lang, ok
	}

	if name, ok := extensionToLanguage[ext]; ok {
		return Get(name)
	}

	return nil, false
}

// SupportedLanguages returns a comma-separated list of registered language
// names in sorted order, or "none" if the registry is empty.
func SupportedLanguages() string {
	langs := List()
	if len(langs) == 0 {
		return "none"
	}
	return strings.Join(langs, ", ")
}
