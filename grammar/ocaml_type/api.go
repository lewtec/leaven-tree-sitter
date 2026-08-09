package grammar_ocaml_type

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for ocaml_type (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ocaml_type()))
}

func init() {
	grammar.Register("ocaml_type", Language())
}
