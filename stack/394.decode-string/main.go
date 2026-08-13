package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(decodeString("3[a]2[bc]")) //aaabcbc
	fmt.Println(decodeString("3[a2[c]]")) //accaccacc
	// fmt.Println(decodeString("2[abc]3[cd]ef")) //abcabccdcdcdef
}

func decodeString(s string) string {
	// k > 0
	// 1 <= s.length <= 30
	// a-z0-9[]
	// integers in s in range [1,300]

	numStack := []int{}
	strStack := []string{}
	currStr := ""

	for i, n := 0, 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
			continue
		}

		if s[i] == '[' {
			numStack = append(numStack, n)
			strStack = append(strStack, currStr)
			n = 0
			currStr = ""
			continue
		}

		// 3[a]2[bc]
		// 3[a2[c]]
		if s[i] == ']' {
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currStr = prevStr + strings.Repeat(currStr, repeat)

			continue
		}

		currStr += string(s[i])
	}

	return currStr
}
