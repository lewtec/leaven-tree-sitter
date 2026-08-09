package grammar_uxntal

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for uxntal (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_uxntal()))
}

func init() {
	grammar.Register("uxntal", Language())
}
