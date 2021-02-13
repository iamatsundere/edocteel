package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	l := len(nums)
	sort.Ints(nums)
	ret := make([][]int, 0)
	set := make(map[string]bool)

	if l >= 3 {
		for i := range nums {
			j := i + 1
			k := len(nums) - 1

			for j < k {
				tot := nums[i] + nums[j] + nums[k]
				if tot == 0 {
					key := fmt.Sprintf("%d_%d_%d", nums[i], nums[j], nums[k])
					if _, ok := set[key]; !ok {
						ret = append(ret, []int{nums[i], nums[j], nums[k]})
						set[key] = true
					}
					j++
					k--
				} else if tot < 0 {
					j++
				} else {
					k--
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
