package grammar

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

// Query wraps a compiled tree-sitter query (leaven core).
//
// Ownership is GC-managed; Delete is optional.
// Methods are safe for concurrent use via an internal mutex.
// ExecuteMatches re-parses source (Trees are pure-Go snapshots), runs the
// query, and frees native state before return.
type Query struct {
	mu      sync.Mutex
	q       *TSQuery
	lang    Language
	cleanup runtime.Cleanup
}

// QueryCursor wraps a query cursor. Cleanup pins the parent *Query.
// Methods are safe for concurrent use; they lock the parent Query.
type QueryCursor struct {
	c       *TSQueryCursor
	query   *Query
	cleanup runtime.Cleanup
}

// TSQueryError is the tree-sitter query compile error code (api.h).
// Leaven does not emit this enum as a named type; values match C.
type TSQueryError int32

const (
	TSQueryErrorNone TSQueryError = iota
	TSQueryErrorSyntax
	TSQueryErrorNodeType
	TSQueryErrorField
	TSQueryErrorCapture
	TSQueryErrorStructure
	TSQueryErrorLanguage
)

// QueryCompileError is returned when NewQuery fails to compile.
type QueryCompileError struct {
	Offset uint32
	Type   TSQueryError
}

func (e *QueryCompileError) Error() string {
	return fmt.Sprintf("query compile error at offset %d: %s", e.Offset, queryErrorName(e.Type))
}

