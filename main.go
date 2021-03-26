package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("straring server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	log.Printf("listening on port %s", port)

	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Node struct {
	value string
	edges []*Node
}

type CellCoords struct {
	row int
	col int
}

func createBoard(src [16]string) [4][4]*Node {
	size := 4

	addEdge := func(src, destination *Node) {
		hasEdge := func(src, destination *Node) bool {
			for i := 0; i < len(src.edges); i++ {
				if src.edges[i] == destination {
					return true
				}
			}

			return false
		}

		if !hasEdge(src, destination) {
			src.edges = append(src.edges, destination)
		}
	}

	getAdjacentCoords := func(row, col int) []CellCoords {
		var result []CellCoords

		for rowOp := -1; rowOp < 2; rowOp++ {
			adjacentRow := row + rowOp
			if adjacentRow < 0 || adjacentRow > size-1 {
				continue
			}

			for colOp := -1; colOp < 2; colOp++ {
				adjacentCol := col + colOp
				if adjacentCol < 0 || adjacentCol > size-1 {
					continue
				}

				if adjacentRow == row && adjacentCol == col {
					continue
				}

				result = append(result, CellCoords{adjacentRow, adjacentCol})
			}
		}

		return result
	}

	var board [4][4]*Node

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			i := row*size + col
			board[row][col] = &Node{value: src[i]}
		}
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			adjacentCoords := getAdjacentCoords(row, col)

			for i := 0; i < len(adjacentCoords); i++ {
				adjacentCell := adjacentCoords[i]
				addEdge(board[row][col], board[adjacentCell.row][adjacentCell.col])
			}
		}
	}

	return board
}

func requestHandler(writer http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(writer, "can't read body", http.StatusBadRequest)
		return
	}

	var letters [16]string
	err = json.Unmarshal(body, &letters)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(writer, "invalid body", http.StatusBadRequest)
		return
	}

	board := createBoard(letters)
	res := ""

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			node := board[row][col]
			res += node.value + ": [ "

			for i := 0; i < len(node.edges); i++ {
				res += node.edges[i].value + ", "
			}

			res += "]\n"
		}
	}

	fmt.Fprint(writer, res)
}
