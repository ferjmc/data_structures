// Trie (we pronounce "try") or prefix tree is a tree data structure,
// which is used for retrieval of a key in a dataset of strings.
// There are various applications of this very efficient data structure such as :
// autocomplete, spell checker, IP routing (Longest prefix matching), T9 predictive text.
package trie

// Trie is a rooted tree. Its nodes have the following fields:
// links: Maximum of 26 links to its children, the number of lowercase latin letters.
// isEnd: which specifies whether the node is the end of a key, or is just a key prefix.
type Trie struct {
	links [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

// Insert a word into the Trie.
func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		idx := ch - 'a'
		if node.links[idx] == nil {
			node.links[idx] = &Trie{}
		}
		node = node.links[idx]
	}
	node.isEnd = true
}

// Search returns if a word is in the Trie.
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		idx := ch - 'a'
		if node.links[idx] == nil {
			return false
		}
		node = node.links[idx]
	}
	return node.isEnd
}

// StartsWith return if there is any word in the trie that starts with the given prefix.
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.links[idx] == nil {
			return false
		}
		node = node.links[idx]
	}
	return true
}
