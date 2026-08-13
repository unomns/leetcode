package main

import (
	"fmt"
)

func main() {
	// fmt.Println(asteroidCollision([]int{5, 10, -5}))          // [5, 10]
	// fmt.Println(asteroidCollision([]int{8, -8}))              // []
	// fmt.Println(asteroidCollision([]int{10, 2, -5}))          // [10]
	// fmt.Println(asteroidCollision([]int{3, 5, -6, 2, -1, 4})) // [-6, 2, 4]
	// fmt.Println(asteroidCollision([]int{-10, 2, -5})) // [-10, -5]
	// fmt.Println(asteroidCollision([]int{10, 2, -5}))  // [10]
	// fmt.Println(asteroidCollision([]int{4, 2, -5}))   // [-5]
	// fmt.Println(asteroidCollision([]int{3, -3, 5}))   // [-5]

	fmt.Println(asteroidCollision2([]int{3, 5, -6, 2, -1, 4})) // [-6, 2, 4]
}

// len >= 2 && <= 10**4
// int >= -1000 && <= 1000
// int != 0
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, curr := range asteroids {
		if len(stack) == 0 || curr > 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, curr)
			continue
		}

		prev := stack[len(stack)-1]

		// go left
		for len(stack) > 1 && prev > 0 && prev < -curr {
			stack = stack[:len(stack)-1]
			prev = stack[len(stack)-1]
		}

		if prev < 0 {
			stack = append(stack, curr)
		} else if prev < -curr {
			stack[len(stack)-1] = curr
		} else if prev == -curr {
			stack = stack[:len(stack)-1]
		}
	}

	return stack
}

func asteroidCollision2(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, curr := range asteroids {
		destroyed := false

		// when do we go left
		for len(stack) > 0 && curr < 0 && stack[len(stack)-1] > 0 {
			prev := stack[len(stack)-1]

			if prev < -curr {
				stack = stack[:len(stack)-1]
				continue
			}

			if prev == -curr {
				stack = stack[:len(stack)-1]
			}

			destroyed = true
			break
		}

		if !destroyed {
			stack = append(stack, curr)
		}
	}

	return stack
}
