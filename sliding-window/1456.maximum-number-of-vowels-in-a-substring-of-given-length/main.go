package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
}

func maxVowels(s string, k int) int {
	max := 0
	for i := range k {
		if isVowel(s[i]) {
			max++
		}
	}

	c := max
	for i := k; i < len(s); i++ {
		if max == k {
			return max
		}

		if isVowel(s[i-k]) {
			c--
		}

		if isVowel(s[i]) {
			c++
		}

		if c > max {
			max = c
		}
	}

	return max
}

func isVowel(b byte) bool {
	switch rune(b) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
