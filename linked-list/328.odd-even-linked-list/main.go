package main

import (
	"fmt"
	"strconv"
)

/**
Given the head of a singly linked list,
group all the nodes with odd indices together followed by the nodes with even indices,
and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.


Example 1:
	Input: head = [1,2,3,4,5]
	Output: [1,3,5,2,4]

Example 2:
	Input: head = [2,1,3,5,6,4,7]
	Output: [2,3,6,7,1,5,4]


Constraints:
	The number of nodes in the linked list is in the range [0, 104].
	-106 <= Node.val <= 106
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{
		Val: 2, Next: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 5, Next: &ListNode{
						Val: 6, Next: &ListNode{
							Val: 4, Next: &ListNode{
								Val: 7, Next: nil}}}}}}}
	// 2,3,6,7,1,5,4

	fmt.Println("before: ", n)
	oddEvenList(n)
	fmt.Println("after: ", n)

	// 1,2,3,4,5
	n = &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: nil}}}}}
	// 1,3,5,2,4

	fmt.Println("before: ", n)
	oddEvenList(n)
	fmt.Println("after: ", n)

	n = &ListNode{Val: 1, Next: nil}
	// 1

	fmt.Println("before: ", n)
	oddEvenList(n)
	fmt.Println("after: ", n)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	eHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = eHead

	return head
}

func (n *ListNode) String() string {
	output := 0
	for n != nil {
		output = output*10 + n.Val
		n = n.Next
	}
	return strconv.Itoa(output)
}
