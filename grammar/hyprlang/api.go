package grammar_hyprlang

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for hyprlang (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_hyprlang()))
}

func init() {
	grammar.Register("hyprlang", Language())
}
