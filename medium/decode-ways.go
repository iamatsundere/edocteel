// https://leetcode.com/problems/decode-ways/

package main

import (
	"fmt"
	"strconv"
)

var ret int

func solve(n string) {
	ranges := []int{1, 2}
	// fmt.Println(n, *ret)

	if n == "" {
		ret++
	} else {
		for _, div := range ranges {
			if len(n) < div {
				continue
			}

			mod, _ := strconv.Atoi(n[len(n)-div : len(n)])
			tmp := n[0 : len(n)-div]
			// fmt.Println(n, mod, tmp)

			switch div {
			case 1:
				if mod >= 1 && mod <= 9 {
					solve(tmp)
				}
				break
			case 2:
				if mod >= 10 && mod <= 26 {
					solve(tmp)
				}
				break
			}
		}
	}
}

func numDecodings(s string) int {
	ret = 0
	if s != "0" {
		solve(s)
	}
	return ret
}

func main() {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("10"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("01"))
	fmt.Println(numDecodings("4757562545844617494555774581341211511296816786586787755257741178599337186486723247528324612117156948"))
}
