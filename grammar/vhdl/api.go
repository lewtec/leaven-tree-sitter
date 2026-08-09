package grammar_vhdl

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for vhdl (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_vhdl()))
}

func init() {
	grammar.Register("vhdl", Language())
}
