package main

func dedupeAnswers(answers []Solution) (res []Solution) {
	keys := make(map[string]bool)
	for _, answer := range answers {
		// Add unique words to the result
		if _, found := keys[answer.word]; !found {
			keys[answer.word] = true
			res = append(res, answer)
		}
	}

	return res
}

func sortAnswers(arr []Solution) {
	// Sort answers in place using insertion sort
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		i := i - 1

		for i >= 0 && len(arr[i].word) < len(key.word) {
			arr[i+1] = arr[i]
			i--
		}

		arr[i+1] = key
	}
}
