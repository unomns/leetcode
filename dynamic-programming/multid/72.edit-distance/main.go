package main

import (
	"fmt"
)

/**
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:
- Insert a character
- Delete a character
- Replace a character


Example 1:
	Input: word1 = "horse", word2 = "ros"
	Output: 3
	Explanation:
	horse -> rorse (replace 'h' with 'r')
	rorse -> rose (remove 'r')
	rose -> ros (remove 'e')

Example 2:
	Input: word1 = "intention", word2 = "execution"
	Output: 5
	Explanation:
	intention -> inention (remove 't')
	inention -> enention (replace 'i' with 'e')
	enention -> exention (replace 'n' with 'x')
	exention -> exection (replace 'n' with 'c')
	exection -> execution (insert 'u')


Constraints:
	0 <= word1.length, word2.length <= 500
	word1 and word2 consist of lowercase English letters.
*/

func main() {
	fmt.Println(minDistance("horse", "ros")) // 3
}

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)

	memo := make([][]int, l1)
	for i := range memo {
		memo[i] = make([]int, l2)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var fn func(i, j int) int

	fn = func(i, j int) int {
		if i >= l1 {
			return l2 - j // need to insert
		}
		if j >= l2 {
			return l1 - i // need to delete
		}

		if val := memo[i][j]; val >= 0 {
			return val
		}

		if word1[i] == word2[j] {
			memo[i][j] = fn(i+1, j+1)
		} else {
			memo[i][j] = 1 + min(
				fn(i, j+1),   // insert
				fn(i+1, j+1), // replace
				fn(i+1, j),   // delete
			)
		}

		return memo[i][j]
	}

	return fn(0, 0)
}
