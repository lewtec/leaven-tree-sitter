package grammar_pymanifest

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for pymanifest (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_pymanifest()))
}

func init() {
	grammar.Register("pymanifest", Language())
}
