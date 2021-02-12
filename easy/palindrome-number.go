// https://leetcode.com/problems/palindrome-number/
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	s := []rune(strconv.Itoa(x))
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(-120))
	fmt.Println(isPalindrome(120))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(0))
}
