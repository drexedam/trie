package trie

import "os"

func Example() {
	tr := New()
	tr.Insert("Test")
	if tr.Find("Test") {
		// Found
	}
	tr.Remove("Test")
}

// ExampleDefaultTrie_PrettyPrint shows how to use PrettyPrint
// Example output:
//. T
//. . e
//. . . s
//. . . . t
//. . . . . e
//. . . . . . r
//. . . . . o
//. . . . . . r
//. . . t
//. . . . r
//. . . . . i
//. . . . . . s
func ExampleDefaultTrie_PrettyPrint() {
	tr := New()
	tr.Insert("Test")
	tr.Insert("Tester")
	tr.Insert("Tetris")
	tr.Insert("Testor")
	tr.PrettyPrint(os.Stdout, "")
}