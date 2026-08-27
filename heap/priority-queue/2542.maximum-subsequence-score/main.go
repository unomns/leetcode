package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/**
You are given two 0-indexed integer arrays nums1 and nums2 of equal length n and a positive integer k.
You must choose a subsequence of indices from nums1 of length k.

For chosen indices i0, i1, ..., ik - 1, your score is defined as:
- The sum of the selected elements from nums1 multiplied with the minimum of the selected elements from nums2.
- It can defined simply as: (nums1[i0] + nums1[i1] +...+ nums1[ik - 1]) * min(nums2[i0] , nums2[i1], ... ,nums2[ik - 1]).

Return the maximum possible score.

A subsequence of indices of an array is a set that can be derived from the set {0, 1, ..., n-1} by deleting some or no elements.


Example 1:
	Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3
	Output: 12
	Explanation:
	The four possible subsequence scores are:
	- We choose the indices 0, 1, and 2 with score = (1+3+3) * min(2,1,3) = 7.
	- We choose the indices 0, 1, and 3 with score = (1+3+2) * min(2,1,4) = 6.
	- We choose the indices 0, 2, and 3 with score = (1+3+2) * min(2,3,4) = 12.
	- We choose the indices 1, 2, and 3 with score = (3+3+2) * min(1,3,4) = 8.
	Therefore, we return the max score, which is 12.

Example 2:
	Input: nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1
	Output: 30
	Explanation:
	Choosing index 2 is optimal: nums1[2] * nums2[2] = 3 * 10 = 30 is the maximum possible score.


Constraints:
	n == nums1.length == nums2.length
	1 <= n <= 105
	0 <= nums1[i], nums2[j] <= 105
	1 <= k <= n
*/

func main() {
	fmt.Println(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))        // 12
	fmt.Println(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1)) // 30
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	n := len(*h)
	l := (*h)[n-1]
	*h = (*h)[:n-1]
	return l
}

var _ heap.Interface = (*MinHeap)(nil)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	type pair struct {
		val, min int
	}

	pairs := make([]pair, len(nums1))
	for i := range nums1 {
		pairs[i] = pair{nums1[i], nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].min > pairs[j].min
	})

	h := &MinHeap{}
	heap.Init(h)

	sum, max := 0, 0
	for _, p := range pairs {
		heap.Push(h, p.val)
		sum += p.val

		if h.Len() > k {
			m := heap.Pop(h).(int)
			sum -= m
		}

		if h.Len() == k {
			score := sum * p.min
			if score > max {
				max = score
			}
		}
	}

	return int64(max)
}
