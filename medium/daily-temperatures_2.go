//https://leetcode.com/problems/daily-temperatures/
package main

import "fmt"

func dailyTemperatures(T []int) []int {
	ret := make([]int, 0)
	stack := make([]int, 0)

	for i := len(T) - 1; i >= 0; i-- {
		for len(stack) > 0 && T[stack[len(stack)-1]] <= T[i] {
			stack = stack[0 : len(stack)-1]
		}

		if len(stack) > 0 {
			ret = append(ret, stack[len(stack)-1]-i)
		} else {
			ret = append(ret, 0)
		}

		stack = append(stack, i)
	}

	for i := 0; i < len(ret)/2; i++ {
		tmp := ret[i]
		ret[i] = ret[len(ret)-1-i]
		ret[len(ret)-1-i] = tmp
	}

	return ret
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
