// https: //leetcode.com/problems/split-array-into-fibonacci-sequence/
package main

import "fmt"

func fillStrings(a, b string) (string, string) {
	for len(a) < len(b) {
		a = "0" + a
	}
	for len(a) > len(b) {
		b = "0" + b
	}
	return a, b
}

func addNumberAsString(a, b string) string {
	ret := ""
	a, b = fillStrings(a, b)
	remember := 0
	for i := len(a) - 1; i >= 0; i-- {
		tmp := int(a[i]+b[i]-'0'-'0') + remember
		remember = tmp / 10
		tmp = tmp % 10
		ret = string(tmp+'0') + ret
	}

	if remember > 0 {
		ret = string(remember+'0') + ret
	}

	return ret
}

func splitIntoFibonacci(S string) []int {
	ret := make([]int, 0)
	return ret
}

func main() {
	fmt.Println(addNumberAsString("111", "2000"))
}
