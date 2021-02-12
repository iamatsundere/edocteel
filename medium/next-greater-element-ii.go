// https://leetcode.com/problems/next-greater-element-ii/
package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	ret := make([]int, 0)
	maxL := len(nums)

	target := -1
	j := 0
	for i, v := range nums {
		target = -1
		j = i + 1
		for true {
			// fmt.Println(i, j)
			if j == maxL {
				j = j % maxL
			}
			if nums[j] > v {
				target = nums[j]
				break
			}

			if i == j {
				break
			}

			j++
		}
		ret = append(ret, target)
	}

	return ret
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 5}))
	fmt.Println(nextGreaterElements([]int{3, 2, 1}))
}
