package grammar_bibtex

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for bibtex (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_bibtex()))
}

func init() {
	grammar.Register("bibtex", Language())
}
