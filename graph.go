package main

type Tile struct {
	value     string
	adjacents []*Tile
}

type Board [4][4]*Tile

var boardSize = 4

type TileCoordinates struct {
	row int
	col int
}

func createBoard(src [16]string) Board {
	addEdge := func(src, destination *Tile) {
		var isAdjacent bool

		for i := 0; i < len(src.adjacents); i++ {
			if src.adjacents[i] == destination {
				isAdjacent = true
				break
			}
		}

		if !isAdjacent {
			src.adjacents = append(src.adjacents, destination)
		}
	}

	getAdjacentCoordinates := func(row, col int) []TileCoordinates {
		var result []TileCoordinates

		for rowOp := -1; rowOp <= 1; rowOp++ {
			adjacentRow := row + rowOp
			if adjacentRow < 0 || adjacentRow > boardSize-1 {
				continue
			}

			for colOp := -1; colOp <= 1; colOp++ {
				adjacentCol := col + colOp
				if adjacentCol < 0 || adjacentCol > boardSize-1 {
					continue
				}

				if adjacentRow == row && adjacentCol == col {
					continue
				}

				result = append(result, TileCoordinates{adjacentRow, adjacentCol})
			}
		}

		return result
	}

	var board Board

	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			i := row*boardSize + col
			board[row][col] = &Tile{value: src[i]}
		}
	}

	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			adjacencies := getAdjacentCoordinates(row, col)

			for i := 0; i < len(adjacencies); i++ {
				adjacent := adjacencies[i]
				addEdge(board[row][col], board[adjacent.row][adjacent.col])
			}
		}
	}

	return board
}
