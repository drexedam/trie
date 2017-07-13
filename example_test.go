package trie

func Example() {
	tr := New()
	tr.Insert("Test")
	if tr.Find("Test") {
		// Found
	}
	tr.Remove("Test")
}
