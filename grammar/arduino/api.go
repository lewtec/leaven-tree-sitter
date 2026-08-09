package grammar_arduino

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for arduino (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_arduino()))
}

func init() {
	grammar.Register("arduino", Language())
}
