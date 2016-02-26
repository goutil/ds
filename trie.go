// Package ds provides some useful data structures that can save you some time
// now and then.
package ds

// TrieNode represents a node the trie tree.
type TrieNode struct {
	Char  byte
	Next  map[byte]*TrieNode
	Value interface{}
}

// Trie represents a trie.
type Trie struct {
	Root *TrieNode
}

func add(node *TrieNode, word string, p int, value interface{}) bool {
	if node == nil {
		return false
	}
	if p == len(word) {
		node.Value = value
		return false
	}
	isNew := false
	char := word[p]
	if node.Next[char] == nil {
		node.Next[char] = &TrieNode{
			Char: char,
			Next: make(map[byte]*TrieNode),
		}
		isNew = true
	}
	return add(node.Next[char], word, p+1, value) || isNew
}

// Insert creates or updates an entry in the trie.
func (t *Trie) Insert(word string, value interface{}) bool {
	return add(t.Root, word, 0, value)
}

func find(node *TrieNode, word string, p int) interface{} {
	if node == nil {
		return nil
	}
	if p == len(word) {
		return node.Value
	}
	char := word[p]
	if node.Next[char] == nil {
		return nil
	}
	return find(node.Next[char], word, p+1)
}

// Find finds word in trie and returns its value. It returns nil if word is not
// found.
func (t *Trie) Find(word string) interface{} {
	return find(t.Root, word, 0)
}

// NewTrie returns an initialized trie.
func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Next: make(map[byte]*TrieNode),
		},
	}
}

func getKeys(n *TrieNode, key string, entries *[]string) {
	if n == nil || entries == nil {
		return
	}
	if len(n.Next) == 0 && key != "" {
		*entries = append(*entries, key)
		return
	}

	for k, v := range n.Next {
		newKey := key
		newKey += string(k)
		getKeys(v, newKey, entries)
	}
}

// Keys returns an array with the trie keys.
func (t *Trie) Keys() []string {
	var entries = make([]string, 0)
	getKeys(t.Root, "", &entries)
	return entries
}
