package main

import (
	"fmt"
)

func main() {
	const DEPTH = 8
	const LIMIT = 12
	letters := Letters{"a", "c", "u", "q", "e", "f", "z", "x", "j", "y", "w", "l", "m", "n", "o", "p"}

	board := createBoard(letters)
	words := findAllPaths(board, DEPTH)
	words = sortByLength(words)
	words = filterPaths(words, LIMIT)

	fmt.Println(words)
}
