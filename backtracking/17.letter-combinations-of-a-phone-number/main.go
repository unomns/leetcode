package main

import (
	"fmt"
)

/**
Given a string containing digits from 2-9 inclusive,
return all possible letter combinations that the number could represent.

Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below.
Note that 1 does not map to any letters.


Example 1:
	Input: digits = "23"
	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
	Input: digits = "2"
	Output: ["a","b","c"]


Constraints:
	1 <= digits.length <= 4
	digits[i] is a digit in the range ['2', '9'].
*/

func main() {
	fmt.Println(letterCombinations("23")) // ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations("2"))  // ["a","b","c"]
}

func letterCombinations(digits string) []string {
	set := [10]string{
		0: "",
		1: "",
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	res := []string{}
	var backtrack func(idx int, curr string)

	backtrack = func(idx int, curr string) {
		if idx == len(digits) {
			res = append(res, curr)
			return
		}

		letters := set[digits[idx]-'0']
		for i := 0; i < len(letters); i++ {
			backtrack(idx+1, curr+string(letters[i]))
		}
	}

	backtrack(0, "")

	return res
}
