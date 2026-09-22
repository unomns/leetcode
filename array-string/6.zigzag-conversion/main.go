package main

import "fmt"

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
	P   A   H   N
	A P L S I I G
	Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:
	string convert(string s, int numRows);

Example 1:
	Input: s = "PAYPALISHIRING", numRows = 3
	Output: "PAHNAPLSIIGYIR"

Example 2:
	Input: s = "PAYPALISHIRING", numRows = 4
	Output: "PINALSIGYAHRPI"
	Explanation:
	P     I    N
	A   L S  I G
	Y A   H R
	P     I

Example 3:
	Input: s = "A", numRows = 1
	Output: "A"

Constraints:
	1 <= s.length <= 1000
	s consists of English letters (lower-case and upper-case), ',' and '.'.
	1 <= numRows <= 1000
*/

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3)) // PAHNAPLSIIGYIR
	fmt.Println(convert("PAYPALISHIRING", 4)) // PINALSIGYAHRPI
	fmt.Println(convert("A", 1))              // A
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	row := 0
	step := 1

	for i := 0; i < len(s); i++ {
		rows[row] = append(rows[row], s[i])

		switch row {
		case 0:
			step = 1
		case numRows - 1:
			step = -1
		}
		row += step
	}

	res := make([]byte, len(s))
	for _, row := range rows {
		res = append(res, row...)
	}
	return string(res)
}
