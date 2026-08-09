package grammar_json_schema

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for json_schema (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_json_schema()))
}

func init() {
	grammar.Register("json_schema", Language())
}
