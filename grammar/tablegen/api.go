package grammar_tablegen

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for tablegen (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_tablegen()))
}

func init() {
	grammar.Register("tablegen", Language())
}
