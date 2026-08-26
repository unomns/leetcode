package main

import "fmt"

/**
You are given an m x n grid where each cell can have one of three values:
	0 representing an empty cell,
	1 representing a fresh orange, or
	2 representing a rotten orange.

Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.
If this is impossible, return -1

Example 1:
	Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
	Output: 4

Example 2:
	Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
	Output: -1
	Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.

Example 3:
	Input: grid = [[0,2]]
	Output: 0
	Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.


Constraints:
	m == grid.length
	n == grid[i].length
	1 <= m, n <= 10
	grid[i][j] is 0, 1, or 2.
*/

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})) // 4
}

func orangesRotting(grid [][]int) int {
	const (
		fresh  = 1
		rotten = 2
	)

	type orange struct {
		y, x int
	}

	m := len(grid)
	n := len(grid[0])

	queue := []orange{}
	freshCnt := 0

	for y, arr := range grid {
		for x, val := range arr {
			switch val {
			case rotten:
				queue = append(queue, orange{y, x})
			case fresh:
				freshCnt++
			}
		}
	}

	if freshCnt == 0 {
		return 0
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ticks := 0
	for len(queue) > 0 && freshCnt > 0 {
		ticks++
		sizeLevel := len(queue)
		for sizeLevel > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, d := range dirs {
				ny, nx := curr.y+d[0], curr.x+d[1]

				if ny < 0 || ny >= m || nx < 0 || nx >= n {
					continue
				}
				if grid[ny][nx] != fresh {
					continue
				}

				freshCnt--
				grid[ny][nx] = rotten
				queue = append(queue, orange{ny, nx})
			}

			sizeLevel--
		}
	}

	if freshCnt == 0 {
		return ticks
	}
	return -1
}
