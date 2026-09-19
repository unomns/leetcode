package main

import "fmt"

/**
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
	Input: s = "A man, a plan, a canal: Panama"
	Output: true
	Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
	Input: s = "race a car"
	Output: false
	Explanation: "raceacar" is not a palindrome.

Example 3:
	Input: s = " "
	Output: true
	Explanation: s is an empty string "" after removing non-alphanumeric characters.
	Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:
	1 <= s.length <= 2 * 10^5
	s consists only of printable ASCII characters.
*/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}

func isPalindrome(s string) bool {
	cl := []rune{}
	for _, c := range s {
		if (c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			if c >= 65 && c <= 90 {
				c += 32
			}
			cl = append(cl, c)
		}
	}

	left, right := 0, len(cl)-1
	for left < right {
		if cl[left] != cl[right] {
			return false
		}
		left++
		right--
	}
	return true
}
