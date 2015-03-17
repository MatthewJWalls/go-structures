package main

import "testing"

func TestTrieInsert(t *testing.T) {

	trie := NewTrie()

	trie.Insert("durrr")

	if trie.Size() != 5 {
		t.Errorf("Size was wrong after insert")
	}

	trie.Insert("dead")

	if trie.Size() != 8 {
		t.Errorf("Size was wrong after insert")
	}

	trie.Insert("hello")

	if trie.Size() != 13 {
		t.Errorf("Size was wrong after insert")
	}

}

func TestTrieInsertSubstrings(t *testing.T) {

	trie := NewTrie()
	trie.Insert("hello")
	trie.Insert("hell")

	if len(trie.Complete("hel")) != 2 {
		t.Error("Substrings failed to work.")
	}

}

func TestGetCompletions(t *testing.T) {

	trie := NewTrie()

	trie.Insert("Hello")
	trie.Insert("Happy")
	trie.Insert("Durka")

	if len(trie.Complete("H")) != 2 {
		t.Error("Incorrect number of completions.")
	}

	if len(trie.Complete("D")) != 1 {
		t.Error("Incorrect number of completions.")
	}

	if len(trie.Complete("B")) != 0 {
		t.Error("Incorrect number of completions.")
	}

	if len(trie.Complete("")) != 3 {
		t.Error("Incorrect number of completions.")
	}

}
