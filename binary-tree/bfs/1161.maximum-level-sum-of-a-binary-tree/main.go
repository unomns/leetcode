package main

import "fmt"

/**
Given the root of a binary tree, the level of its root is 1,
the level of its children is 2, and so on.

Return the smallest level x
such that the sum of all the values of nodes at level x is maximal.

Example 1:
	Input: root = [1,7,0,7,-8,null,null]
	Output: 2
	Explanation:
	Level 1 sum = 1.
	Level 2 sum = 7 + 0 = 7.
	Level 3 sum = 7 + -8 = -1.
	So we return the level with the maximum sum which is level 2.

Example 2:
	Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
	Output: 2


Constraints:
	The number of nodes in the tree is in the range [1, 104].
	-105 <= Node.val <= 105
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -8},
		},
		Right: &TreeNode{Val: 0},
	}

	tree2 := &TreeNode{Val: 989,
		Right: &TreeNode{Val: 10250,
			Left: &TreeNode{Val: 98693},
			Right: &TreeNode{Val: -89388,
				Right: &TreeNode{Val: -32127},
			},
		},
	}

	fmt.Println(maxLevelSum(tree1))
	fmt.Println(maxLevelSum(tree2))
}

func maxLevelSum(root *TreeNode) int {
	deep := 0
	max := root.Val
	min := 1

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		deep++
		levelSize := len(queue)
		sum := 0
		for levelSize > 0 {
			curr := queue[0]
			queue = queue[1:]

			sum += curr.Val

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			levelSize--
		}
		if sum > max {
			max = sum
			min = deep
		}
	}

	return min
}
