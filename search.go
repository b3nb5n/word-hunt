package main

import "strconv"

type VisitedMatrix = [BOARD_SIZE][BOARD_SIZE]bool

func search(row, col int, word, path string, visited VisitedMatrix, node *TrieNode) {
	if row < 0 || row >= BOARD_SIZE || col < 0 || col >= BOARD_SIZE {
		return
	} else if visited[row][col] {
		return
	}

	letter := string(board[row][col])

	if _, found := node.children[letter]; !found {
		return
	}

	word += letter
	visited[row][col] = true

	if len(word) > 3 && node.children[letter].fullWord {
		ans = append(ans, Answer{
			word: word,
			path: path,
		})
	}

	for xStep := -1; xStep <= 1; xStep++ {
		for yStep := -1; yStep <= 1; yStep++ {
			xDst := col + xStep
			yDst := row + yStep

			yInBounds := yDst >= 0 && yDst < BOARD_SIZE
			xInBounds := xDst >= 0 && xDst < BOARD_SIZE

			if yInBounds && xInBounds && !visited[yDst][xDst] {
				newPath := path + ", (" + strconv.Itoa(xDst) + "," + strconv.Itoa(yDst) + ")"
				search(yDst, xDst, word, newPath, visited, node.children[letter])
			}
		}
	}

	visited[row][col] = false
}
