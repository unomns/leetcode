package main

import "fmt"

/**
Given the root of a binary tree and an integer targetSum,
return the number of paths where the sum of the values along the path equals targetSum.

The path does not need to start or end at the root or a leaf,
but it must go downwards (i.e., traveling only from parent nodes to child nodes).

Example 1:
	Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
	Output: 3
	Explanation: The paths that sum to 8 are shown.

Example 2:
	Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
	Output: 3


Constraints:
	The number of nodes in the tree is in the range [0, 1000].
	-109 <= Node.val <= 109
	-1000 <= targetSum <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: -2,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}

	fmt.Println(pathSum(t, 8))
}

// O(n**2) complexity
func pathSum(root *TreeNode, targetSum int) int {
	var counter int
	var dpsCount func(node *TreeNode, sum int)
	var dps func(node *TreeNode, fn func(node *TreeNode, sum int))

	dpsCount = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum += node.Val
		if sum == targetSum {
			counter++
		}

		dpsCount(node.Left, sum)
		dpsCount(node.Right, sum)
	}

	dps = func(node *TreeNode, fn func(node *TreeNode, sum int)) {
		if node == nil {
			return
		}

		fn(node, 0)

		dps(node.Left, fn)
		dps(node.Right, fn)
	}

	dps(root, dpsCount)

	return counter
}

// todo: solve problem with O(n)
func pathSumWithPrefix(root *TreeNode, targetSum int) int {
	return 0
}
