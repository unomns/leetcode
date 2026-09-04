package main

import "fmt"

/**
You have two types of tiles: a 2 x 1 domino shape and a tromino shape (L). You may rotate these shapes.

Given an integer n, return the number of ways to tile an 2 x n board.
Since the answer may be very large, return it modulo 10**9 + 7.

In a tiling, every square must be covered by a tile.
Two tilings are different if and only if there are two 4-directionally adjacent cells on the board
such that exactly one of the tilings has both squares occupied by a tile.


Example 1:
	Input: n = 3
	Output: 5
	Explanation: The five different ways are shown above.


Example 2:
	Input: n = 1
	Output: 1


Constraints:
	1 <= n <= 1000
*/

func main() {
	fmt.Println(numTilings(3)) //5
	fmt.Println(numTilings(1)) //1
}

func numTilings(n int) int {
	const MOD = 1_000_000_007
	const (
		full = iota
		top
		bot
	)

	type pair struct {
		i, state int
	}

	memo := map[pair]int{}

	var fn func(i, state int) int
	add := func(a, b int) int {
		return (a + b) % MOD
	}

	fn = func(i, state int) int {
		if i > n {
			return 0
		}

		pair := pair{i, state}
		if val, ok := memo[pair]; ok {
			return val
		}

		if i == n {
			if state == full {
				return 1
			}
			return 0
		}

		ways := 0
		switch state {
		case full:
			// vert
			ways = add(ways, fn(i+1, full))
			// horiz
			ways = add(ways, fn(i+2, full))
			// L top
			ways = add(ways, fn(i+1, top))
			// L bot
			ways = add(ways, fn(i+1, bot))

		case top:
			// horiz
			ways = add(ways, fn(i+1, bot))
			// L bot
			ways = add(ways, fn(i+2, full))

		case bot:
			// horiz
			ways = add(ways, fn(i+1, top))
			// L top
			ways = add(ways, fn(i+2, full))
		}

		memo[pair] = ways

		return ways
	}

	return fn(0, full)
}
