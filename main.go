package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Answer struct {
	word, path string
}

const BOARD_SIZE = 4
const TILE_COUNT = BOARD_SIZE * BOARD_SIZE

var board = [BOARD_SIZE]string{}
var ans = make([]Answer, 0)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Missing letters from argument list")
	}

	letters := os.Args[1]
	if len(letters) > TILE_COUNT {
		letters = letters[:TILE_COUNT]
	} else if len(letters) < TILE_COUNT {
		log.Fatalln("Not enough letters to fill the board")
	}

	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = letters[BOARD_SIZE*i : BOARD_SIZE*(i+1)]
	}

	start := time.Now()
	root := &TrieNode{}
	var visited VisitedMatrix
	err := makeTrie(root)
	if err != nil {
		log.Fatalf("Error making trie: %v", err)
	}

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			path := "(" + strconv.Itoa(row) + "," + strconv.Itoa(col) + ")"
			search(row, col, "", path, visited, root)
		}
	}

	ans = dedupe(ans)
	insertionSort(ans)
	for _, word := range ans {
		fmt.Printf("%v: %v\n", word.word, word.path)
	}

	fmt.Printf("Runtime: %v\n", time.Now().Sub(start))
}
