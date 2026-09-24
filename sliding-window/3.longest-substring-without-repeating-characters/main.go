package main

import "fmt"

/**
Given a string s, find the length of the longest substring without duplicate characters.

Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:
	0 <= s.length <= 10^5
	s consists of English letters, digits, symbols and spaces.
*/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
}

func lengthOfLongestSubstring(s string) int {
	set := [256]int{}
	max, left := 0, 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if set[ch] > left {
			left = set[ch]
		}

		set[ch] = right + 1

		if right-left+1 > max {
			max = right - left + 1
		}
	}

	return max
}
