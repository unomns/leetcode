package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalPairs([][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}))

	fmt.Println(equalPairs([][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}))
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	cnt := 0

	rows := map[[200]int]int{}

	for y := range n {
		row := [200]int{}
		copy(row[:], grid[y])
		rows[row]++
	}

	for x := range n {
		col := [200]int{}
		for y := range n {
			col[y] = grid[y][x]
		}

		cnt += rows[col]
	}

	return cnt
}
