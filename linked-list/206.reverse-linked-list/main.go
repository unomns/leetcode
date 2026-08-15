package main

import (
	"fmt"
	"strconv"
)

/**

Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
	Input: head = []
	Output: []

Constraints:
	The number of nodes in the list is the range [0, 5000].
	-5000 <= Node.val <= 5000
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: nil}}}}}

	fmt.Println("before: ", n)
	n = reverseList(n)
	fmt.Println("after: ", n)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	curr := head

	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

func (n *ListNode) String() string {
	output := 0
	for n != nil {
		output = output*10 + n.Val
		n = n.Next
	}
	return strconv.Itoa(output)
}
