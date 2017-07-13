// Package trie provides implementations of trie trees
package trie

// Trie provides a basic interface for any tree implementation
type Trie interface {
	// Insert inserts a string into the tree
	Insert(s string)
	// Find searches for the given value
	Find(s string) bool
}

// New creates a new instance of the naive default implementation using
// a map for holding children
func New() Trie {
	return &defaultTrie{newMapNode()}
}