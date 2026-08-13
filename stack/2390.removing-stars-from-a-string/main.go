package main

import "fmt"

func main() {
	fmt.Println(removeStars("**eet**cod*e")) // leecoe
}

func removeStars(s string) string {
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
