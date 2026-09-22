package main

import (
	"fmt"
	"sort"
)

/**
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
	Input: s = "abc", t = "ahbgdc"
	Output: true

Example 2:
	Input: s = "axc", t = "ahbgdc"
	Output: false


Constraints:
	0 <= s.length <= 100
	0 <= t.length <= 10^4
	s and t consist only of lowercase English letters.


Follow up:
	Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence.
	In this scenario, how would you change your code?
*/

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false

	f := newFollowUp("ahbgdc")
	fmt.Println(f.isSubsequence("abc")) // true
	fmt.Println(f.isSubsequence("axc")) // false
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	found := 0
	for i := 0; i < len(t); i++ {
		if s[found] == t[i] {
			found++
			if found == len(s) {
				return true
			}
		}
	}

	return false
}

type followUp struct {
	pos [26][]int
}

func newFollowUp(t string) followUp {
	pos := [26][]int{}
	for i := range t {
		idx := t[i] - 'a'
		pos[idx] = append(pos[idx], i)
	}
	return followUp{pos}
}

func (f followUp) isSubsequence(s string) bool {
	prev := -1

	for i := range s {
		idx := s[i] - 'a'

		indices := f.pos[idx]
		if len(indices) == 0 {
			return false
		}

		iPos := sort.Search(len(indices), func(j int) bool {
			return indices[j] > prev
		})

		if iPos == len(indices) {
			return false
		}

		prev = indices[iPos]
	}

	return true
}
