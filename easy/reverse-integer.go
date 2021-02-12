// https://leetcode.com/problems/reverse-integer/
package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	MIN := -2147483648
	MAX := 2147483647

	ret := 0
	s := []rune(strconv.Itoa(x))
	sx := ""

	if s[0] == '-' {
		sx += "-"
		s = s[1:]
	}

	if s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}

	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = tmp
	}

	sx += string(s)
	ret, _ = strconv.Atoi(sx)

	if ret > MAX || ret < MIN {
		return 0
	}

	return ret
}

func main() {
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(1230))
	fmt.Println(reverse(-1234506))
	fmt.Println(reverse(120))
}
