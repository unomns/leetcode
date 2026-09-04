package main

import "fmt"

/**
Write a function sum_possible that takes in an amount and a list of positive numbers.
The function should return a boolean indicating whether or not it is possible to create the amount by summing numbers of the list.
You may reuse numbers of the list as many times as necessary.

You may assume that the target amount is non-negative.

test_00:
	sum_possible (8, [5, 12, 4]) # → True, 4 + 4

test_01:
	sum_possible (15, [6, 2, 10, 19]) # → False

test_02:
	sum_possible (13, [6, 2, 1]) # → True
*/

func main() {
	fmt.Println(sum_possible(8, []int{5, 12, 4}))      // true
	fmt.Println(sum_possible(15, []int{6, 2, 10, 19})) // false
}

func sum_possible(n int, nums []int) bool {
	var _sum_possible func(s int, memo map[int]bool) bool

	_sum_possible = func(s int, memo map[int]bool) bool {
		if val, ok := memo[s]; ok {
			return val
		}
		if s == 0 {
			return true
		}

		if s < 0 {
			return false
		}

		for i := range nums {
			if _sum_possible(s-nums[i], memo) {
				memo[s] = true
				return true
			}
		}

		memo[s] = false
		return false
	}

	return _sum_possible(n, make(map[int]bool))
}
