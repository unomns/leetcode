package main

import (
	"fmt"
)

/**
You are given an array of variable pairs equations and an array of real numbers values,
where equations[i] = [Ai, Bi] and values[i] represent the equation Ai / Bi = values[i].
Each Ai or Bi is a string that represents a single variable.

You are also given some queries,
where queries[j] = [Cj, Dj] represents the jth query
where you must find the answer for Cj / Dj = ?.

Return the answers to all queries.
If a single answer cannot be determined, return -1.0.

Note:
	The input is always valid.
	You may assume that evaluating the queries will not result in division by zero
	and that there is no contradiction.

Note:
	The variables that do not occur in the list of equations are undefined,
	so the answer cannot be determined for them.


Example 1:
	Input: equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
	Output: [6.00000,0.50000,-1.00000,1.00000,-1.00000]
	Explanation:
	Given: a / b = 2.0, b / c = 3.0
	queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
	return: [6.0, 0.5, -1.0, 1.0, -1.0 ]
	note: x is undefined => -1.0

Example 2:
	Input: equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0], queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
	Output: [3.75000,0.40000,5.00000,0.20000]

Example 3:
	Input: equations = [["a","b"]], values = [0.5], queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
	Output: [0.50000,2.00000,-1.00000,-1.00000]


Constraints:
	1 <= equations.length <= 20
	equations[i].length == 2
	1 <= Ai.length, Bi.length <= 5
	values.length == equations.length
	0.0 < values[i] <= 20.0
	1 <= queries.length <= 20
	queries[i].length == 2
	1 <= Cj.length, Dj.length <= 5
	Ai, Bi, Cj, Dj consist of lower case English letters and digits.
*/

func main() {
	equations := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2.0, 3.0}
	queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	fmt.Println(calcEquation(equations, values, queries)) // [6.00000,0.50000,-1.00000,1.00000,-1.00000]
}

func calcEquation(
	equations [][]string,
	values []float64,
	queries [][]string,
) []float64 {
	type node struct {
		v string
		w float64
	}

	adj := make(map[string][]node)

	for i, v := range equations {
		a, b := v[0], v[1]
		weight := values[i]

		adj[a] = append(adj[a], node{v: b, w: weight})
		adj[b] = append(adj[b], node{v: a, w: 1 / weight})
	}

	var dfs func(c, j string, res float64, visited map[string]bool) float64

	dfs = func(c, d string, res float64, visited map[string]bool) float64 {
		visited[c] = true
		if c == d {
			return res
		}

		for _, v := range adj[c] {
			if !visited[v.v] {
				if res := dfs(v.v, d, res*v.w, visited); res != -1 {
					return res
				}
			}
		}

		return -1
	}

	res := make([]float64, len(queries))

	for i, v := range queries {
		c, d := v[0], v[1]

		if _, ok := adj[c]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := adj[d]; !ok {
			res[i] = -1
			continue
		}
		res[i] = dfs(c, d, 1, make(map[string]bool))
	}

	return res
}
