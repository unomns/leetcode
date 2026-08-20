package main

import "fmt"

/**
There are n cities. Some of them are connected, while some are not.
If city a is connected directly with city b, and city b is connected directly with city c,
then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities
and no other cities outside of the group.

You are given an n x n matrix isConnected
where isConnected[i][j] = 1 if the ith city and the jth city are directly connected,
and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.

Example 1:
	Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
	Output: 2

Example 2:
	Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
	Output: 3


Constraints:
	1 <= n <= 200
	n == isConnected.length
	n == isConnected[i].length
	isConnected[i][j] is 1 or 0.
	isConnected[i][i] == 1
	isConnected[i][j] == isConnected[j][i]
*/

func main() {
	cities := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}

	fmt.Println(findCircleNum(cities))
}

func findCircleNum(isConnected [][]int) int {
	provincies := 0
	visited := make([]bool, len(isConnected))

	var dfs func(city int)

	dfs = func(city int) {
		visited[city] = true

		for i, k := range isConnected[city] {
			if k == 0 {
				continue
			}
			if !visited[i] {
				dfs(i)
			}
		}
	}

	for i := range isConnected {
		if !visited[i] {
			dfs(i)
			provincies++
		}
	}

	return provincies
}
