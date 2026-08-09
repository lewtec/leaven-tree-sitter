package grammar_core_schema

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for core_schema (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_core_schema()))
}

func init() {
	grammar.Register("core_schema", Language())
}
