package main

// question: https://leetcode.com/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	space := make([]int, len(nums))
	for i := range space {
		space[i] = 1
	}

	for i := 0; i < len(space); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && space[j]+1 > space[i] {
				space[i] = space[j] + 1
			}
		}
	}

	var m int

	for _, e := range space {
		if e > m {
			m = e
		}
	}

	return m
}
