package main

import (
	"container/heap"
	"fmt"
)

/**
You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

Implement the SmallestInfiniteSet class:
- SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
- int popSmallest() Removes and returns the smallest integer contained in the infinite set.
- void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.


Example 1:
	Input
	["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
	[[], [2], [], [], [], [1], [], [], []]
	Output
	[null, null, 1, 2, 3, null, 1, 4, 5]

	Explanation
	SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
	smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
	smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
	smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
	smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
	smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
	smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
									   // is the smallest number, and remove it from the set.
	smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
	smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.


Constraints:
	1 <= num <= 1000
	At most 1000 calls will be made in total to popSmallest and addBack.
*/

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
func main() {
	params := []int{}
	obj := Constructor()
	obj.AddBack(2)
	for range 3 {
		params = append(params, obj.PopSmallest())
	}
	obj.AddBack(1)
	for range 3 {
		params = append(params, obj.PopSmallest())
	}

	fmt.Println(params)
}

type MinHeap []int

var _ heap.Interface = (*MinHeap)(nil)

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	n := len(*h)
	l := (*h)[n-1]
	(*h) = (*h)[:n-1]
	return l
}

type SmallestInfiniteSet struct {
	next   int
	h      *MinHeap
	inHeap map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := &MinHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{1, h, make(map[int]bool)}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.h.Len() > 0 {
		n := heap.Pop(this.h).(int)

		if n < this.next {
			delete(this.inHeap, n)
			return n
		}

		heap.Push(this.h, n)
	}
	this.next++
	return this.next - 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.next || this.inHeap[num] {
		return
	}

	heap.Push(this.h, num)
	this.inHeap[num] = true
}
