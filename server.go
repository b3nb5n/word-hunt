package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func requestHandler(writer http.ResponseWriter, req *http.Request) {
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

	board := createBoard(letters)
	fmt.Printf("letters: %v\n", letters)
	words := findAllWords(board, 4)
	fmt.Printf("strings: %v\n", len(words))

	res := "[ "
	for i := 0; i < len(words); i++ {
		res += words[i] + ", "
	}
	res += "]"

	fmt.Fprint(writer, res)
}
