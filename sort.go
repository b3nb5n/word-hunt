package main

import (
	"fmt"
	"time"
)

func sortByLength(src []string) []string {
	start := time.Now()

	var buckets [10][]string

	for i := 0; i < len(src); i++ {
		length := len(src[i])
		buckets[length] = append(buckets[length], src[i])
	}

	src = src[0:0]
	bI := 9

	for bI > 0 {
		bucket := buckets[bI]

		if len(bucket) > 0 {
			src = append(src, bucket[len(bucket)-1])
			buckets[bI][len(bucket)-1] = ""
			buckets[bI] = buckets[bI][:len(bucket)-1]
		} else {
			buckets[bI] = nil
			bI--
		}
	}

	fmt.Printf("sort words: %v\n", time.Since(start))
	return src
}
