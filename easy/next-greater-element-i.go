// https: //leetcode.com/problems/next-greater-element-i/
package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	pairsVI := make(map[int]int)
	pairsIV := make(map[int]int)
	maxL := len(nums2)

	for i, v := range nums2 {
		pairsVI[v] = i
		pairsIV[i] = v
	}

	target := -1
	for _, v := range  nums1 {
		target = -1
		i := pairsVI[v]
		for i < maxL {
			i++
			if pairsIV[i] > v {
				target = pairsIV[i]
				break
			}
		}
		ret = append(ret, target)
	}

	return ret
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
