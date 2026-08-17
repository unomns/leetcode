package main

import (
	"context"
)

func similar(root1 *TreeNode, root2 *TreeNode) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer close(ch1)
		findLeafs(root1, ch1, ctx)
	}()

	go func() {
		defer close(ch2)
		findLeafs(root2, ch2, ctx)
	}()

	for {
		r1, ok1 := <-ch1
		r2, ok2 := <-ch2

		if !ok1 && !ok2 {
			return true
		} else if !ok1 || !ok2 {
			return false
		} else if r1 != r2 {
			return false
		}
	}
}

func findLeafs(node *TreeNode, ch chan int, ctx context.Context) {
	if node == nil {
		return
	}

	select {
	case <-ctx.Done():
		return
	default:
		if node.Left == nil && node.Right == nil {
			select {
			case ch <- node.Val:
			case <-ctx.Done():
			}
			return
		}

		findLeafs(node.Left, ch, ctx)
		findLeafs(node.Right, ch, ctx)
	}
}
