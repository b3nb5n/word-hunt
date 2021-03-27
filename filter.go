package main

import (
	"fmt"
	"time"
)

func filterWords(src []string) []string {
	start := time.Now()

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
