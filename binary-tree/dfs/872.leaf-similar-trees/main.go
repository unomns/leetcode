package main

import (
	"fmt"
	"slices"
)

/**
Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar if their leaf value sequence is the same.
Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.


Example 1:
	Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
	Output: true

Example 2:
	Input: root1 = [1,2,3], root2 = [1,3,2]
	Output: false


Constraints:

	The number of nodes in each tree will be in the range [1, 200].
	Both of the given trees will have values in the range [0, 200].
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	r2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
	}
	fmt.Println("[recursive] leafSimilar: ", leafSimilar(r1, r2))
	fmt.Println("[goroutine] similar: ", similar(r1, r2))
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1Leafs := []int{}
	r2Leafs := []int{}

	findLastLeafs(root1, &r1Leafs)
	findLastLeafs(root2, &r2Leafs)

	return slices.Equal(r1Leafs, r2Leafs)
}

func findLastLeafs(root *TreeNode, container *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*container = append(*container, root.Val)
		return
	}

	findLastLeafs(root.Left, container)
	findLastLeafs(root.Right, container)
}
