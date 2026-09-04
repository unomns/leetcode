package main

import "fmt"

/**
There is a robot on an m x n grid.
The robot is initially located at the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]).
The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.


Example 1:
	Input: m = 3, n = 7
	Output: 28

Example 2:
	Input: m = 3, n = 2
	Output: 3
	Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
	1. Right -> Down -> Down
	2. Down -> Down -> Right
	3. Down -> Right -> Down


Constraints:
	1 <= m, n <= 100
*/

func main() {
	fmt.Println(uniquePaths(3, 7)) // 28
	fmt.Println(uniquePaths(3, 2)) // 3
}

func uniquePaths(m int, n int) int {
	type pair struct{ y, x int }
	memo := map[pair]int{}

	var fn func(y, x int) int

	fn = func(y, x int) int {
		if y >= m || x >= n {
			return 0
		}

		p := pair{y, x}
		if val, ok := memo[p]; ok {
			return val
		}

		if y == m-1 && x == n-1 {
			return 1
		}

		memo[p] = fn(y+1, x) + fn(y, x+1)
		return memo[p]
	}

	return fn(0, 0)
}
