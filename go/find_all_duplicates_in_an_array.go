package main

import "math"

// question: https://leetcode.com/problems/find-all-duplicates-in-an-array/
/*
because of condition 1 ≤ a[i] ≤ n, we can use array itself as a hash,
for each value a[i], set a[a[i] - 1] negative, to represent a[i] has been seen
*/
func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[int(math.Abs(float64(nums[i])))-1] < 0 {
			res = append(res, int(math.Abs(float64(nums[i]))))
		} else {
			nums[int(math.Abs(float64(nums[i])))-1] = -1 * nums[int(math.Abs(float64(nums[i])))-1]
		}
	}

	return res
}