// QueryCapture is one capture in a QueryMatch.
type QueryCapture struct {
	Index     uint32 `json:"index"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	StartByte uint32 `json:"start_byte"`
	EndByte   uint32 `json:"end_byte"`
	Text      string `json:"text,omitempty"`
}

// QueryMatch is one match from ExecuteMatches.
type QueryMatch struct {
	ID           uint32         `json:"id"`
	PatternIndex uint16         `json:"pattern_index"`
	Captures     []QueryCapture `json:"captures"`
}

// NewQuery compiles a query. Callers need not Delete; the GC will free it.
func NewQuery(lang Language, source string) (*Query, error) {
	if lang == nil {
		return nil, &QueryCompileError{Type: TSQueryErrorLanguage}
	}
	buf := nulTerminate([]byte(source))
	var errOffset int32
	var errType int32
	q := ts_query_new(lang, &buf[0], int32(len(source)), &errOffset, &errType)
	if q == nil {
		return nil, &QueryCompileError{
			Offset: uint32(errOffset),
			Type:   TSQueryError(errType),
		}
	}
	out := &Query{q: q, lang: lang}
	out.cleanup = runtime.AddCleanup(out, freeQuery, q)
	return out, nil
}

func freeQuery(q *TSQuery) {
	if q != nil {
		ts_query_delete(q)
	}
}

func freeCursor(r cursorRes) {
	if r.c == nil || r.query == nil {
		return
	}
	r.query.mu.Lock()
	defer r.query.mu.Unlock()
	freeCursorUnlocked(r)
}

type cursorRes struct {
	c     *TSQueryCursor
	query *Query
}

// freeCursorUnlocked deletes the native cursor. Caller must hold r.query.mu.
func freeCursorUnlocked(r cursorRes) {
	if r.c == nil {
		return
	}
	ts_query_cursor_delete(r.c)
}

// Delete eagerly frees the query. Optional: the GC will free it if omitted.
func (q *Query) Delete() {
	if q == nil {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.q == nil {
		return
	}
	q.cleanup.Stop()
	freeQuery(q.q)
	q.q = nil
	q.lang = nil
}

// NewCursor creates a cursor for this query.
func (q *Query) NewCursor() *QueryCursor {
	if q == nil {
		return &QueryCursor{}
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.newCursorUnlocked()
}

func (q *Query) newCursorUnlocked() *QueryCursor {
	if q.q == nil {
		return &QueryCursor{}
	}
	c := ts_query_cursor_new()
	out := &QueryCursor{c: c, query: q}
	if c != nil {
		out.cleanup = runtime.AddCleanup(out, freeCursor, cursorRes{c: c, query: q})
	}
	return out
}

// Delete eagerly frees the cursor. Optional: the GC will free it if omitted.
func (c *QueryCursor) Delete() {
	if c == nil || c.c == nil || c.query == nil {
		return
	}
	c.cleanup.Stop()
	freeCursor(cursorRes{c: c.c, query: c.query})
	c.c = nil
	c.query = nil
}

// ExecuteMatches runs the query over root and returns all matches.
//
// Trees are pure-Go snapshots, so this re-parses source with a temporary
// parser, walks root's child-index path to the matching native node, runs the
// query, then frees all native state before return.
// Returns nil if the query is unusable, root is nil/null, or reparse fails.
// Safe for concurrent use.
func (q *Query) ExecuteMatches(root *Node, source []byte) []QueryMatch {
	if q == nil || q.q == nil {
		return nil
	}
	if root == nil || root.data == nil || root.IsNull() {
		return nil
	}

	var lang Language
	if root.tree != nil {
		lang = root.tree.lang
	}
	if lang == nil {
		lang = q.lang
	}
	if lang == nil {
		return nil
	}

	q.mu.Lock()
	defer q.mu.Unlock()
	if q.q == nil {
		return nil
	}

	p := NewParser()
	defer p.Delete()
	if !p.SetLanguage(lang) {
		return nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	tree := p.parseNativeLocked(source)
	if tree == nil {
		return nil
	}
	defer ts_tree_delete(tree)

	var nativeRoot TSNode
	ts_tree_root_node(&nativeRoot, tree)
	nativeNode := walkChildPath(&nativeRoot, root.data.path)
	if ts_node_is_null(nativeNode) {
		return nil
	}

	cursor := q.newCursorUnlocked()
	if cursor.c == nil {
		return nil
	}
	defer func() {
		cursor.cleanup.Stop()
		freeCursorUnlocked(cursorRes{c: cursor.c, query: q})
		cursor.c = nil
		cursor.query = nil
	}()

	ts_query_cursor_exec(cursor.c, q.q, nativeNode)

	matches := make([]QueryMatch, 0)
	var rawMatch TSQueryMatch
	for ts_query_cursor_next_match(cursor.c, &rawMatch) {
		nCap := int(rawMatch.F2) // capture_count
		captures := make([]QueryCapture, 0, nCap)
		if nCap > 0 && rawMatch.F3 != nil {
			rawCaps := unsafe.Slice(rawMatch.F3, nCap)
			for i := range rawCaps {
				cap := &rawCaps[i]
				node := &cap.F0
				idx := uint32(cap.F1)
				start := uint32(ts_node_start_byte(node))
				end := uint32(ts_node_end_byte(node))
				qc := QueryCapture{
					Index:     idx,
					Name:      q.captureNameUnlocked(idx),
					Type:      cString(ts_node_type(node)),
					StartByte: start,
					EndByte:   end,
				}
				if int(end) <= len(source) && start <= end {
					qc.Text = string(source[start:end])
				}
				captures = append(captures, qc)
			}
		}
		matches = append(matches, QueryMatch{
			ID:           uint32(rawMatch.F0),
			PatternIndex: uint16(rawMatch.F1),
			Captures:     captures,
		})
	}
	return matches
}

// walkChildPath follows child indices from root (empty path = root).
// The returned *TSNode points at a heap copy valid for the caller.
func walkChildPath(root *TSNode, path []uint32) *TSNode {
	if root == nil {
		return nil
	}
	cur := *root
	for _, i := range path {
		if ts_node_is_null(&cur) {
			out := cur
			return &out
		}
		var child TSNode
		ts_node_child(&child, &cur, int32(i))
		cur = child
	}
	out := cur
	return &out
}

func (q *Query) captureNameUnlocked(captureIndex uint32) string {
	var length int32
	ptr := ts_query_capture_name_for_id(q.q, int32(captureIndex), &length)
	if ptr == nil || length == 0 {
		return ""
	}
	return string(unsafe.Slice(ptr, int(length)))
}

func queryErrorName(errType TSQueryError) string {
	switch errType {
	case TSQueryErrorSyntax:
		return "syntax"
	case TSQueryErrorNodeType:
		return "node_type"
	case TSQueryErrorField:
		return "field"
	case TSQueryErrorCapture:
		return "capture"
	case TSQueryErrorStructure:
		return "structure"
	case TSQueryErrorLanguage:
		return "language"
	default:
		return "unknown"
	}
}
