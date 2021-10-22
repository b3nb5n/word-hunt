package main

import (
	"bufio"
	"os"
)

type TrieNode struct {
	children map[string]*TrieNode
	fullWord bool
}

func (node *TrieNode) init() *TrieNode {
	node.children = make(map[string]*TrieNode)
	return node
}

func makeTrie(root *TrieNode) error {
	if root.children == nil {
		root.init()
	}

	file, err := os.Open("./dictionary.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr := root
		word := scanner.Text()

		for i, ch := range word {
			letter := string(ch)
			if _, found := curr.children[letter]; !found {
				curr.children[letter] = (&TrieNode{}).init()
			}
			curr = curr.children[letter]
			if i == len(word)-1 {
				curr.fullWord = true
			}
		}
	}

	return nil
}
