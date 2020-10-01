package main

/*
question: https://leetcode.com/problems/sliding-window-maximum/
*/

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	queue := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for len(queue) != 0 && queue[0] < i-k+1 {
			queue = remove(queue, 0)
		}

		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = remove(queue, len(queue)-1)
		}

		queue = append(queue, i)

		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}
	}
	return result
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
