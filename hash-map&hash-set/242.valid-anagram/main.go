package main

import "fmt"

/**
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true

Example 2:
	Input: s = "rat", t = "car"
	Output: false


Constraints:
	1 <= s.length, t.length <= 5 * 10^4
	s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counts := [26]int{}

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, cnt := range counts {
		if cnt != 0 {
			return false
		}
	}
	return true
}
