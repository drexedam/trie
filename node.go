package trie

// treeNode provides an interface for node implementations
type treeNode interface {
	// Child returns the child node for the given rune or nil if else
	Child(r rune) treeNode
	// AddChild adds a new child to the node
	AddChild(r rune, n treeNode)
	// Value returns the value of the node or nil
	Value() string
	// SetValue sets the value of the node
	SetValue(s string)
}

// newMapNode creates a new mapNode
func newMapNode() treeNode {
	return &mapNode{children: make(map[rune]*mapNode)}
}

// mapNode uses a map to store its children
type mapNode struct {
	children map[rune]*mapNode
	value    string
}

func (mn *mapNode) Child(r rune) treeNode {
	value, ok := mn.children[r]
	if !ok {
		return nil
	}
	return value
}

func (mn *mapNode) Value() string {
	return mn.value
}

func (mn *mapNode) AddChild(r rune, n treeNode) {
	switch typedN := n.(type) {
	case *mapNode:
		_, ok := mn.children[r]
		if !ok { // Prevent replacement of old child
			mn.children[r] = typedN
		}
	default:
		break //Unsupported treeNode type
	}

}
func (mn *mapNode) SetValue(s string) {
	mn.value = s
}
