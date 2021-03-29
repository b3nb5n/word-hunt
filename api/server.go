package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
)

type Body struct {
	Letters [16]string
	Depth   int
	Limit   int
}

func requestHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()

	fmt.Println("serving request...")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if req.Method == "OPTIONS" {
		fmt.Print("configured CORS headers\n\n")
		return
	}

	var data Body
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		fmt.Printf("error reading request body: %v\n\n", err)
		return
	}

	data.Depth = int(math.Min(10, float64(data.Depth)))
	data.Depth = int(math.Max(3, float64(data.Depth)))

	if data.Limit == 0 {
		data.Limit = 128
	} else {
		data.Limit = int(math.Min(128, float64(data.Limit)))
	}

	fmt.Printf("depth: %v\n", data.Depth)

	board := createBoard(data.Letters)
	words := findAllWords(board, data.Depth)
	words = sortByLength(words)
	words = filterWords(words, data.Limit)

	res, err := json.Marshal(words)
	if err != nil {
		http.Error(w, "error encoding response", 500)
		fmt.Printf("error encoding response: %v\n\n", err)
		return
	}

	fmt.Fprint(w, string(res))

	fmt.Printf("total: %v\n\n", time.Since(start))
}
