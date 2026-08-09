package grammar_ocaml

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for ocaml (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ocaml()))
}

func init() {
	grammar.Register("ocaml", Language())
}
