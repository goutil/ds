package ds

import (
	"sort"
	"testing"
)

func TestInsertEmptyString(t *testing.T) {
	trie := NewTrie()
	if trie.Insert("", nil) {
		t.Errorf("Empty string is already in the trie by definition.")
	}
}

func TestInsertWord(t *testing.T) {
	trie := NewTrie()
	if !trie.Insert("foo", 1) {
		t.Errorf("%q is a new word, it should return true.", "foo")
	}
}

func TestFindNonExistentWord(t *testing.T) {
	trie := NewTrie()
	trie.Insert("foo", 1)
	if trie.Find("bar") != nil {
		t.Errorf("%q is not in the trie, it should return nil.", "bar")
	}
}

func TestFindExistentWord(t *testing.T) {
	trie := NewTrie()
	trie.Insert("foo", 1)
	if trie.Find("foo") != 1 {
		t.Errorf("%q is in the trie, it should return 1.", "foo")
	}
}

func TestKeys(t *testing.T) {
	trie := NewTrie()
	trie.Insert("foo", nil)
	trie.Insert("bar", nil)
	e := trie.Keys()
	sort.Strings(e)
	if len(e) != 2 {
		t.Errorf("Expected %d, got %d\n", 2, len(e))
	}
	if e[0] != "bar" {
		t.Errorf("Expected %v to be %v", e[0], "bar")
	}
	if e[1] != "foo" {
		t.Errorf("Expected %v to be %v", e[1], "foo")
	}
}
