package main

func longestPalindromeSubseq(s string) int {
	src := []rune(s)
	dp := make([][]int, len(src))
	for i := range dp {
		dp[i] = make([]int, len(src))
	}

	for gap := 0; gap < len(dp); gap++ {
		for i := 0; i+gap < len(dp); i++ {
			if gap == 0 {
				dp[i][i+gap] = 1
			} else if gap == 1 {
				if src[i] == src[i+gap] {
					dp[i][i+gap] = 2
				} else {
					dp[i][i+gap] = 1
				}
			} else {
				if src[i] == src[i+gap] {
					dp[i][i+gap] = dp[i+1][i+gap-1] + 2
				} else {
					dp[i][i+gap] = max(dp[i+1][i+gap], dp[i][i+gap-1])
				}
			}

		}
	}

	return dp[0][len(dp)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
