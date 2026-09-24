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
	if len(s) <= 1 {
		return len(s)
	}

	max, left, right := 0, 0, 1

	set := [256]bool{}
	set[s[left]] = true

	for left < right && right < len(s) {
		if set[s[right]] {
			if right-left > max {
				max = right - left
			}
			for left < right && s[left] != s[right] {
				set[s[right]] = false
				left++
			}
			set[s[right]] = false
			left++
		}

		set[s[right]] = true
		right++
		if right-left > max {
			max = right - left
		}
	}

	return max
}
