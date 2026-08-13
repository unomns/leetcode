package main

import "fmt"

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                 // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0
}

func minSubArrayLen(target int, nums []int) int {
	var size, sum, left int
	min := len(nums) + 1

	for right := range nums {
		sum += nums[right]
		size++

		for sum > target {
			sum -= nums[left]
			left++
			size--
		}

		if sum == target && size < min {
			min = size
		}
	}

	if min > len(nums) {
		return 0
	}

	return min
}
