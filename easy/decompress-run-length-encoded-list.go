// https://leetcode.com/problems/decompress-run-length-encoded-list/
package main

import "fmt"

func decompressRLElist(nums []int) []int {
	ret := make([]int, 0)
	for i := range nums {
		if i%2 == 0 {
			// fmt.Println(i, nums[i], nums[i+1])
			for nums[i] > 0 {
				ret = append(ret, nums[i+1])
				nums[i]--
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}
