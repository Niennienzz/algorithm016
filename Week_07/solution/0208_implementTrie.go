package solution

// Implement a trie with insert, search, and startsWith methods.
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.
//
// Examples:
//
// trie := NewTrie()
// trie.insert("apple");
// trie.search("apple");   // returns true
// trie.search("app");     // returns false
// trie.startsWith("app"); // returns true
// trie.insert("app");
// trie.search("app");     // returns true

type Trie struct {
	root *trieNode
}

func NewTrie() Trie {
	return Trie{root: newTrieNode()}
}

func (x *Trie) Insert(word string) {
	node := x.root
	for i := range word {
		currChar := word[i]
		if !node.containsKey(currChar) {
			node.put(currChar, newTrieNode())
		}
		node = node.get(currChar)
	}
	node.isEnd = true
}

func (x *Trie) Search(word string) bool {
	node := x.searchTrieNode(word)
	return node != nil && node.isEnd
}

func (x *Trie) StartsWith(word string) bool {
	return x.searchTrieNode(word) != nil
}

func (x *Trie) searchTrieNode(word string) *trieNode {
	node := x.root
	for i := range word {
		currChar := word[i]
		if node.containsKey(currChar) {
			node = node.get(currChar)
		} else {
			return nil
		}
	}
	return node
}

type trieNode struct {
	links []*trieNode
	isEnd bool
}

func newTrieNode() *trieNode {
	return &trieNode{
		links: make([]*trieNode, 26, 26),
		isEnd: false,
	}
}

func (x trieNode) containsKey(v byte) bool {
	return x.links[v-'a'] != nil
}

func (x trieNode) get(v byte) *trieNode {
	return x.links[v-'a']
}

func (x *trieNode) put(v byte, node *trieNode) {
	x.links[v-'a'] = node
}
