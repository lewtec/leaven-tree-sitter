package grammar_godot_resource

import (
	"unsafe"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

// Language returns the TSLanguage for godot_resource (leaven-generated).
func Language() grammar.Language {
	return (*grammar.TSLanguage)(unsafe.Pointer(tree_sitter_godot_resource()))
}

func init() {
	grammar.Register("godot_resource", Language())
}
