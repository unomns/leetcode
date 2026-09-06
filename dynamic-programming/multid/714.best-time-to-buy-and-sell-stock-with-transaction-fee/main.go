package main

import (
	"fmt"
	"math"
)

/**
You are given an array prices where prices[i] is the price of a given stock on the ith day,
and an integer fee representing a transaction fee.

Find the maximum profit you can achieve.
You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.

Note:
- You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
- The transaction fee is only charged once for each stock purchase and sale.


Example 1:
	Input: prices = [1,3,2,8,4,9], fee = 2
	Output: 8
	Explanation: The maximum profit can be achieved by:
	- Buying at prices[0] = 1
	- Selling at prices[3] = 8
	- Buying at prices[4] = 4
	- Selling at prices[5] = 9
	The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

Example 2:
	Input: prices = [1,3,7,5,10,3], fee = 3
	Output: 6


Constraints:
	1 <= prices.length <= 5 * 10^4
	1 <= prices[i] < 5 * 10^4
	0 <= fee < 5 * 10^4
*/

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2)) // 8
}

func maxProfit(prices []int, fee int) int {
	const (
		buy  = 0
		sell = 1
	)

	memo := make([][2]int, len(prices))
	for i := range memo {
		memo[i][0] = math.MinInt
		memo[i][1] = math.MinInt
	}

	var fn func(i, state int) int

	fn = func(i, state int) int {
		if i >= len(prices) {
			return 0
		}

		if val := memo[i][state]; val > math.MinInt {
			return val
		}

		res := 0
		switch state {
		case buy:
			res = max(fn(i+1, sell)-prices[i], fn(i+1, buy))
		case sell:
			res = max(fn(i+1, buy)+prices[i]-fee, fn(i+1, sell))
		}

		memo[i][state] = res
		return res
	}

	return fn(0, buy)
}
