package main

import "fmt"

func main() {
	// fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
	//0,-4, -7, -9, -10,-6,-3,-1

	fmt.Println(largestAltitude([]int{52, -91, 72}))
	//0,52,-39,33
	fmt.Println(largestAltitudeWithoutAltdsInMemory([]int{52, -91, 72}))
	//0,52,-39,33
}

func largestAltitude(gain []int) int {
	peak := 0

	altds := []int{0}
	for i := range gain {
		val := altds[i] + gain[i]
		altds = append(altds, val)

		if val > peak {
			peak = val
		}
	}

	fmt.Println(altds)

	return peak
}

func largestAltitudeWithoutAltdsInMemory(gain []int) int {
	peak := 0
	h := peak
	for _, net := range gain {
		h += net

		if h > peak {
			peak = h
		}
	}

	return peak
}
