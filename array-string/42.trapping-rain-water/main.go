package main

import "fmt"

/**
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

Example 1:
	Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6
	Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

Example 2:
	Input: height = [4,2,0,3,2,5]
	Output: 9

Constraints:
	n == height.length
	1 <= n <= 2 * 10^4
	0 <= height[i] <= 10^5
*/

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
	fmt.Println(trap([]int{4, 2, 3}))                            // 1
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	water := 0
	n := len(height)

	peak := 0
	for i := range n {
		if height[i] > height[peak] {
			peak = i
		}
	}

	left := 0

	for left < peak {
		if height[left] == 0 {
			left++
			continue
		}

		right := left + 1

		for right <= peak && height[right] < height[left] {
			right++
		}

		max := height[left]
		for _, v := range height[left+1 : right] {
			if max-v > 0 {
				water += max - v
			}
		}

		left = right
	}

	right := n - 1
	for right > peak {
		if height[right] == 0 {
			right--
			continue
		}

		left := right - 1
		for left >= peak && height[left] < height[right] {
			left--
		}

		max := height[right]
		for _, v := range height[left+1 : right] {
			if max-v > 0 {
				water += max - v
			}
		}

		right = left
	}

	return water
}
