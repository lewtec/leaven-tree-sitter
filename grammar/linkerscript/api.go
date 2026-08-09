package grammar_linkerscript

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for linkerscript (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_linkerscript()))
}

func init() {
	grammar.Register("linkerscript", Language())
}
