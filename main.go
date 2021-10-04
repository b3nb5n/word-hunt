package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Missing letters from argument list")
	}

	letters := strings.Split(os.Args[1], "")
	if len(letters) > BOARD_SIZE*BOARD_SIZE {
		letters = letters[:16]
	} else if len(letters) < BOARD_SIZE*BOARD_SIZE {
		log.Fatalln("Not enough letters to fill the board")
	}

	depth := 8
	if len(os.Args) >= 3 {
		var err error
		depth, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid depth argument: %v\n", err)
		}
	}

	limit := 24
	if len(os.Args) >= 4 {
		var err error
		limit, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalf("Invalid limit argument: %v\n", err)
		}
	}

	board := createBoard(letters)
	paths := findAllPaths(board, depth)
	sortByLength(paths)
	words := filterPaths(paths, limit)

	fmt.Println(words)
}
