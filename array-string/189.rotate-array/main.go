package main

import "fmt"

/**
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.


Example 1:
	Input: nums = [1,2,3,4,5,6,7], k = 3
	Output: [5,6,7,1,2,3,4]
	Explanation:
	rotate 1 steps to the right: [7,1,2,3,4,5,6]
	rotate 2 steps to the right: [6,7,1,2,3,4,5]
	rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
	Input: nums = [-1,-100,3,99], k = 2
	Output: [3,99,-1,-100]
	Explanation:
	rotate 1 steps to the right: [99,-1,-100,3]
	rotate 2 steps to the right: [3,99,-1,-100]


Constraints:
	1 <= nums.length <= 10^5
	-2^31 <= nums[i] <= 2^31 - 1
	0 <= k <= 10^5


Follow up:
	- Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
	- Could you do it in-place with O(1) extra space?
*/

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)  // [5,6,7,1,2,3,4]
	rotate2([]int{1, 2, 3, 4, 5, 6, 7}, 3) // [5,6,7,1,2,3,4]
	rotate([]int{-1, -100, 3, 99}, 2)      // [3,99,-1,-100]
	rotate2([]int{-1, -100, 3, 99}, 2)     // [3,99,-1,-100]
}

func rotate(nums []int, k int) {
	k %= len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))

	fmt.Println(nums)
}

func rotate2(nums []int, k int) {
	k %= len(nums)

	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

	fmt.Println(nums)
}
