package main

import (
	"bufio"
	"fmt"
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

func readDictionary(path string, out chan string) error {
	file, err := os.Open("./dictionary.txt")
	if err != nil {
		return fmt.Errorf("Error reading file %v: %v", path, err)
	}

	// Add each new word to the out channel
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out <- scanner.Text()
	}

	file.Close()
	close(out)
	return nil
}

func makeTrie(dictPath string) (*Node, error) {
	root := newNode()
	wordChan := make(chan string)
	go readDictionary(dictPath, wordChan)

	// For each word (new line deliminated)
	for word := range wordChan {
		curr := root
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
