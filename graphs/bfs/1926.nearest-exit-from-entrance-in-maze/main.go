package main

import "fmt"

/**
You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
You are also given the entrance of the maze,
where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

In one step, you can move one cell up, down, left, or right.
You cannot step into a cell with a wall, and you cannot step outside the maze.
Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze.
The entrance does not count as an exit.

Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.


Example 1:
	Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
	Output: 1
	Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
	Initially, you are at the entrance cell [1,2].
	- You can reach [1,0] by moving 2 steps left.
	- You can reach [0,2] by moving 1 step up.
	It is impossible to reach [2,3] from the entrance.
	Thus, the nearest exit is [0,2], which is 1 step away.


Example 2:
	Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
	Output: 2
	Explanation: There is 1 exit in this maze at [1,2].
	[1,0] does not count as an exit since it is the entrance cell.
	Initially, you are at the entrance cell [1,0].
	- You can reach [1,2] by moving 2 steps right.
	Thus, the nearest exit is [1,2], which is 2 steps away.


Example 3:
	Input: maze = [[".","+"]], entrance = [0,0]
	Output: -1
	Explanation: There are no exits in this maze.


Constraints:
	maze.length == m
	maze[i].length == n
	1 <= m, n <= 100
	maze[i][j] is either '.' or '+'.
	entrance.length == 2
	0 <= entrancerow < m
	0 <= entrancecol < n
	entrance will always be an empty cell.
*/

func main() {
	fmt.Println(nearestExit([][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}, []int{1, 2}))
}

func nearestExit(maze [][]byte, entrance []int) int {
	mH, mW := len(maze), len(maze[0])

	visited := make([][]bool, mH)
	for i := range visited {
		visited[i] = make([]bool, mW)
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	type cell struct {
		y, x int
	}

	queue := []cell{{entrance[0], entrance[1]}}
	visited[entrance[0]][entrance[1]] = true
	steps := 0

	for len(queue) > 0 {
		steps++
		levelSize := len(queue)
		for levelSize > 0 {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				ny, nx := curr.y+d[0], curr.x+d[1]

				if ny < 0 || ny >= mH || nx < 0 || nx >= mW {
					continue
				}

				if visited[ny][nx] || maze[ny][nx] == '+' {
					continue
				}

				if ny == 0 || ny == mH-1 || nx == 0 || nx == mW-1 {
					return steps
				}

				visited[ny][nx] = true
				queue = append(queue, cell{ny, nx})
			}

			levelSize--
		}
	}

	return -1
}
