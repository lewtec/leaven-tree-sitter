package grammar_ungrammar

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for ungrammar (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ungrammar()))
}

func init() {
	grammar.Register("ungrammar", Language())
}
