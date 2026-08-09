package grammar_wgsl_bevy

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for wgsl_bevy (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_wgsl_bevy()))
}

func init() {
	grammar.Register("wgsl_bevy", Language())
}
