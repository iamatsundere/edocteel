//https://leetcode.com/problems/daily-temperatures/
package main

import "fmt"

func dailyTemperatures(T []int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(T); i++ {
		lenx := 0
		t := T[i]
		for j := i + 1; j < len(T); j++ {
			tx := T[j]
			if tx > t {
				lenx = j - i
				break
			}
		}
		ret = append(ret, lenx)
	}
	return ret
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
