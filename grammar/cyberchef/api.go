package grammar_cyberchef

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for cyberchef (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_cyberchef()))
}

func init() {
	grammar.Register("cyberchef", Language())
}
