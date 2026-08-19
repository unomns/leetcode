package main

import "fmt"

/*
Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example 1:
	Input: root = [1,2,3,null,5,null,4]
	Output: [1,3,4]

Example 2:
	Input: root = [1,2,3,4,null,null,null,5]
	Output: [1,3,4,5]

Example 3:
	Input: root = [1,null,3]
	Output: [1,3]

Example 4:
	Input: root = []
	Output: []


Constraints:
	The number of nodes in the tree is in the range [0, 100].
	-100 <= Node.val <= 100
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	/*
			  1
			/   \
		   2     3
			\     \
			 5     4

			   output: 1, 3, 4
	*/
	fmt.Println(rightSideView(tree))
	fmt.Println(rightSideViewRecursive(tree))
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nums := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := range levelSize {
			curr := queue[0]
			queue = queue[1:]

			if i == 0 {
				nums = append(nums, curr.Val)
			}

			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
		}
	}

	return nums
}

func rightSideViewRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nums := []int{}
	var dfs func(node *TreeNode, d int)

	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}

		if d == len(nums) {
			nums = append(nums, node.Val)
		}

		dfs(node.Right, d+1)
		dfs(node.Left, d+1)
	}

	dfs(root, 0)

	return nums
}
