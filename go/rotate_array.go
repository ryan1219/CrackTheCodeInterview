package main

// question: https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}

		nums[0] = tmp
	}
}
