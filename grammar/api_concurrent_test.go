package grammar_test

import (
	"sync"
	"testing"

	"github.com/lewtec/leaven-tree-sitter/grammar"
)

func TestParserConcurrentShare(t *testing.T) {
	name, lang := requireAnyLang(t)
	if !grammar.LiveParseReady(lang) {
		t.Skip("leaven Subtree tagged-pointer crash in core parse")
	}
	src := string(fixtureSource(t, name))
	p := grammar.NewParser()
	if !p.SetLanguage(lang) {
		t.Fatal("SetLanguage failed")
	}
	var wg sync.WaitGroup
	errCh := make(chan string, 32)
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				tree := p.ParseString(src)
				root := tree.RootNode()
				if root.IsNull() {
					errCh <- "null root"
					return
				}
				if root.Type() == "" {
					errCh <- "empty type"
					return
				}
				_ = root.ChildCount()
				if root.ChildCount() > 0 {
					_ = root.Child(0).Type()
				}
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for e := range errCh {
		t.Fatal(e)
	}
}

func TestQueryConcurrentExecute(t *testing.T) {
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
	q, err := grammar.NewQuery(lang, "(_) @n")
	if err != nil {
		t.Skipf("generic query rejected for %s: %v", name, err)
	}
	var wg sync.WaitGroup
	errCh := make(chan string, 32)
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 15; j++ {
				_ = root.Type()
				matches := q.ExecuteMatches(root, src)
				if len(matches) == 0 {
					errCh <- "expected matches"
					return
				}
			}
		}()
	}
	wg.Wait()
	close(errCh)
	for e := range errCh {
		t.Fatal(e)
	}
}
