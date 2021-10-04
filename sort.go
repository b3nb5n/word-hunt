package main

func maxLen(paths []string) (max int) {
	for _, path := range paths {
		if len(path) > max {
			max = len(path)
		}
	}

	return max
}

func sortByLength(paths []string) []string {
	const BASE = 3
	rng := maxLen(paths) - BASE + 1
	buckets := make([][]string, rng)

	for _, path := range paths {
		i := len(path) - BASE
		buckets[i] = append(buckets[i], path)
	}

	var i int
	for bi := rng - 1; bi >= 0; bi-- {
		for _, path := range buckets[bi] {
			paths[i] = path
			i++
		}
	}

	return paths
}
