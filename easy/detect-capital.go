// https://leetcode.com/problems/detect-capital/
package main

import "fmt"

func detectCapitalUse(word string) bool {
	l := len(word)
	count := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			count++
		}
	}

	fmt.Println(count)

	if count == l || count == 0 {
		return true
	} else {
		if count == 1 {
			if word[0] >= 'A' && word[0] <= 'Z' {
				return true
			}
		} else {
			return false
		}
		return false
	}
}

func main() {
	fmt.Println(detectCapitalUse("Leetcode"))
}
