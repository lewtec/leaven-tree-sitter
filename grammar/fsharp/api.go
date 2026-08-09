package grammar_fsharp

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for fsharp (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_fsharp()))
}

func init() {
	grammar.Register("fsharp", Language())
}
