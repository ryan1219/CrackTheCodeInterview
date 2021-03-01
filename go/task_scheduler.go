package main

// question: https://leetcode.com/problems/task-scheduler/
func leastInterval(tasks []byte, n int) int {
	counter := make([]int, 26)
	max := 0
	maxCount := 0

	for _, v := range tasks {
		counter[v-'A']++
		if counter[v-'A'] == max {
			maxCount++
		} else if counter[v-'A'] > max {
			max = counter[v-'A']
			maxCount = 1
		}
	}

	gapNumber := max - 1
	gapLength := n - (maxCount - 1)
	gapSlots := gapNumber * gapLength
	idle := 0
	if len(tasks)-max*maxCount < gapSlots {
		idle = gapSlots - (len(tasks) - max*maxCount)
	}

	return len(tasks) + idle
}
