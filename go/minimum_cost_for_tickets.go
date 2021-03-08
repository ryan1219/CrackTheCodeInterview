package main

import "fmt"

// question: https://leetcode.com/problems/minimum-cost-for-tickets/
/*
We track the minimum cost for all calendar days in dp.
For non-travel days, the cost stays the same as for the previous day.
For travel days, it's a minimum of yesterday's cost plus single-day ticket, or cost for 8 days ago plus 7-day pass, or cost 31 days ago plus 30-day pass.
*/
func mincostTickets(days []int, costs []int) int {
	dp := make([]int, days[len(days)-1]+1)
	dayIndex := 0
	for i := 1; i < len(dp); i++ {
		if i != days[dayIndex] {
			dp[i] = dp[i-1]
		} else {
			dayIndex++
			dp[i] = min(dp[i-1]+costs[0], dp[max(0, i-7)]+costs[1], dp[max(0, i-30)]+costs[2])
		}
	}
	return dp[len(dp)-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= c && b <= a {
		return b
	}

	return c
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}
