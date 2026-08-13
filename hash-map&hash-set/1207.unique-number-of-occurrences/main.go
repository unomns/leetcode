package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})) // true - the value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences

	fmt.Println(uniqueOccurrences([]int{1, 2}))                             // false
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})) // true
}

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}

	for _, n := range arr {
		m[n]++
	}

	s := map[int]struct{}{}

	for _, c := range m {
		if _, ok := s[c]; ok {
			return false
		}

		s[c] = struct{}{}
	}

	return true
}
