package grammar_kconfig

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for kconfig (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_kconfig()))
}

func init() {
	grammar.Register("kconfig", Language())
}
