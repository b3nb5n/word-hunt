package main

import (
	"fmt"
	"time"
)

func filterWords(src []string, max int) []string {
	start := time.Now()
	var result []string

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

	resultIncludes := func(target string) bool {
		for i := 0; i < len(result); i++ {
			if result[i] == target {
				return true
			}
		}

		return false
	}

	for i := 0; i < len(src) && len(result) < max; i++ {
		if isEnglishWord(src[i]) && !resultIncludes(src[i]) {
			result = append(result, src[i])
		}
	}

	fmt.Printf("removed %v strings in %v\n", len(src)-len(result), time.Since(start))
	return result
}
