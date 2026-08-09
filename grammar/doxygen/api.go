package grammar_doxygen

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for doxygen (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_doxygen()))
}

func init() {
	grammar.Register("doxygen", Language())
}
