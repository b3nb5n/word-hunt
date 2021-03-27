package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

func requestHandler(writer http.ResponseWriter, req *http.Request) {
	start := time.Now()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panicf("Error reading body: %v", err)
		http.Error(writer, "can't read body", http.StatusBadRequest)
		return
	}

	var letters [16]string
	err = json.Unmarshal(body, &letters)
	if err != nil {
		log.Panicf("Error reading body: %v", err)
		http.Error(writer, "invalid body", http.StatusBadRequest)
		return
	}

	depths := req.URL.Query()["depth"]
	depth := 4

	if len(depths) > 0 {
		depth, err = strconv.Atoi(depths[0])
		if err != nil {
			log.Panicf("Error reading depth: %v", err)
			http.Error(writer, "invalid depth", http.StatusBadRequest)
			return
		}

		depth = int(math.Min(10, float64(depth)))
		depth = int(math.Max(3, float64(depth)))
	}

	fmt.Printf("depth: %v\n", depth)

	board := createBoard(letters)
	words := findAllWords(board, depth)
	words = filterWords(words)
	words = sortByLength(words)

	res := "[ "
	for i := 0; i < len(words); i++ {
		res += words[i] + ", "
	}
	res += "]"

	fmt.Printf("total: %v\n\n", time.Since(start))
	fmt.Fprint(writer, res)
}
