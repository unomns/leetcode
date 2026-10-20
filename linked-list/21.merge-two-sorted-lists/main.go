package main

import (
	"fmt"
	"strings"
)

/**
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Example 1:
	Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]

Example 2:
	Input: list1 = [], list2 = []
	Output: []

Example 3:
	Input: list1 = [], list2 = [0]
	Output: [0]


Constraints:
	The number of nodes in both lists is in the range [0, 50].
	-100 <= Node.val <= 100
	Both list1 and list2 are sorted in non-decreasing order.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	r := []string{}
	for l != nil {
		r = append(r, fmt.Sprint(l.Val))
		l = l.Next
	}
	return strings.Join(r, " -> ")
}

func main() {
	ex1 := func() *ListNode {
		l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
		l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

		return mergeTwoLists(l1, l2)
	}

	fmt.Println(ex1())
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}

	return head.Next
}
