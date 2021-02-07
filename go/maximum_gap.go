package main

// question: https://leetcode.com/problems/maximum-gap/
/*
general approcah is first get list of numbers sorted, then loop to find maximum gap
sorting algorithm affects the runtime. A comparison sort cannot perform better than O(n log n).
non-comparison sorts are not limited to Ω(n log n)
*/
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	radixSort(nums)
	maxGap := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > maxGap {
			maxGap = nums[i+1] - nums[i]
		}
	}

	return maxGap
}
