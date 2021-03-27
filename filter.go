package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func filterWords(src []string) []string {
	data, err := ioutil.ReadFile("words.json")
	if err != nil {
		log.Fatal(err)
	}

	var englishWords []string
	err = json.Unmarshal(data, &englishWords)
	if err != nil {
		log.Fatal(err)
	}

	isEnglishWord := func(word string) bool {
		upper := len(englishWords) - 1
		lower := 0

		for lower <= upper {
			mid := (lower + upper) / 2

			if englishWords[mid] < word {
				lower = mid + 1
			} else if englishWords[mid] > word {
				upper = mid - 1
			} else {
				return true
			}
		}

		return false
	}

	start := time.Now()

	i := 0
	for i < len(src) {
		if !isEnglishWord(src[i]) {
			src[i] = src[len(src)-1]
			src[len(src)-1] = ""
			src = src[:len(src)-1]
		} else {
			i++
		}
	}

	fmt.Printf("filter strings: %v\n", time.Since(start))

	return src
}
