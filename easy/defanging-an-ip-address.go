// https://leetcode.com/problems/defanging-an-ip-address/
package main

import "fmt"

func defangIPaddr(address string) string {
	retx := make([]rune, 0)
	for _, c := range address {
		if c != '.' {
			retx = append(retx, c)
		} else {
			retx = append(retx, '[', '.', ']')
		}
	}
	return string(retx)
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
