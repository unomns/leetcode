package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6})) // i: 3
	fmt.Println(pivotIndex([]int{2, 1, -1}))         // i: 0
	fmt.Println(pivotIndex([]int{1, -1, 2}))         // i: 2
}

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}

	leftSum := 0
	for i, n := range nums {
		if totalSum-leftSum-n == leftSum {
			return i
		}

		leftSum += n
	}

	return -1
}
