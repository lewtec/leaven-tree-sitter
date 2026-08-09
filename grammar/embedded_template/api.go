package grammar_embedded_template

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for embedded_template (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_embedded_template()))
}

func init() {
	grammar.Register("embedded_template", Language())
}
