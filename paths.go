package main

const MIN_LENGTH = 3

func findAllPaths(board Board, depth int) (pathStrings []string) {
	addPath := func(path []*Tile) {
		var pathString string

		for i := 0; i < len(path); i++ {
			pathString += path[i].value
		}

		pathStrings = append(pathStrings, pathString)
	}

	findAllPathsBetween := func(src, dst *Tile) {
		currentPath := []*Tile{src}

		currentPathIncludes := func(target *Tile) bool {
			for _, tile := range currentPath {
				if tile == target {
					return true
				}
			}

			return false
		}

		var search func()
		search = func() {
			nextTiles := currentPath[len(currentPath)-1].adjacents
			for _, tile := range nextTiles {
				if len(currentPath) >= depth {
					break
				} else if tile == dst {
					newPath := append(currentPath, tile)
					if len(newPath) >= MIN_LENGTH {
						addPath(newPath)
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

	for srcI := 0; srcI < BOARD_SIZE*BOARD_SIZE; srcI++ {
		srcRow := srcI / BOARD_SIZE
		srcCol := srcI % BOARD_SIZE
		src := board[srcRow][srcCol]

		for dstI := 0; dstI < BOARD_SIZE*BOARD_SIZE; dstI++ {
			if srcI == dstI {
				continue
			}

			dstRow := dstI / BOARD_SIZE
			dstCol := dstI % BOARD_SIZE
			dst := board[dstRow][dstCol]

			findAllPathsBetween(src, dst)
		}
	}

	return pathStrings
}
