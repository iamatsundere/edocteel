// https://leetcode.com/problems/count-and-say/
package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	s := ""
	ns := "1"
	for n-1 > 0 {
		s = ""
		flag := ns[0]
		count := 0
		for i := range ns {
			if flag == ns[i] {
				count++
			} else {
				s += fmt.Sprintf("%d%s", count, string(flag))
				count = 1
				flag = ns[i]
			}
		}

		s += fmt.Sprintf("%d%s", count, string(flag))
		ns = s
		// fmt.Println(ns)
		n--
	}
	return s
}

func main() {
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(30))
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
}
