package main

const BOARD_SIZE = 4

type Board [BOARD_SIZE][BOARD_SIZE]*Tile

type Tile struct {
	value     string
	adjacents []*Tile
}

type TileCoordinates struct {
	row int
	col int
}

func createBoard(letters []string) (board Board) {
	addEdge := func(src, dst *Tile) {
		var adjacent bool

		for _, tile := range src.adjacents {
			if tile == dst {
				adjacent = true
				break
			}
		}

		if !adjacent {
			src.adjacents = append(src.adjacents, dst)
		}
	}

	getAdjacentCoordinates := func(row, col int) (result []TileCoordinates) {
		for rowOp := -1; rowOp <= 1; rowOp++ {
			adjRow := row + rowOp
			if adjRow < 0 || adjRow >= BOARD_SIZE {
				continue
			}

			for colOp := -1; colOp <= 1; colOp++ {
				adjCol := col + colOp
				if adjCol < 0 || adjCol >= BOARD_SIZE {
					continue
				} else if adjRow == row && adjCol == col {
					continue
				}

				result = append(result, TileCoordinates{adjRow, adjCol})
			}
		}

		return result
	}

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			i := row*BOARD_SIZE + col
			board[row][col] = &Tile{value: letters[i]}
		}
	}

	for row := 0; row < BOARD_SIZE; row++ {
		for col := 0; col < BOARD_SIZE; col++ {
			adjacencies := getAdjacentCoordinates(row, col)

			for i := 0; i < len(adjacencies); i++ {
				adjacent := adjacencies[i]
				addEdge(board[row][col], board[adjacent.row][adjacent.col])
			}
		}
	}

	return board
}
