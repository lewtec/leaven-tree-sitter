package grammar_vimdoc

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for vimdoc (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_vimdoc()))
}

func init() {
	grammar.Register("vimdoc", Language())
}
