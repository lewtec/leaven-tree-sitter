package grammar_gitcommit

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for gitcommit (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_gitcommit()))
}

func init() {
	grammar.Register("gitcommit", Language())
}
