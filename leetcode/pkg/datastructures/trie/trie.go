package trie

// TrieNode represents a node in the Trie.
const alphabetSize = 26 // Assuming lowercase English letters

type TrieNode struct {
	children [alphabetSize]*TrieNode
	isEndOfWord bool
}

// Trie represents the Trie itself.
type Trie struct {
	root *TrieNode
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert inserts a word into the Trie.
func (t *Trie) Insert(word string) {
	curr := t.root
	for _, char := range word {
		index := char - 'a' // Assuming lowercase English letters
		if curr.children[index] == nil {
			curr.children[index] = NewTrieNode()
		}
		curr = curr.children[index]
	}
	curr.isEndOfWord = true
}

// Search searches for a word in the Trie.
// Returns true if the word is present, false otherwise.
func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, char := range word {
		index := char - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.isEndOfWord
}

// StartsWith searches if there is any word in the Trie that starts with the given prefix.
// Returns true if there is such a word, false otherwise.
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, char := range prefix {
		index := char - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return true
}