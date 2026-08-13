package main

import "fmt"

func main() {
	/**
	  NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
	  numArray.sumRange(0, 2); // return (-2) + 0 + 3 = 1
	  numArray.sumRange(2, 5); // return 3 + (-5) + 2 + (-1) = -1
	  numArray.sumRange(0, 5); // return (-2) + 0 + 3 + (-5) + 2 + (-1) = -3
	*/

	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2)) // 1
	fmt.Println(numArray.SumRange(2, 5)) // -1
	fmt.Println(numArray.SumRange(0, 5)) // -3

}

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	pref := make([]int, len(nums))

	s := 0
	for i, n := range nums {
		s += n
		pref[i] = s
	}

	return NumArray{nums: pref}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}
	return this.nums[right] - this.nums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
