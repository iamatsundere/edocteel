package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	numsx := make([]int, 0)
	l := len(nums)
	ret := make([][]int, 0)
	setNums := make(map[int]bool)
	for i := range nums {
		if _, ok := setNums[nums[i]]; !ok {
			numsx = append(numsx, nums[i])
			setNums[nums[i]] = true
		}
	}

	for len(numsx) < l {
		numsx = append(numsx, numsx[0])
	}

	setSlice := make(map[string]bool)
	if l >= 3 {
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				for k := j + 1; k < l; k++ {
					if numsx[i]+numsx[j]+numsx[k] == 0 {
						tmp := []int{numsx[i], numsx[j], numsx[k]}
						sort.Ints(tmp)
						key := fmt.Sprintf("%d_%d_%d", tmp[0], tmp[1], tmp[2])
						if _, ok := setSlice[key]; !ok {
							setSlice[key] = true
							ret = append(ret, tmp)
						}
					}
				}
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
