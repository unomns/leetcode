package main

import "fmt"

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
	Input: strs = ["flower","flow","flight"]
	Output: "fl"

Example 2:
	Input: strs = ["dog","racecar","car"]
	Output: ""
	Explanation: There is no common prefix among the input strings.


Constraints:
	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lowercase English letters if it is non-empty.
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	for i := 0; i < len(strs[0]); i++ {
		cmp := strs[0][i]
		for j := 1; j < n; j++ {
			if i >= len(strs[j]) || cmp != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
