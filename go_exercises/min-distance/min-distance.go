package main

import (
	"sort"
)

func MinDistance(lst []float64) int {
	/*
	 * Your solution
	 */

	sort.Float64s(lst)

	var closest float64 = lst[len(lst)-1]

	for i, value := range lst[:len(lst)-1] {

		if lst[i+1]-value < closest {
			closest = lst[i+1] - value
		}

	}

	return int(closest)
}

func main() {
	/*
	 * Code written here or in other files will be ignored by the automated tests.
	 * You can run this code when running the file with `go run min-distance.go`
	 */
}
