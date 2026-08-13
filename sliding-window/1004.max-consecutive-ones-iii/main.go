package main

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))

	fmt.Println(longestOnes([]int{0, 0, 0}, 1))
	fmt.Println(longestOnes([]int{0, 0, 0}, 2))
}

func longestOnes(nums []int, k int) int {
	var max, nMax, left int

	zeroCnt := 0

	for right := range nums {
		if nums[right] == 0 {
			zeroCnt++
		}

		for zeroCnt > k {
			if nums[left] == 0 {
				zeroCnt--
			}
			left++
		}

		max = right - left + 1
		if max > nMax {
			nMax = max
		}
	}

	return nMax
}
