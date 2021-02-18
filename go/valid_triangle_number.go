package main

import "sort"

// question: https://leetcode.com/problems/valid-triangle-number/
/*
Assume a is the longest edge, b and c are shorter ones, to form a triangle, they need to satisfy len(b) + len(c) > len(a).
sort the list, let each number be the longest edge
*/
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0

	for i := 2; i < len(nums); i++ {
		left := 0
		right := i - 1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}
	return count
}
