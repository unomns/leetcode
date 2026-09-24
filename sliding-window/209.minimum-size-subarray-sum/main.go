package main

import (
	"fmt"
	"math"
)

/**
Given an array of positive integers nums and a positive integer target,
return the minimal length of a subarray whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.

Example 1:
	Input: target = 7, nums = [2,3,1,2,4,3]
	Output: 2
	Explanation: The subarray [4,3] has the minimal length under the problem constraint.

Example 2:
	Input: target = 4, nums = [1,4,4]
	Output: 1

Example 3:
	Input: target = 11, nums = [1,1,1,1,1,1,1,1]
	Output: 0


Constraints:
	1 <= target <= 10^9
	1 <= nums.length <= 10^5
	1 <= nums[i] <= 10^4

Follow up: If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log(n)).
*/

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
	fmt.Println(minSubArrayLen(7, []int{2, 1, 5, 2}))              // 0
}

func minSubArrayLen(target int, nums []int) int {
	min := math.MaxInt
	left, right, s := 0, 0, 0

	for right < len(nums) && left <= right {
		s += nums[right]
		for s >= target {
			if right-left+1 < min {
				min = right - left + 1
			}
			s -= nums[left]
			left++
		}
		right++
	}

	if min == math.MaxInt {
		return 0
	}
	return min
}
