package main

// question: https://leetcode.com/problems/interleaving-string/
/*
recursion

i := index of s1, [0, len(s1))
j := index of s2, [0, len(s2))

func isInterleave(s1 string, s2 string, i, j int, s3 string)

base case: i + j == len(s3) return true
recursive case:
1. when next char can either from s1 or s2
s3[i+j] == s1[i] && s3[i+j] == s2[j]
return isInterleave(s1, s2, i+1, j, s3) || isInterleave(s1, s2, i, j+1, s3)
2. when next char can only from s1
s3[i+j] == s1[i]
return isInterleave(s1, s2, i+1, j, s3)
3. when next char can only from s2
s3[i+j] == s2[j]
return isInterleave(s1, s2, i, j+1, s3)
*/

/*
DP
dp[i][j]

i:= chars in s1 until index i-1(included)
j:= chars in s2 until index j-1(included)

dp[i][j] == true, means first i chars in s1 and first j chars in s2 are interleave for first i+j chars in s3
goal: calcualte dp[i][j] where i = len(s1), j = len(s2)

*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) < len(s3) {
		return false
	}
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true
	for i := 1; i <= len(s1); i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j <= len(s2); j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i-1+j] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}

			if s2[j-1] == s3[j-1+i] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
