package main

import "fmt"

/**
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.

Example 1:
	Input: head = [3,2,0,-4], pos = 1
	Output: true
	Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

Example 2:
	Input: head = [1,2], pos = 0
	Output: true
	Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

Example 3:
	Input: head = [1], pos = -1
	Output: false
	Explanation: There is no cycle in the linked list.


Constraints:
	The number of the nodes in the list is in the range [0, 10^4].
	-10^5 <= Node.val <= 10^5
	pos is -1 or a valid index in the linked-list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ex1 := func() *ListNode {
		n0 := ListNode{Val: 3}
		n1 := ListNode{Val: 2}
		n2 := ListNode{Val: 0}
		n3 := ListNode{Val: -4}

		n0.Next = &n1
		n1.Next = &n2
		n2.Next = &n3
		n3.Next = &n2

		return &n0
	}

	ex2 := func() *ListNode {
		n0 := ListNode{Val: 1}
		n1 := ListNode{Val: 2}

		n0.Next = &n1
		n1.Next = &n0

		return &n0
	}

	ex3 := func() *ListNode {
		return &ListNode{Val: 1}
	}

	fmt.Println(hasCycle(ex1()))
	fmt.Println(hasCycle(ex2()))
	fmt.Println(hasCycle(ex3()))
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fst, snd := head, head.Next

	for fst != snd && snd != nil && snd.Next != nil {
		fst = fst.Next
		snd = snd.Next.Next
	}

	return fst == snd
}
