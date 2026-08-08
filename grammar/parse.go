package grammar

// ParseOutput is the JSON envelope for a serialized parse tree (e.g. cmd/parse).
// Language and File identify the grammar and input path; Root holds the tree.
// Matches is optional and omitted from JSON when empty.
type ParseOutput struct {
	Language string       `json:"language"`
	File     string       `json:"file"`
	Root     *ParseNode   `json:"root"`
	Matches  []QueryMatch `json:"matches,omitempty"`
}

// ParseNode is a JSON-friendly view of a syntax node: type, optional field
// name relative to the parent, byte span in the source, and children.
// Text is set only on leaves (no children); internal nodes omit it.
type ParseNode struct {
	Type      string       `json:"type"`
	Field     string       `json:"field,omitempty"`
	StartByte uint32       `json:"start_byte"`
	EndByte   uint32       `json:"end_byte"`
	Text      string       `json:"text,omitempty"`
	Children  []*ParseNode `json:"children,omitempty"`
}

// BuildParseNode copies n into a ParseNode tree for JSON encoding.
// n is already a pure-Go snapshot (no native lock). fieldName is the field
// name on the parent (empty for the root). Leaf Text is source[start:end]
// when the span is in range; returns nil if n is nil or null.
func BuildParseNode(n *Node, source []byte, fieldName string) *ParseNode {
	if n == nil || n.IsNull() {
		return nil
	}

	start := n.StartByte()
	end := n.EndByte()
	count := n.ChildCount()
	node := &ParseNode{
		Type:      n.Type(),
		Field:     fieldName,
		StartByte: start,
		EndByte:   end,
	}
	if count == 0 {
		if int(end) <= len(source) && start <= end {
			node.Text = string(source[start:end])
		}
		return node
	}

	node.Children = make([]*ParseNode, 0, count)
	for i := uint32(0); i < count; i++ {
		childNode := BuildParseNode(n.Child(i), source, n.FieldNameForChild(i))
		if childNode != nil {
			node.Children = append(node.Children, childNode)
		}
	}
	return node
}
