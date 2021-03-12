package main

import "math"

// question: https://leetcode.com/problems/integer-replacement/
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n == math.MaxInt32 {
		return 32
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	} else {
		return int(math.Min(float64(1+integerReplacement(n+1)), float64(1+integerReplacement(n-1))))
	}
}

// a clean dp solution but run out of memory
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 0
	dp[2] = 1

	for i := 3; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2] + 1
		} else {
			dp[i] = int(math.Min(float64(dp[i-1]+1), float64(dp[(i+1)/2]+2)))
		}
	}

	return dp[n]
}
