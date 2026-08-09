package grammar_ssh_config

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for ssh_config (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_ssh_config()))
}

func init() {
	grammar.Register("ssh_config", Language())
}
