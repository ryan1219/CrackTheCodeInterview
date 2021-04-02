package main

import "sort"

// question: https://leetcode.com/problems/two-city-scheduling/
func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	var total int
	for i := 0; i < len(costs)/2; i++ {
		// Every cost up to n is city A, everyone else is city B
		total += costs[i][0] + costs[i+len(costs)/2][1]
	}
	return total
}
