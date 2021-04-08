package main

// question: https://leetcode.com/problems/height-checker/
func heightChecker(heights []int) int {
	count := make([]int, 101)

	for _, height := range heights {
		count[height]++
	}

	unordered := 0
	head := 0

	for _, height := range heights {
		for count[head] == 0 {
			head++
		}
		if head != height {
			unordered++
		}

		count[head]--
	}

	return unordered
}
