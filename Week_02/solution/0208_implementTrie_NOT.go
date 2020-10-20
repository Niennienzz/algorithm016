package solution

// Implement a trie with insert, search, and startsWith methods.
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.
//
// Examples:
//
// Trie trie = new Trie();
//
// trie.insert("apple");
// trie.search("apple");   // returns true
// trie.search("app");     // returns false
// trie.startsWith("app"); // returns true
// trie.insert("app");
// trie.search("app");     // returns true

type Trie struct {
}

func NewTrie() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
}

func (this *Trie) Search(word string) bool {
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}
