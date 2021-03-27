package main

import (
	"fmt"
	"time"
)

func filterWords(src []string, max int) []string {
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

	var result []string
	for i := 0; i < len(src) && len(result) < max; i++ {
		if isEnglishWord(src[i]) {
			result = append(result, src[i])
		}
	}

	fmt.Printf("filter strings: %v\n", time.Since(start))
	return result
}
