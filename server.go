package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"
)

func requestHandler(writer http.ResponseWriter, req *http.Request) {
	start := time.Now()

	var letters [16]string
	err := json.NewDecoder(req.Body).Decode(&letters)
	if err != nil {
		http.Error(writer, "invalid body", http.StatusBadRequest)
		return
	}

	depths := req.URL.Query()["depth"]
	depth := 4

	if len(depths) > 0 {
		depth, err = strconv.Atoi(depths[0])
		if err != nil {
			http.Error(writer, "invalid depth", http.StatusBadRequest)
			return
		}

		depth = int(math.Min(10, float64(depth)))
		depth = int(math.Max(3, float64(depth)))
	}

	resultLims := req.URL.Query()["resultLim"]
	resultLim := 128

	if len(resultLims) > 0 {
		resultLim, err = strconv.Atoi(resultLims[0])
		if err != nil {
			http.Error(writer, "invalid result limit", http.StatusBadRequest)
			return
		}
	}

	fmt.Printf("depth: %v\n", depth)

	board := createBoard(letters)
	words := findAllWords(board, depth)
	words = filterWords(words, resultLim)
	words = sortByLength(words)

	err = json.NewEncoder(writer).Encode(words)
	if err != nil {
		http.Error(writer, "error encoding response", 500)
		return
	}

	fmt.Printf("total: %v\n\n", time.Since(start))
}
