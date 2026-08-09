package grammar_firrtl

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for firrtl (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_firrtl()))
}

func init() {
	grammar.Register("firrtl", Language())
}
