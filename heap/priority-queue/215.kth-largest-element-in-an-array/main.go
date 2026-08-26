package main

import (
	"container/heap"
	"fmt"
)

/**
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.
Can you solve it without sorting?


Example 1:
	Input: nums = [3,2,1,5,6,4], k = 2
	Output: 5

Example 2:
	Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
	Output: 4


Constraints:
	1 <= k <= nums.length <= 105
	-104 <= nums[i] <= 104
*/

type MinHeap []int

var _ heap.Interface = (*MinHeap)(nil)

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	n := len(*h)
	l := (*h)[n-1]
	*h = (*h)[:n-1]
	return l
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))          // 5
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // 4
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, n := range nums {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}
