package main

import (
	"fmt"
	"strings"
)

/**
Given an input string s, reverse the order of the words.

A word is defined as a sequence of non-space characters.
The words in s will be separated by at least one space.

Return a string of the words in reverse order concatenated by a single space.

Note that s may contain leading or trailing spaces or multiple spaces between two words.
The returned string should only have a single space separating the words.
Do not include any extra spaces.

Example 1:
	Input: s = "the sky is blue"
	Output: "blue is sky the"

Example 2:
	Input: s = "  hello world  "
	Output: "world hello"
	Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
	Input: s = "a good   example"
	Output: "example good a"
	Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.


Constraints:
	1 <= s.length <= 10^4
	s contains English letters (upper-case and lower-case), digits, and spaces ' '.
	There is at least one word in s.

Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?
*/

func main() {
	fmt.Println(reverseWords("the sky is blue"))   // blue is sky the
	fmt.Println(reverseWords("a good    example")) // good example a
	fmt.Println(reverseWords("the sky"))           // good example a
	fmt.Println(reverseWords("a "))                // good example a
}

func reverseWords(s string) string {
	buf := []string{}

	wl := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			wl++
		}

		if wl > 0 && (s[i] == ' ' || i == 0) {
			start := i
			if s[i] == ' ' {
				start++
			}
			buf = append(buf, s[start:start+wl])
			wl = 0
		}
	}

	return strings.Join(buf, " ")
}
