package grammar

import "sync"

var (
	liveParseOnce sync.Once
	liveParseOK   bool
)

// LiveParseReady reports whether leaven-translated core can complete a parse
// with lang without panicking (tree-sitter tagged Subtree handling).
//
// The probe runs at most once process-wide (sync.Once). The first caller's
// lang is used; later calls ignore lang and return the cached result.
// Thread-safe for concurrent first callers.
func LiveParseReady(lang Language) bool {
	liveParseOnce.Do(func() {
		liveParseOK = probeLiveParse(lang)
	})
	return liveParseOK
}

func probeLiveParse(lang Language) (ok bool) {
	if lang == nil {
		return false
	}
	defer func() {
		if recover() != nil {
			ok = false
		}
	}()
	p := NewParser()
	defer p.Delete()
	if !p.SetLanguage(lang) {
		return false
	}
	// Tiny source: enough to exercise lex + accept + Subtree summarize.
	tree := p.ParseBytes([]byte("0"))
	root := tree.RootNode()
	return root != nil && !root.IsNull()
}
