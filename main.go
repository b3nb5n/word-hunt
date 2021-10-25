package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const BOARD_SIZE = 4
const TILE_COUNT = BOARD_SIZE * BOARD_SIZE

var board = [BOARD_SIZE]string{}
var solutions = make([]Solution, 0)

var score = map[int]int{
	3: 100,
	4: 400,
	5: 800,
	6: 1400,
	7: 1800,
	8: 2200,
}

func main() {
	// Read letters from cli arguments
	if len(os.Args) < 2 {
		log.Fatalln("Missing letters from argument list")
	}

	letters := os.Args[1]
	if len(letters) > TILE_COUNT {
		letters = letters[:TILE_COUNT]
	} else if len(letters) < TILE_COUNT {
		log.Fatalln("Not enough letters to fill the board")
	}

	// Add letters to the board
	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = letters[BOARD_SIZE*i : BOARD_SIZE*(i+1)]
	}

	root, err := makeTrie("./dictionary.txt")
	if err != nil {
		log.Fatalf("Error making trie: %v", err)
	}

	// Search for words starting at each letter
	var visited VisitedMatrix
	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			path := "(" + strconv.Itoa(row) + "," + strconv.Itoa(col) + ")"
			search(row, col, "", path, visited, root)
		}
	}

	// Sort and dedupe solutions
	solutions = dedupeAnswers(solutions)
	sortAnswers(solutions)
	for _, word := range solutions {
		fmt.Printf("%v: %v\n", word.word, word.path)
	}
}
