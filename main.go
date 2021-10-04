package main

import (
	"fmt"
)

func main() {
	const DEPTH = 8
	const LIMIT = 12
	letters := [16]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

	board := createBoard(letters)
	words := findAllWords(board, DEPTH)
	words = sortByLength(words)
	words = filterWords(words, LIMIT)

	fmt.Println(words)
}
