package main

// question: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
func twoSum(numbers []int, target int) []int {
	result := make([]int, 0)
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] < target {
			i++
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			break
		}
	}
	result = append(result, i+1, j+1)
	return result
}
