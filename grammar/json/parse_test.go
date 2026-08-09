package grammar_json_test

import (
	"testing"

	"github.com/lewtec/leaven-tree-sitter/grammar"
	grammar_json "github.com/lewtec/leaven-tree-sitter/grammar/json"
)

func TestParseObject(t *testing.T) {
	p := grammar.NewParser()
	t.Cleanup(p.Delete)
	if !p.SetLanguage(grammar_json.Language()) {
		t.Fatal("SetLanguage failed")
	}
	src := []byte(`{"a":1}`)
	tree := p.ParseBytes(src)
	t.Cleanup(tree.Delete)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}
	if root.HasError() {
		t.Fatalf("parse error:\n%s", root.PrintTree())
	}
	if root.Type() == "" {
		t.Fatal("empty root type")
	}
	if root.ChildCount() == 0 {
		t.Fatal("no children")
	}
	child := root.Child(0)
	if child.IsNull() {
		t.Fatal("child 0 null")
	}
	q, err := grammar.NewQuery(grammar_json.Language(), "(pair) @p")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(q.Delete)
	matches := q.ExecuteMatches(root, src)
	if len(matches) == 0 {
		t.Fatal("expected pair capture")
	}
}
