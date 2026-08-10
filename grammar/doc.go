// Package grammar binds tree-sitter as pure Go (leaven only: no CGO, no modernc).
//
// # Generated files
//
// core.go and per-language grammar/<lang>/grammar.go are raw leaven output
// (mise run codegen: clang 22 → LLVM IR → leaven). Do not edit by hand.
//
// # Hand-written API
//
// api.go, query.go, parse.go, and registry.go wrap leaven symbols with a
// thread-safe Go API: Parser/Query serialize native work under a mutex; Tree
// wraps the native *TSTree and Node methods call into it.
// LiveParseReady(lang) probes once (sync.Once) that parse completes.
//
// # Registry
//
// Blank-import a language package so its init Registers the language:
//
//	import _ "github.com/lewtec/leaven-tree-sitter/grammar/json"
//
// # Line index
//
// LineIndex converts tree-sitter byte offsets to 1-based lines and 0-based
// columns (byte offset within the line).
package grammar
