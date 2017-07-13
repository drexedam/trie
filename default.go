package trie

// defaultTrie is a basic implementation of trie
type defaultTrie struct {
	root treeNode
}

func (d *defaultTrie) Insert(s string) {
	runesValue := []rune(s)
	node := d.root
	i := 0
	len := len(runesValue)

	for _, r := range runesValue {
		if newnode := node.Child(r); newnode != nil {
			node = newnode
			i++
		} else {
			break
		}
	}

	var newChild treeNode
	for i < len {
		newChild = newMapNode()
		node.AddChild(runesValue[i], newChild)
		node = newChild
		i++
	}

	if newChild != nil {
		newChild.SetValue(string(runesValue[:i]))
	}

}

func (d *defaultTrie) Find(s string) bool {
	node := d.root
	for _, r := range s {
		if newnode := node.Child(r); newnode != nil {
			node = newnode
		} else {
			return false
		}
	}

	return node.Value() == s
}
