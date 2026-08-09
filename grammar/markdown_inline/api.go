package grammar_markdown_inline

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for markdown_inline (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_markdown_inline()))
}

func init() {
	grammar.Register("markdown_inline", Language())
}
