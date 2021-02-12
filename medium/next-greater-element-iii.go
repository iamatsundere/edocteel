// https://leetcode.com/problems/next-greater-element-iii/
package main

import "fmt"

func nextGreaterElement(n int) int {
	ret := 0
	nx := make([]int, 0)
	tmp := n

	for tmp != 0 {
		nx = append(nx, tmp%10)
		tmp = tmp / 10
	}

	for i := 0; i < len(nx)/2; i++ {
		tmp = nx[i]
		nx[i] = nx[len(nx)-1-i]
		nx[len(nx)-1-i] = tmp
	}

	pos := 0
	if len(nx) >= 3 {
		for i := 1; i < len(nx)-1; i++ {
			if (nx[i]*nx[i-1])*(nx[i+1]-nx[i]) > 0 {
				continue
			} else {
				pos = i
				break
			}
		}
	}

	fmt.Println(pos)

	return ret
}

func main() {
	fmt.Println("ret", nextGreaterElement(15234))
}
