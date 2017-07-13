package trie

import "testing"

func TestDefaultTrie(t *testing.T) {
	tr := New()
	tr.Insert("Test")
	if !tr.Find("Test") {
		t.Error("Could not find inserted value")
	}
}

func TestDefaultTrieUmlauts(t *testing.T) {
	tr := New()
	tr.Insert("Töst")
	if !tr.Find("Töst") {
		t.Error("Could not find inserted value")
	}
}

func TestSpecialCharacter(t *testing.T) {
	tr := New()
	tr.Insert("T&st")
	if !tr.Find("T&st") {
		t.Error("Could not find inserted value")
	}
}
