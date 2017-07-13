package trie

import (
	"testing"
)

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

func TestDefaultTrie_Remove(t *testing.T) {
	tr := New()

	tr.Insert("Test")
	if !tr.Find("Test") {
		t.Error("Could not find inserted value")
	}

	tr.Insert("Tetris")
	if !tr.Find("Tetris") {
		t.Error("Could not find inserted value")
	}

	tr.Insert("Tester")
	if !tr.Find("Tester") {
		t.Error("Could not find inserted value")
	}

	tr.Remove("Test")

	if tr.Find("Test") {
		t.Error("Value was not removed")
	}

	if !tr.Find("Tetris") {
		t.Error("Removed too much (Tetris)")
	}

	if !tr.Find("Tester") {
		t.Error("Removed too muc (Tester)")
	}
}

func TestDefaultTrie_Remove2(t *testing.T) {
	tr := New()

	tr.Insert("Test")
	if !tr.Find("Test") {
		t.Error("Could not find inserted value")
	}

	tr.Insert("Tetris")
	if !tr.Find("Tetris") {
		t.Error("Could not find inserted value")
	}

	tr.Insert("Tester")
	if !tr.Find("Tester") {
		t.Error("Could not find inserted value")
	}

	tr.Remove("Tetris")

	if tr.Find("Tetris") {
		t.Error("Value was not removed")
	}

	if !tr.Find("Test") {
		t.Error("Removed too much (Test)")
	}

	if !tr.Find("Tester") {
		t.Error("Removed too muc (Tester)")
	}
}