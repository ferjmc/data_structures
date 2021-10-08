package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if len(trie.links) == 0 {
		t.Fatal("inserting apple cannot result in trie with empty links")
	}
	if !trie.Search("apple") {
		t.Fatal("did not find apple avobe inserted")
	}
	if trie.Search("app") {
		t.Fatal("find app when search a full word, given that app is just a prefix")
	}
	if !trie.StartsWith("app") {
		t.Fatal("startsWith app returns false given that the trie have the word apple")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Fatal("did not find app avobe inserted")
	}
}
