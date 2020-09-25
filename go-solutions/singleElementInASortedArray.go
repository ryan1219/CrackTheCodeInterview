package main

/*
question:https://leetcode.com/problems/single-element-in-a-sorted-array/submissions/
key note:
- binary search
- if all numbers in pair, 0, 0, 1, 1, 3, 3, value at even index == value at odd index
- one single value introduced, 0, 0, 1, 2, 2, 4, 4, if value at even index != value at odd index,
must be single value on the left so all value shift one index up
*/
func singleNonDuplicate(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) / 2
		if mid%2 != 0 {
			mid -= 1
		}
		if nums[mid] == nums[mid+1] {
			start = mid + 2
		} else {
			end = mid
		}
	}
	return nums[start]
}
