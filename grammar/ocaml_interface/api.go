package grammar_ocaml_interface

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for ocaml_interface (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ocaml_interface()))
}

func init() {
	grammar.Register("ocaml_interface", Language())
}
