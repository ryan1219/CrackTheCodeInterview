package main

// question: https://leetcode.com/problems/regular-expression-matching/
func isMatch(s string, p string) bool {
	// dp[i][j] denotes if s.substring(0,i) is valid for pattern p.substring(0,j)
	if p == "" {
		return s == ""
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true

	for j := 2; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for j := 1; j <= len(p); j++ {
		for i := 1; i <= len(s); i++ {
			if p[j-1] == '*' {
				// * is not used
				dp[i][j] = dp[i][j-2]

				// * is used
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else {
				if p[j-1] == s[i-1] || p[j-1] == '.' {
					dp[i][j] = dp[i-1][j-1]
				}
			}

		}
	}

	return dp[len(s)][len(p)]
}

// i is the index for string s, j is the index for string j
func isMatchRecursive(s string, i int, p string, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	// if there is still another char after index j, check if it is *
	if j+1 < len(p) && p[j+1] == '*' {
		// * is not used
		if isMatchRecursive(s, i, p, j+2) {
			return true
		}
		// * is used
		if i < len(s) && (p[j] == '.' || s[i] == p[j]) {
			return isMatchRecursive(s, i+1, p, j)
		}
		// second char is not *
	} else {
		if i < len(s) && (p[j] == '.' || s[i] == p[j]) {
			return isMatchRecursive(s, i+1, p, j+1)
		}
	}
	return false
}
