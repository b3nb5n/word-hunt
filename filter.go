package main

func filterPaths(paths []string, limit int) (words []string) {
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
		for i := 0; i < len(words); i++ {
			if words[i] == target {
				return true
			}
		}

		return false
	}

	for _, path := range paths {
		if !resultIncludes(path) && isEnglishWord(path) {
			words = append(words, path)
			if len(words) == limit {
				break
			}
		}
	}

	return words
}
