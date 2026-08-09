package grammar

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/lewtec/leaven/libc"
)

// Language is an immutable tree-sitter language. Safe to share across parsers
// and goroutines.
type Language = *TSLanguage

// Parser is a tree-sitter parser backed by leaven-translated core.
//
// Ownership is GC-managed (runtime.Cleanup); Delete is optional.
//
// Methods are safe for concurrent use: an internal mutex serializes native
// parse work. Parse returns a Tree that wraps the native *TSTree; node
// methods call into that tree. Prefer one Parser per goroutine under heavy
// parallel load.
type Parser struct {
	mu      sync.Mutex
	p       *TSParser
	lang    Language
	cleanup runtime.Cleanup
}

// Tree wraps a native tree-sitter tree. Keep it reachable while using Nodes
// from it. Concurrent reads are safe if the tree is not edited or Deleted.
// Delete is optional; the GC frees the native tree.
type Tree struct {
	native  *TSTree
	lang    Language
	cleanup runtime.Cleanup
}

// Node is a handle into a Tree. Keep the *Tree reachable while using the Node.
type Node struct {
	n    TSNode
	tree *Tree
}

// NewParser creates a parser. Callers need not Delete; the GC will free it.
func NewParser() *Parser {
	p := ts_parser_new()
	out := &Parser{p: p}
	if p != nil {
		out.cleanup = runtime.AddCleanup(out, freeParser, p)
	}
	return out
}

// SetLanguage sets the language for parsing.
func (p *Parser) SetLanguage(lang Language) bool {
	if p == nil {
		return false
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.p == nil {
		return false
	}
	if !ts_parser_set_language(p.p, lang) {
		return false
	}
	p.lang = lang
	return true
}

// ParseString parses a string into a Tree.
func (p *Parser) ParseString(source string) *Tree {
	return p.ParseBytes([]byte(source))
}

// ParseBytes parses a contiguous UTF-8 source buffer into a Tree.
// The buffer is copied into a NUL-terminated scratch for the native call.
func (p *Parser) ParseBytes(source []byte) *Tree {
	if p == nil {
		return &Tree{}
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.p == nil {
		return &Tree{lang: p.lang}
	}
	return wrapTree(p.lang, p.parseNativeLocked(source))
}

func wrapTree(lang Language, native *TSTree) *Tree {
	t := &Tree{native: native, lang: lang}
	if native != nil {
		t.cleanup = runtime.AddCleanup(t, ts_tree_delete, native)
	}
	return t
}

// parseNativeLocked parses source and returns the native tree.
// Caller must hold p.mu and must either wrap or ts_tree_delete the result.
func (p *Parser) parseNativeLocked(source []byte) *TSTree {
	if p.p == nil {
		return nil
	}
	buf := nulTerminate(source)
	return ts_parser_parse_string(p.p, nil, &buf[0], int32(len(source)))
}

func freeParser(p *TSParser) {
	if p != nil {
		ts_parser_delete(p)
	}
}

// Delete eagerly frees the parser. Optional: the GC will free it if omitted.
func (p *Parser) Delete() {
	if p == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.p == nil {
		return
	}
	p.cleanup.Stop()
	freeParser(p.p)
	p.p = nil
	p.lang = nil
}

// Delete eagerly frees the native tree. Optional: the GC will free it if omitted.
// Do not use Nodes from this Tree after Delete.
func (t *Tree) Delete() {
	if t == nil || t.native == nil {
		return
	}
	t.cleanup.Stop()
	ts_tree_delete(t.native)
	t.native = nil
}

// RootNode returns the root node of the tree.
func (t *Tree) RootNode() *Node {
	if t == nil || t.native == nil {
		return &Node{}
	}
	var root TSNode
	ts_tree_root_node(&root, t.native)
	return &Node{n: root, tree: t}
}

func (n *Node) live() bool {
	return n != nil && n.tree != nil && n.tree.native != nil && !ts_node_is_null(&n.n)
}

// Type returns the node type as a string.
func (n *Node) Type() string {
	if !n.live() {
		return ""
	}
	return cString(ts_node_type(&n.n))
}

// ChildCount returns the number of children.
func (n *Node) ChildCount() uint32 {
	if !n.live() {
		return 0
	}
	return uint32(ts_node_child_count(&n.n))
}

// Child returns the child at the given index.
func (n *Node) Child(index uint32) *Node {
	if !n.live() {
		return &Node{}
	}
	var child TSNode
	ts_node_child(&child, &n.n, int32(index))
	return &Node{n: child, tree: n.tree}
}

// FieldNameForChild returns the field name for the child at the given index.
func (n *Node) FieldNameForChild(index uint32) string {
	if !n.live() {
		return ""
	}
	return cString(ts_node_field_name_for_child(&n.n, int32(index)))
}

// NamedChildCount returns the number of named children.
func (n *Node) NamedChildCount() uint32 {
	if !n.live() {
		return 0
	}
	return uint32(ts_node_named_child_count(&n.n))
}

// NamedChild returns the named child at the given index.
func (n *Node) NamedChild(index uint32) *Node {
	if !n.live() {
		return &Node{}
	}
	var child TSNode
	ts_node_named_child(&child, &n.n, int32(index))
	return &Node{n: child, tree: n.tree}
}

// StartByte returns the start byte offset.
func (n *Node) StartByte() uint32 {
	if !n.live() {
		return 0
	}
	return uint32(ts_node_start_byte(&n.n))
}

// EndByte returns the end byte offset.
func (n *Node) EndByte() uint32 {
	if !n.live() {
		return 0
	}
	return uint32(ts_node_end_byte(&n.n))
}

// String returns the S-expression representation of the node.
func (n *Node) String() string {
	if !n.live() {
		return ""
	}
	p := ts_node_string(&n.n)
	if p == nil {
		return ""
	}
	s := cString(p)
	libc.Free(p)
	return s
}

// IsNull returns true if the node is null.
func (n *Node) IsNull() bool {
	return !n.live()
}

// IsNamed returns true if the node is named.
func (n *Node) IsNamed() bool {
	return n.live() && ts_node_is_named(&n.n)
}

// IsExtra returns true if the node is extra.
func (n *Node) IsExtra() bool {
	return n.live() && ts_node_is_extra(&n.n)
}

// IsError returns true if the node is an error.
func (n *Node) IsError() bool {
	return n.live() && ts_node_is_error(&n.n)
}

// HasError returns true if the node or any descendant has an error.
func (n *Node) HasError() bool {
	return n.live() && ts_node_has_error(&n.n)
}

// HasChanges returns true if the node has changed.
func (n *Node) HasChanges() bool {
	return n.live() && ts_node_has_changes(&n.n)
}

// PrintTree returns the node tree in S-expression format, or "(null)" if null.
func (n *Node) PrintTree() string {
	if n.IsNull() {
		return "(null)"
	}
	return n.String()
}

// nulTerminate returns a copy of b with a trailing NUL for C string APIs.
func nulTerminate(b []byte) []byte {
	out := make([]byte, len(b)+1)
	copy(out, b)
	return out
}

// cString converts a NUL-terminated *byte to a Go string.
func cString(p *byte) string {
	if p == nil {
		return ""
	}
	n := 0
	for *(*byte)(unsafe.Add(unsafe.Pointer(p), n)) != 0 {
		n++
	}
	if n == 0 {
		return ""
	}
	return string(unsafe.Slice(p, n))
}
