package grammar_qmldir

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for qmldir (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_qmldir()))
}

func init() {
	grammar.Register("qmldir", Language())
}
