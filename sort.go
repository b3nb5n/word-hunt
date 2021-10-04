package main

import (
	"fmt"
	"time"
)

func sortByLength(src []string) []string {
	start := time.Now()
	var buckets [10][]string

	for i := 0; i < len(src); i++ {
		length := len(src[i]) - 1
		buckets[length] = append(buckets[length], src[i])
	}

	src = src[0:0]

	i := 9
	for i > 0 {
		length := len(buckets[i])

		if length > 0 {
			src = append(src, buckets[i][length-1])
			buckets[i][length-1] = ""
			buckets[i] = buckets[i][:length-1]
		} else {
			buckets[i] = nil
			i--
		}
	}

	fmt.Printf("sorted %v strings in %v\n", len(src), time.Since(start))
	return src
}
