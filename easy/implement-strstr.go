// https://leetcode.com/problems/implement-strstr/
package main

import "fmt"

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		// fmt.Println(haystack[i : i+len(needle)])
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", ""))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}
