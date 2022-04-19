package main

const BOARD_SIZE = 4
const TILE_COUNT = BOARD_SIZE * BOARD_SIZE

type Board struct {
	tiles [BOARD_SIZE][BOARD_SIZE]rune
	solutions Solutions
}

func makeBoard(letters [TILE_COUNT]rune) Board {
	board := Board {
		solutions: make(Solutions),
	}

	for i := 0; i < BOARD_SIZE; i++ {
		copy(board.tiles[i][:], letters[BOARD_SIZE*i:BOARD_SIZE*(i+1)])
	}

	return board
}