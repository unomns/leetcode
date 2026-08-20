package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
Given a root node reference of a BST and a key,
delete the node with the given key in the BST.
Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:
1. Search for a node to remove.
2. If the node is found, delete the node.

Example 1:
	Input: root = [5,3,6,2,4,null,7], key = 3
	Output: [5,4,6,2,null,null,7]
	Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
	One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
	Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

Example 2:
	Input: root = [5,3,6,2,4,null,7], key = 0
	Output: [5,3,6,2,4,null,7]
	Explanation: The tree does not contain a node with value = 0.


Example 3:
	Input: root = [], key = 0
	Output: []


Constraints:
	The number of nodes in the tree is in the range [0, 104].
	-105 <= Node.val <= 105
	Each node has a unique value.
	root is a valid binary search tree.
	-105 <= key <= 105
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{Val: 5,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 6,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(deleteNode(tree, 3))
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		min := findMin(root.Right)
		root.Val = min.Val

		root.Right = deleteNode(root.Right, min.Val)
	}

	return root
}

func findMin(n *TreeNode) *TreeNode {
	curr := n
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func (t *TreeNode) String() string {
	nums := []int{}

	stack := []*TreeNode{t}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		nums = append(nums, curr.Val)

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}

	strs := make([]string, len(nums))
	for i := range nums {
		strs[i] = strconv.Itoa(nums[i])
	}

	return strings.Join(strs, "  ")
}
