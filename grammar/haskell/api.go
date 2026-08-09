package grammar_haskell

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for haskell (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_haskell()))
}

func init() {
	grammar.Register("haskell", Language())
}
