// https://leetcode.com/problems/longest-substring-without-repeating-characters/solution/
package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	l := len(s)
	x := make([]int, l)
	has := make(map[byte]bool)

	for i := range s {
		x[i] = 1
		has[s[i]] = true
		for j := i + 1; j < l; j++ {
			if _, ok := has[s[j]]; !ok {
				x[i]++
				has[s[j]] = true
			} else {
				break
			}
		}
		has = make(map[byte]bool)
	}

	// fmt.Println(x)
	max := x[0]
	for i := range x {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("aa"))
	fmt.Println(lengthOfLongestSubstring("vbxpvwkkteaiob"))
}
