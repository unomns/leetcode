package main

import "fmt"

/**
Given an m x n integer matrix 'matrix', if an element is 0, set its entire row and column to 0's.

You must do it in place.

Example 1:
	Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
	Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
	Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]


Constraints:
	m == matrix.length
	n == matrix[0].length
	1 <= m, n <= 200
	-2^31 <= matrix[i][j] <= 2^31 - 1

Follow up:
- A straightforward solution using O(mn) space is probably a bad idea.
- A simple improvement uses O(m + n) space, but still not the best solution.
- Could you devise a constant space solution?
*/

func main() {
	setZeroes3([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}) // [[1,0,1],[0,0,0],[1,0,1]]

	setZeroes3([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5}}) // [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	zeros := [][2]int{}

	for y := range m {
		for x := range n {
			if matrix[y][x] == 0 {
				zeros = append(zeros, [2]int{y, x})
			}
		}
	}

	for _, coords := range zeros {
		for y := range m {
			matrix[y][coords[1]] = 0
		}
		for x := range n {
			matrix[coords[0]][x] = 0
		}
	}
	fmt.Println(matrix)
}

func setZeroes2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	zeroRows := make([]bool, m)
	zeroCols := make([]bool, n)

	for y := range m {
		for x := range n {
			if matrix[y][x] == 0 {
				zeroRows[y] = true
				zeroCols[x] = true
			}
		}
	}

	for y, val := range zeroRows {
		if val {
			for x := range n {
				matrix[y][x] = 0
			}
		}
	}

	for x, val := range zeroCols {
		if val {
			for y := range m {
				matrix[y][x] = 0
			}
		}
	}

	fmt.Println(matrix)
}

func setZeroes3(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	fstRowZero, fstColZero := false, false
	for x := range n {
		if matrix[0][x] == 0 {
			fstRowZero = true
			break
		}
	}
	for y := range m {
		if matrix[y][0] == 0 {
			fstColZero = true
			break
		}
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if matrix[y][x] == 0 {
				matrix[0][x] = 0 // cols
				matrix[y][0] = 0 // rows
			}
		}
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if matrix[y][0] == 0 || matrix[0][x] == 0 {
				matrix[y][x] = 0
			}
		}
	}

	if fstColZero {
		for y := range m {
			matrix[y][0] = 0
		}
	}
	if fstRowZero {
		for x := range n {
			matrix[0][x] = 0
		}
	}

	fmt.Println(matrix)
}
