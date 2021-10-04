package main

import (
	"bufio"
	"log"
	"os"
)

func loadDictionary(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func filterPaths(paths []string, limit int) (words []string) {
	dictionary, err := loadDictionary("./dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}

	inDictionary := func(word string) bool {
		upper := len(dictionary) - 1
		lower := 0

		for lower <= upper {
			mid := (lower + upper) / 2

			if dictionary[mid] < word {
				lower = mid + 1
			} else if dictionary[mid] > word {
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
		if !resultIncludes(path) && inDictionary(path) {
			words = append(words, path)
			if len(words) == limit {
				break
			}
		}
	}

	return words
}
