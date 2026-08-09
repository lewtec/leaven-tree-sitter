package grammar_tsv

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for tsv (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_tsv()))
}

func init() {
	grammar.Register("tsv", Language())
}
