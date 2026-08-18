package main

import "fmt"

/**
You are given the root of a binary tree.

A ZigZag path for a binary tree is defined as follow:
1. Choose any node in the binary tree and a direction (right or left).
2. If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
3. Change the direction from right to left or from left to right.
4. Repeat the second and third steps until you can't move in the tree.

Zigzag length is defined as the number of nodes visited - 1.
(A single node has a length of 0).

Return the longest ZigZag path contained in that tree.


Example 1:
	Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
	Output: 3
	Explanation: Longest ZigZag path in blue nodes (right -> left -> right).

Example 2:
	Input: root = [1,1,1,null,1,null,null,1,1,null,1]
	Output: 4
	Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).

Example 3:
	Input: root = [1]
	Output: 0


Constraints:
	The number of nodes in the tree is in the range [1, 5 * 104].
	1 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	btree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(longestZigZag(btree))
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var max int
	const left = -1
	const right = 1

	var dfs func(node *TreeNode, dir, length int)

	dfs = func(node *TreeNode, dir, length int) {
		if node == nil {
			return
		}

		if length > max {
			max = length
		}

		if dir == left {
			dfs(node.Right, right, length+1)
			dfs(node.Left, left, 1)
		} else {
			dfs(node.Left, left, length+1)
			dfs(node.Right, right, 1)
		}
	}

	dfs(root.Left, left, 1)
	dfs(root.Right, right, 1)

	return max
}
