package main

import "strconv"

type Solutions map[string]string
type VisitedMatrix = [BOARD_SIZE][BOARD_SIZE]bool

// Recursivley searches for words on the board from the given trie node
func (board *Board) search(row, col int, word, path string, visited VisitedMatrix, node *TrieNode) {
	// Check if coordinates are in bounds
	if row < 0 || row >= BOARD_SIZE || col < 0 || col >= BOARD_SIZE {
		return
	} else if visited[row][col] {
		return
	}

	// Check if this character is a valid next ch
	ch := board.tiles[row][col]
	if _, found := node.children[ch]; !found {
		return
	}

	// Add the letter to the current word and check if its a valid solution
	word += string(ch)
	visited[row][col] = true
	if node.children[ch].fullWord {
		board.solutions[word] = path
	}

	// Recurse with every adjacent letter
	for xStep := -1; xStep <= 1; xStep++ {
		for yStep := -1; yStep <= 1; yStep++ {
			xDst := col + xStep
			yDst := row + yStep

			yInBounds := yDst >= 0 && yDst < BOARD_SIZE
			xInBounds := xDst >= 0 && xDst < BOARD_SIZE

			if yInBounds && xInBounds && !visited[yDst][xDst] {
				newPath := path + ", (" + strconv.Itoa(xDst) + "," + strconv.Itoa(yDst) + ")"
				board.search(yDst, xDst, word, newPath, visited, node.children[ch])
			}
		}
	}

	// Step back
	visited[row][col] = false
}