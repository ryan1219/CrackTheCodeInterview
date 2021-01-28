package main

import "sort"

// question: https://leetcode.com/problems/subsets-ii/
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, make([]int, 0))
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	start := 0
	stop := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			start = stop
		} else {
			start = 0
		}
		stop = len(res)
		for j := start; j < stop; j++ {
			cpy := make([]int, len(res[j]))
			copy(cpy, res[j])
			res = append(res, append(cpy, nums[i]))
		}
	}
	return res
}
