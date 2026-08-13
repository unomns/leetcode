package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                // 3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) // 5
}

func longestSubarray(nums []int) int {
	zeros := 0

	var left, max, nMax int

	for right := range nums {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		max = right - left + 1
		if max > nMax {
			nMax = max
		}
	}

	return nMax - 1
}
