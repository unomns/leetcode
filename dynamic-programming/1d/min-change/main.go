package main

import (
	"fmt"
	"math"
)

/**
Write a function min_change that takes in an amount and a list of coins.
The function should return the minimum number of coins required to create the amount.
You may use each coin as many times as necessary.

If it is not possible to create the amount, then return -1.

test_00:
	min_change(8, [1, 5, 4, 12]$ # → 2, because 4*4 is the minimum coins possible

test_01:
	min_change (13, [1, 9, 5, 14, 301) # → 5

test_02:
	min_change (23, [2, 5, 71) # → 4
*/

func main() {
	fmt.Println(min_change(8, []int{1, 5, 4, 12}))       // 2
	fmt.Println(min_change(13, []int{1, 9, 5, 14, 301})) // 5
}

func min_change(amount int, coins []int) int {
	max := math.MaxInt32
	var _min_change func(a int, memo map[int]int) int

	_min_change = func(a int, memo map[int]int) int {
		if val, ok := memo[a]; ok {
			return val
		}

		if a == 0 {
			return 0
		}

		if a < 0 {
			return max
		}

		min := max
		for _, c := range coins {
			m := 1 + _min_change(a-c, memo)
			if m < min {
				min = m
			}
		}

		memo[a] = min

		return min
	}

	res := _min_change(amount, make(map[int]int))
	if res == max {
		return -1
	}
	return res
}
