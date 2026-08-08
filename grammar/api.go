package grammar

import (
	"runtime"
	"strings"
	"sync"
	"unsafe"
)

// Language is an immutable tree-sitter language. Safe to share across parsers
// and goroutines.
type Language = *TSLanguage

// Parser is a tree-sitter parser backed by leaven-translated core.
//
// Ownership is GC-managed (runtime.Cleanup); Delete is optional.
//
// Methods are safe for concurrent use: an internal mutex serializes native
// parse work. Parse returns a pure-Go Tree snapshot and frees the native tree
// before return, so Tree/Node methods never touch the parser.
// Prefer one Parser per goroutine under heavy parallel load.
type Parser struct {
	mu      sync.Mutex
	p       *TSParser
	lang    Language
	cleanup runtime.Cleanup
}

// Tree is an immutable pure-Go snapshot of a parse tree.
// Concurrent use needs no locking. Delete only drops the root for earlier GC.
type Tree struct {
	root *nodeData
	lang Language
}

// Node is a handle into a Tree snapshot. Keep the *Tree reachable while using
// Nodes from it. Concurrent reads are safe.
type Node struct {
	data *nodeData
	tree *Tree
}

// nodeData is one syntax node in the pure-Go snapshot.
type nodeData struct {
	typ        string
	start, end uint32
	named      bool
	extra      bool
	isError    bool
	hasError   bool
	hasChanges bool
	// path is child indices from the tree root (empty for root). Used to
	// re-locate the node after a temporary reparse for queries.
	path     []uint32
	fields   []string
	children []*nodeData
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

// ParseString parses a string into a pure-Go Tree snapshot.
func (p *Parser) ParseString(source string) *Tree {
	return p.ParseBytes([]byte(source))
}

// ParseBytes parses a contiguous UTF-8 source buffer into a pure-Go Tree.
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
	tree := p.parseNativeLocked(source)
	return snapshotAndFree(p, tree)
}

// snapshotAndFree copies the native tree into a pure-Go Tree and deletes it.
// Caller must hold p.mu.
func snapshotAndFree(p *Parser, tree *TSTree) *Tree {
	t := &Tree{lang: p.lang}
	if tree == nil {
		return t
	}
	var root TSNode
	ts_tree_root_node(&root, tree)
	t.root = captureNode(&root, nil)
	ts_tree_delete(tree)
	return t
}

// captureNode recursively copies one native node and its descendants.
func captureNode(node *TSNode, path []uint32) *nodeData {
	if node == nil || ts_node_is_null(node) {
		return nil
	}
	d := &nodeData{
		typ:        cString(ts_node_type(node)),
		start:      uint32(ts_node_start_byte(node)),
		end:        uint32(ts_node_end_byte(node)),
		named:      ts_node_is_named(node),
		extra:      ts_node_is_extra(node),
		isError:    ts_node_is_error(node),
		hasError:   ts_node_has_error(node),
		hasChanges: ts_node_has_changes(node),
		path:       path,
	}
	count := uint32(ts_node_child_count(node))
	if count == 0 {
		return d
	}
	d.children = make([]*nodeData, 0, count)
	d.fields = make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		field := cString(ts_node_field_name_for_child(node, int32(i)))
		var child TSNode
		ts_node_child(&child, node, int32(i))
		childPath := append(append([]uint32(nil), path...), i)
		d.fields = append(d.fields, field)
		d.children = append(d.children, captureNode(&child, childPath))
	}
	return d
}

// parseNativeLocked parses source and returns the native tree without
// snapshotting. Caller must hold p.mu and must ts_tree_delete the result.
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

// Delete drops the snapshot root so a large tree can be collected earlier.
func (t *Tree) Delete() {
	if t == nil {
		return
	}
	t.root = nil
}

// RootNode returns the root node of the tree.
func (t *Tree) RootNode() *Node {
	if t == nil || t.root == nil {
		return &Node{}
	}
	return &Node{data: t.root, tree: t}
}

// Type returns the node type as a string.
func (n *Node) Type() string {
	if n == nil || n.data == nil {
		return ""
	}
	return n.data.typ
}

// ChildCount returns the number of children.
func (n *Node) ChildCount() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return uint32(len(n.data.children))
}

// Child returns the child at the given index.
func (n *Node) Child(index uint32) *Node {
	if n == nil || n.data == nil || int(index) >= len(n.data.children) {
		return &Node{}
	}
	return &Node{data: n.data.children[index], tree: n.tree}
}

// FieldNameForChild returns the field name for the child at the given index.
func (n *Node) FieldNameForChild(index uint32) string {
	if n == nil || n.data == nil || int(index) >= len(n.data.fields) {
		return ""
	}
	return n.data.fields[index]
}

// NamedChildCount returns the number of named children.
func (n *Node) NamedChildCount() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	var c uint32
	for _, ch := range n.data.children {
		if ch != nil && ch.named {
			c++
		}
	}
	return c
}

// NamedChild returns the named child at the given index.
func (n *Node) NamedChild(index uint32) *Node {
	if n == nil || n.data == nil {
		return &Node{}
	}
	var seen uint32
	for _, ch := range n.data.children {
		if ch == nil || !ch.named {
			continue
		}
		if seen == index {
			return &Node{data: ch, tree: n.tree}
		}
		seen++
	}
	return &Node{}
}

// StartByte returns the start byte offset.
func (n *Node) StartByte() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return n.data.start
}

// EndByte returns the end byte offset.
func (n *Node) EndByte() uint32 {
	if n == nil || n.data == nil {
		return 0
	}
	return n.data.end
}

// String returns the S-expression representation of the node.
func (n *Node) String() string {
	if n == nil || n.data == nil {
		return ""
	}
	var b strings.Builder
	writeSexpr(&b, n.data)
	return b.String()
}

func writeSexpr(b *strings.Builder, d *nodeData) {
	if d == nil {
		return
	}
	if !d.named && len(d.children) == 0 {
		b.WriteString(quoteSexprAtom(d.typ))
		return
	}
	b.WriteByte('(')
	b.WriteString(d.typ)
	for i, ch := range d.children {
		if ch == nil {
			continue
		}
		b.WriteByte(' ')
		if i < len(d.fields) && d.fields[i] != "" {
			b.WriteString(d.fields[i])
			b.WriteString(": ")
		}
		writeSexpr(b, ch)
	}
	b.WriteByte(')')
}

func quoteSexprAtom(s string) string {
	return `"` + strings.ReplaceAll(s, `"`, `\"`) + `"`
}

// IsNull returns true if the node is null.
func (n *Node) IsNull() bool {
	return n == nil || n.data == nil
}

// IsNamed returns true if the node is named.
func (n *Node) IsNamed() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.named
}

// IsExtra returns true if the node is extra.
func (n *Node) IsExtra() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.extra
}

// IsError returns true if the node is an error.
func (n *Node) IsError() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.isError
}

// HasError returns true if the node or any descendant has an error.
func (n *Node) HasError() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.hasError
}

// HasChanges returns true if the node has changed.
func (n *Node) HasChanges() bool {
	if n == nil || n.data == nil {
		return false
	}
	return n.data.hasChanges
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
