package main

import "fmt"

/**
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:
- Starting with any positive integer, replace the number by the sum of the squares of its digits.
- Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
- Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.

Example 1:
	Input: n = 19
	Output: true
	Explanation:
		1^2 + 9^2 = 82
		8^2 + 2^2 = 68
		6^2 + 8^2 = 100
		1^2 + 0^2 + 0^2 = 1

Example 2:
	Input: n = 2
	Output: false

Constraints:
	1 <= n <= 2^31 - 1
*/

func main() {
	fmt.Println(isHappy(19)) // true
	fmt.Println(isHappy(9))  // true
	// 9^2 = 81
	// 8^2 + 1^2 = 65
	// 6^2 + 5^2 = 61
	// 6^2 + 1^2 = 37
	// 3^2 + 7^2 = 58
	// 25 + 64 = 89
	// 86 + 81 = 167
	// 1 + 36 + 49 = 86
	// 64 + 36 = 100
	// 1
}

func isHappy(n int) bool {
	seen := map[int]bool{}

	for n != 1 && !seen[n] {
		seen[n] = true

		sum := 0
		for n > 0 {
			d := n % 10
			sum += d * d
			n /= 10
		}
		n = sum
	}

	return n == 1
}
