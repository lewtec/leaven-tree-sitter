package grammar_test

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

func requireAnyLang(t *testing.T) (name string, lang grammar.Language) {
	t.Helper()
	if strconv.IntSize == 32 {
		t.Skip("leaven parse fatals on 32-bit (GOARCH=386 traceback unwind)")
	}
	names := grammar.List()
	if len(names) == 0 {
		t.Skip("no languages registered")
	}
	for _, n := range names {
		l, ok := grammar.Get(n)
		if ok && l != nil {
			return n, l
		}
	}
	t.Skip("no non-nil language registered")
	return "", nil
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 6; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			if _, err := os.Stat(filepath.Join(dir, "testdata")); err == nil {
				return dir
			}
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatal("repo root with testdata not found")
	return ""
}

func fixtureSource(t *testing.T, langName string) []byte {
	t.Helper()
	dir := filepath.Join(repoRoot(t), "testdata", langName)
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Skipf("no testdata/%s", langName)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(name, ".golden.json") || name == "SOURCE.txt" || strings.HasPrefix(name, ".") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatal(err)
		}
		return b
	}
	t.Skipf("no source fixture in testdata/%s", langName)
	return nil
}

func TestBuildParseNodeNilAndNull(t *testing.T) {
	if got := grammar.BuildParseNode(nil, nil, ""); got != nil {
		t.Fatalf("nil node: got %#v want nil", got)
	}
	if got := grammar.BuildParseNode(&grammar.Node{}, []byte("{}"), ""); got != nil {
		t.Fatalf("zero node: got %#v want nil", got)
	}
}

func TestBuildParseNodeFromFixture(t *testing.T) {
	name, lang := requireAnyLang(t)
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	src := fixtureSource(t, name)

	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}
	if root.HasError() {
		t.Fatalf("parse has errors:\n%s", root.PrintTree())
	}

	doc := grammar.BuildParseNode(root, src, "")
	if doc == nil {
		t.Fatal("BuildParseNode returned nil")
	}
	if doc.Type == "" {
		t.Fatal("root type empty")
	}
	if doc.Field != "" {
		t.Fatalf("root field: got %q want empty", doc.Field)
	}
	if doc.StartByte != 0 || doc.EndByte != uint32(len(src)) {
		t.Fatalf("root span: [%d,%d) want [0,%d)", doc.StartByte, doc.EndByte, len(src))
	}
}

func TestParseTreeChildAfterDeleteIsNull(t *testing.T) {
	name, lang := requireAnyLang(t)
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	src := fixtureSource(t, name)
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}
	if root.ChildCount() == 0 {
		t.Skip("empty tree")
	}
	child := root.Child(0)
	if child.IsNull() {
		t.Fatal("child 0 null")
	}
	tree.Delete()
	if !root.IsNull() {
		t.Fatal("root still live after Tree.Delete")
	}
	if !child.IsNull() {
		t.Fatal("child still live after Tree.Delete")
	}
}

func TestBuildParseNodeRootFieldName(t *testing.T) {
	name, lang := requireAnyLang(t)
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	src := fixtureSource(t, name)
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	tree := p.ParseBytes(src)
	root := tree.RootNode()
	if root.IsNull() {
		t.Fatal("null root")
	}

	n := grammar.BuildParseNode(root, src, "from_caller")
	if n == nil {
		t.Fatal("BuildParseNode returned nil")
	}
	if n.Field != "from_caller" {
		t.Fatalf("root field: got %q want from_caller", n.Field)
	}
}
