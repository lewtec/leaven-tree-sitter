package grammar_squirrel

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for squirrel (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_squirrel()))
}

func init() {
	grammar.Register("squirrel", Language())
}
