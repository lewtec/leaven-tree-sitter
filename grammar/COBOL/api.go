package grammar_COBOL

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for COBOL (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_COBOL()))
}

func init() {
	grammar.Register("COBOL", Language())
}
