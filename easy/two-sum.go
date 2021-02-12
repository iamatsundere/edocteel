// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	pairsV2Is := make(map[int][]int)
	for i, v := range nums {
		pairsV2Is[v] = append(pairsV2Is[v], i)
	}

	for v, is := range pairsV2Is {
		if len(is) >= 2 && v*2 == target {
			return []int{is[0], is[1]}
		}

		if len(pairsV2Is[target-v]) > 0 && target-v != v {
			return []int{is[0], pairsV2Is[target-v][0]}
		}
	}

	// i := 0
	// j := 0
	// for ; i < len(nums)-1; i++ {
	// 	for j = i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == target {
	// 			ret[0] = i
	// 			ret[1] = j
	// 			return ret
	// 		}
	// 	}
	// }
	return ret
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{1, 0, 0, 0, 0, 0, 3, 4}, 5))
}
