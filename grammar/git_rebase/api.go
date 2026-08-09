package grammar_git_rebase

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for git_rebase (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_git_rebase()))
}

func init() {
	grammar.Register("git_rebase", Language())
}
