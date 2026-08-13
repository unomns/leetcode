package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(closeStrings("abc", "bca"))       // true
	fmt.Println(closeStrings("a", "aa"))          // false
	fmt.Println(closeStrings("cabbba", "abbccc")) // true
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	var m1, m2 [26]int

	for i := range len(word1) {
		m1[word1[i]-'a']++
		m2[word2[i]-'a']++
	}

	for i := range 26 {
		if (m1[i] == 0) != (m2[i] == 0) {
			return false
		}
	}

	slices.Sort(m1[:])
	slices.Sort(m2[:])

	return m1 == m2
}
