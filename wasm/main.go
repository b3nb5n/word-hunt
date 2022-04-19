package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func findWords() js.Func {
	return js.FuncOf(func (this js.Value, args []js.Value) any {
		if len(args) < 2 {
			return "Invalid number of args"
		}

		for _, arg := range args {
			fmt.Println(arg.Type().String())
		}

		return nil
	})
}

func p(letters [TILE_COUNT]rune, words []string) Solutions {
	board := makeBoard(letters)
	trieRoot := makeTrie(words)

	// Search for words starting at each letter
	var visited VisitedMatrix
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			path := "(" + strconv.Itoa(row) + "," + strconv.Itoa(col) + ")"
			board.search(row, col, "", path, visited, trieRoot)
		}
	}
	
	return board.solutions
}