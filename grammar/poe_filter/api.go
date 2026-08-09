package grammar_poe_filter

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for poe_filter (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_poe_filter()))
}

func init() {
	grammar.Register("poe_filter", Language())
}
