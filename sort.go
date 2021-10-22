package main

func dedupe(answers []Answer) (res []Answer) {
	keys := make(map[string]bool)
	for _, answer := range answers {
		if _, found := keys[answer.word]; !found {
			keys[answer.word] = true
			res = append(res, answer)
		}
	}

	return res
}

func insertionSort(arr []Answer) {
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
