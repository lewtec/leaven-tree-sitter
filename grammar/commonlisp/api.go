package grammar_commonlisp

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for commonlisp (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_commonlisp()))
}

func init() {
	grammar.Register("commonlisp", Language())
}
