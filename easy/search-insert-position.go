// https://leetcode.com/problems/search-insert-position/
package main

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	} else {
		for i, v := range nums {
			if v >= target {
				return i
			}
		}
	}
	return 0
}
