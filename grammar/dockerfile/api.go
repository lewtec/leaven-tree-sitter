package grammar_dockerfile

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for dockerfile (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_dockerfile()))
}

func init() {
	grammar.Register("dockerfile", Language())
}
