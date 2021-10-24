package main

import (
	"bufio"
	"os"
)

type Node struct {
	children map[string]*Node
	fullWord bool
}

func newNode() *Node {
	return &Node{
		children: make(map[string]*Node),
	}
}

func makeTrie() (*Node, error) {
	// Load dictionary file
	file, err := os.Open("./dictionary.txt")
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// For each word (new line deliminated)
	root := newNode()
	for scanner.Scan() {
		curr := root
		word := scanner.Text()

		for i, ch := range word {
			// Add new letters to the nodes children
			letter := string(ch)
			if _, found := curr.children[letter]; !found {
				curr.children[letter] = newNode()
			}
			curr = curr.children[letter]

			// Check for end of word
			if i == len(word)-1 {
				curr.fullWord = true
			}
		}
	}

	return root, nil
}
