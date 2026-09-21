package main

import "fmt"

/**
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
	Input: haystack = "sadbutsad", needle = "sad"
	Output: 0
	Explanation: "sad" occurs at index 0 and 6.
	The first occurrence is at index 0, so we return 0.

Example 2:
	Input: haystack = "leetcode", needle = "leeto"
	Output: -1
	Explanation: "leeto" did not occur in "leetcode", so we return -1.


Constraints:
	1 <= haystack.length, needle.length <= 104
	haystack and needle consist of only lowercase English characters.
*/

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))  // 0
	fmt.Println(strStr("leetcode", "leeto")) // -1
}

func strStr(haystack string, needle string) int {
	n := len(haystack)
	nn := len(needle)

	for i := 0; i+nn <= n; i++ {
		if haystack[i:i+nn] == needle {
			return i
		}
	}

	return -1
}
