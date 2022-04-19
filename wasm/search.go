package wasm

import "strconv"

type Solutions map[string]string
type VisitedMatrix = [BOARD_SIZE][BOARD_SIZE]bool

// Recursivley searches for words on the board from the given trie node
func (board *Board) Search(row, col int, word, path string, visited VisitedMatrix, node *TrieNode) {
	// Verify that the coordinates are in bounds and this tile hasn't already been visited
	xInBounds := col >= 0 && col < BOARD_SIZE
	yInBounds := row >= 0 && row < BOARD_SIZE
	if !xInBounds || !yInBounds || visited[row][col] {
		return
	}

	// Verify that this tile is a valid next character
	ch := board.Tiles[row][col]
	if _, found := node.children[ch]; !found {
		return
	}

	// Add the letter to the current word and check if its a valid solution
	word += string(ch)
	visited[row][col] = true
	if node.children[ch].fullWord {
		board.Solutions[word] = path
	}

	// Recurse with every adjacent letter
	for xStep := -1; xStep <= 1; xStep++ {
		for yStep := -1; yStep <= 1; yStep++ {
			xDst := col + xStep
			yDst := row + yStep

			newPath := path + ", (" + strconv.Itoa(xDst) + "," + strconv.Itoa(yDst) + ")"
			board.Search(yDst, xDst, word, newPath, visited, node.children[ch])
		}
	}

	// Step back
	visited[row][col] = false
}