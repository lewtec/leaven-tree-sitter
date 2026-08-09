package grammar_bitbake

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for bitbake (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_bitbake()))
}

func init() {
	grammar.Register("bitbake", Language())
}
