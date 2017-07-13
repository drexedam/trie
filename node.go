package trie

// treeNode provides an interface for node implementations
type treeNode interface {
	// Child returns the child node for the given rune or nil if else
	Child(r rune) treeNode
	// AddChild adds a new child to the node
	AddChild(r rune, n treeNode)
	// Remove removes an entry if it exists
	Remove(s string)
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

func (mn *mapNode) Remove(s string) {
	runes := []rune(s)
	value, ok := mn.children[runes[0]]
	if ok {
		if value.remove(mn, runes, 1) {
			delete(mn.children, runes[0])
		}
	}
}

// remove tries to remove entries recursive
func (mn *mapNode) remove(previous *mapNode, runes []rune, position int) bool {
	if mn.value == string(runes) { // Found last node
		mn.value = ""
		return len(mn.children) == 0 // Remove me if i have no children
	}

	curKey := runes[position]
	value, ok := mn.children[curKey]
	position++

	if ok {
		if value.remove(mn, runes, position) { // Can i remove my child?
			delete(mn.children, curKey)
		}
	}

	return len(mn.children) == 0 // Remove me if i have no children
}