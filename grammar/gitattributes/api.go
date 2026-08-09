package grammar_gitattributes

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for gitattributes (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_gitattributes()))
}

func init() {
	grammar.Register("gitattributes", Language())
}
