package main

import (
	"fmt"
	"math"
	"time"
)

func findAllWords(board Board, depth int) []string {
	var words []string

	addWord := func(path []*Tile) {
		var word string

		for i := 0; i < len(path); i++ {
			word += path[i].value
		}

		words = append(words, word)
	}

	findAllPathsBetween := func(src, dst *Tile) {
		currentPath := []*Tile{src}

		currentPathIncludes := func(target *Tile) bool {
			for i := 0; i < len(currentPath); i++ {
				if currentPath[i] == target {
					return true
				}
			}

			return false
		}

		var search func()
		search = func() {
			nextTiles := currentPath[len(currentPath)-1].adjacents
			for i := 0; i < len(nextTiles); i++ {
				if len(currentPath) >= depth {
					break
				}

				tile := nextTiles[i]
				if tile == dst {
					newPath := append(currentPath, tile)
					if len(newPath) >= 3 {
						addWord(newPath)
					}
				} else if !currentPathIncludes(tile) {
					currentPath = append(currentPath, tile)
					search()
					currentPath[len(currentPath)-1] = nil
					currentPath = currentPath[:len(currentPath)-1]
				}
			}
		}

		search()
	}

	start := time.Now()

	for srcI := 0; srcI < boardSize; srcI++ {
		srcRow := srcI / boardSize
		srcCol := int(math.Mod(float64(srcI), float64(boardSize)))

		for dstI := 0; dstI < boardSize; dstI++ {
			if srcI == dstI {
				continue
			}

			dstRow := dstI / boardSize
			dstCol := int(math.Mod(float64(dstI), float64(boardSize)))

			findAllPathsBetween(board[srcRow][srcCol], board[dstRow][dstCol])
		}
	}

	fmt.Printf("get combinations: %v\n", time.Since(start))

	return words
}
