package main

// quetsion: https://leetcode.com/problems/contiguous-array/
func findMaxLength(nums []int) int {
	count := 0
	maxLength := 0

	// key:value -> after taking the value of the number at index i into account, the count is the key, value is the index i
	table := make(map[int]int)
	table[0] = -1

	for i, v := range nums {
		if v == 0 {
			count--
		} else {
			count++
		}

		if value, contained := table[count]; contained {
			maxLength = max(maxLength, i-value)
		} else {
			table[count] = i
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
