package main

// question: https://leetcode.com/problems/longest-valid-parentheses/
// solution1: dp
// dp array where iith element of dp represents the length of the longest valid substring ending at ith index.
func longestValidParentheses_dp(s string) int {
	max := 0
	dp := make([]int, len(s))
	sBytes := []byte(s)

	for i := 1; i < len(sBytes); i++ {
		if sBytes[i] == ')' {
			if sBytes[i-1] == '(' {
				if i-2 > 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && sBytes[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 > 0 {
					dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}

	return max
}

// solution2: stack
func longestValidParentheses_stack(s string) int {
	max := 0
	sBytes := []byte(s)
	stack := make([]int, 0)
	stack = append(stack, -1)

	for i := 0; i < len(sBytes); i++ {
		if sBytes[i] == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				tmp := i - stack[len(stack)-1]
				if tmp > max {
					max = tmp
				}
			}
		} else {
			stack = append(stack, i)
		}
	}

	return max
}
