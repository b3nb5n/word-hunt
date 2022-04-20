package main

type TrieNode struct {
	children map[rune]*TrieNode
	fullWord bool
}

// Returns a new trie node with an instantiated `children` map
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

// Inserts a new word starting at the given `root`
func (root *TrieNode) Insert(word string) {
	curr := root;

	for _, ch := range word {
		// Add new letters to the nodes children
		if _, exists := curr.children[ch]; !exists {
			curr.children[ch] = NewTrieNode()
		}

		// step to the new node to add the next character
		curr = curr.children[ch]
	}

	curr.fullWord = true
}

// Returns a new trie node with each of the given `words` as children
func MakeTrie(words []string) *TrieNode {
	root := NewTrieNode()
	for _, word := range words {
		root.Insert(word)
	}

	return root
}