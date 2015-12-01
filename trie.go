// Basic Trie implementation
package ds

type Node struct {
	Char  byte
	Next  map[byte]*Node
	Value interface{}
}

type Trie struct {
	Root *Node
}

func add(node *Node, word string, p int, value interface{}) bool {
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
		node.Next[char] = &Node{
			Char: char,
			Next: make(map[byte]*Node),
		}
		isNew = true
	}
	return add(node.Next[char], word, p+1, value) || isNew
}

func (t *Trie) Insert(word string, value interface{}) bool {
	return add(t.Root, word, 0, value)
}

func find(node *Node, word string, p int) interface{} {
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

func (t *Trie) Find(word string) interface{} {
	return find(t.Root, word, 0)
}

func NewTrie() *Trie {
	return &Trie{
		Root: &Node{
			Next: make(map[byte]*Node),
		},
	}
}

func getKeys(n *Node, key string, entries *[]string) {
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

func (t *Trie) Keys() []string {
	entries := make([]string, 0)
	getKeys(t.Root, "", &entries)
	return entries
}
