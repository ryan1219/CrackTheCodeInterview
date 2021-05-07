package main

// question: https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/
func numRollsToTarget(d int, f int, target int) int {
	mod := 1000000000 + 7
	dp := make([][]int, d+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1)
	}

	dp[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := 0; j <= target; j++ {
			for k := 1; k <= j && k <= f; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][j-k]) % mod
			}
		}
	}

	return dp[d][target]
}
