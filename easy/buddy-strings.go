// https://leetcode.com/problems/buddy-strings/
package main

import "fmt"

func buddyStrings(A string, B string) bool {
	ret := false
	diffs := make([]int, 0)
	chars := make(map[byte]int)

	if len(A) == len(B) {
		l := len(A)
		for i := 0; i < l; i++ {
			if A[i] != B[i] {
				diffs = append(diffs, i)
			}

			chars[A[i]]++
		}

		if len(diffs) == 3 {
			return false
		}

		if len(diffs) == 2 {
			if A[diffs[0]] == B[diffs[1]] && A[diffs[1]] == B[diffs[0]] {
				return true
			}
			return false
		}

		if len(diffs) == 0 {
			for _, v := range chars {
				if v > 1 {
					return true
				}
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(buddyStrings("abab", "abab"))
	fmt.Println(buddyStrings("ab", "ab"))
	fmt.Println(buddyStrings("aa", "aa"))
}
