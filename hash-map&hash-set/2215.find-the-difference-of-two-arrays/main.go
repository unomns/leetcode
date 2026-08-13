package main

import "fmt"

func main() {
	fmt.Println(findDifference(
		[]int{1, 2, 3, 1},
		[]int{2, 4, 6},
	))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)

	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, n := range nums1 {
		set1[n] = struct{}{}
	}

	for _, n := range nums2 {
		set2[n] = struct{}{}
	}

	for n := range set1 {
		if _, ok := set2[n]; !ok {
			res[0] = append(res[0], n)
		}
	}
	for n := range set2 {
		if _, ok := set1[n]; !ok {
			res[1] = append(res[1], n)
		}
	}

	return res
}
