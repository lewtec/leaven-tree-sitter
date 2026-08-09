package grammar_gdscript

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for gdscript (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_gdscript()))
}

func init() {
	grammar.Register("gdscript", Language())
}
