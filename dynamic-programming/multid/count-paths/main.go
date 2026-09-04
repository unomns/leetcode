package main

import "fmt"

/**
Write a function, count_paths, that takes in a grid as an argument.
In the grid, 'X' epresents walls and '0' represents open spaces.
You may only move down ory to the right anc annot pass through walls.
The function should return the number of ways possible to travel from the top-left corner of the grid to the bottom-right corner.

test_00:
	grid = [
		["O","O"],
		["O","O"],
	]
	count_paths (grid) # → 2

test_01:
	grid = [
		["O","O", "X"],
		["O","O", "O"],
		["O","O", "O"],
	]
	count_paths (grid) # → 5
*/

func main() {
	fmt.Println(count_paths([][]string{
		{"O", "O"},
		{"O", "O"},
	}))

	fmt.Println(count_paths([][]string{
		{"O", "O", "X"},
		{"O", "O", "O"},
		{"O", "O", "O"},
	}))
}

func count_paths(grid [][]string) int {
	type pair struct{ y, x int }
	memo := map[pair]int{}
	m, n := len(grid), len(grid[0])
	var _count_paths func(y, x int) int

	_count_paths = func(y, x int) int {
		p := pair{y, x}
		if val, ok := memo[p]; ok {
			return val
		}

		if y == m || x == n || grid[y][x] == "X" {
			return 0
		}

		if y == m-1 && x == n-1 {
			return 1
		}

		down := _count_paths(y+1, x)
		right := _count_paths(y, x+1)

		memo[p] = down + right
		return memo[p]
	}

	return _count_paths(0, 0)
}
