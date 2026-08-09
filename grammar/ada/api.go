package grammar_ada

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for ada (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ada()))
}

func init() {
	grammar.Register("ada", Language())
}
