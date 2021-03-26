package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		for i := 0; i < len(englishWords); i++ {
			if englishWords[i] == word {
				return true
			}
		}

		return false
	}

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

	return src
}
