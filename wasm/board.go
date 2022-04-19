package wasm

const BOARD_SIZE = 4
const TILE_COUNT = BOARD_SIZE * BOARD_SIZE

type Board struct {
	Tiles [BOARD_SIZE][BOARD_SIZE]rune
	Solutions Solutions
}

// Returns a new `Board` with an instantiated solutions map and tiles populated from `letters`
func MakeBoard(letters [TILE_COUNT]rune) Board {
	board := Board {
		Solutions: make(Solutions),
	}

	for i := 0; i < BOARD_SIZE; i++ {
		copy(board.Tiles[i][:], letters[BOARD_SIZE*i:BOARD_SIZE*(i+1)])
	}

	return board
}