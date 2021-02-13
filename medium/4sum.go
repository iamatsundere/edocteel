// https://leetcode.com/problems/4sum/
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	size := len(nums)
	set := make(map[string]bool)
	if size > 3 {
		for i := 0; i < size-3; i++ {
			for j := i + 1; j < size-2; j++ {
				k := j + 1
				l := size - 1

				for k < l {
					tot := nums[i] + nums[j] + nums[k] + nums[l]
					if tot == target {
						key := fmt.Sprintf("%d_%d_%d_%d", nums[i], nums[j], nums[k], nums[l])
						if _, ok := set[key]; !ok {
							ret = append(ret, []int{nums[i], nums[j], nums[k], nums[l]})
							set[key] = true
						}
						k++
						l--
					} else if tot < target {
						k++
					} else {
						l--
					}
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
}
