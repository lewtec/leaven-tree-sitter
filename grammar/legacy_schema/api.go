package grammar_legacy_schema

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for legacy_schema (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_legacy_schema()))
}

func init() {
	grammar.Register("legacy_schema", Language())
}
