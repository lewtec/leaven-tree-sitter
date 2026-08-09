package grammar_fsharp_signature

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for fsharp_signature (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_fsharp_signature()))
}

func init() {
	grammar.Register("fsharp_signature", Language())
}
