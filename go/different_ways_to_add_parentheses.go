package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// question: https://leetcode.com/problems/different-ways-to-add-parentheses/
func diffWaysToCompute(input string) []int {
	return diffWaysToComputeRecursive(input)
}

/*
recursive

break the equation base on operator '-' '+' '*', get left part and right part, each part is another subproblem.
after return values from each part, try different combination from each part.
base case: only one number is on subproblem
*/
func diffWaysToComputeRecursive(input string) []int {
	res := make([]int, 0)

	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToComputeRecursive(input[:i])
			right := diffWaysToComputeRecursive(input[i+1:])

			for _, l := range left {
				for _, r := range right {
					if input[i] == '+' {
						res = append(res, l+r)
					} else if input[i] == '-' {
						res = append(res, l-r)
					} else if input[i] == '*' {
						res = append(res, l*r)
					}
				}
			}
		}
	}

	// base case, only one number passed in
	if len(res) == 0 {
		i, _ := strconv.Atoi(input)
		res = append(res, i)
	}

	return res
}

func diffWaysToComputeDP(input string) []int {
	// dp[i][j] stores all possible results from i-th index to the j-th(inclusive) index
	dp := make([][][]int, len(input))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, len(input))
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, 0)
		}
	}

	// if from i to j there are all digits, then index i to j is a number
	for i := 0; i < len(input); i++ {
		if input[i] != '+' && input[i] != '-' && input[i] != '*' {
			j := i + 1
			for j < len(input) && unicode.IsDigit(int32(input[j])) {
				j++
			}
			value, _ := strconv.Atoi(input[i:j])
			dp[i][j-1] = append(dp[i][j-1], value)
		}
	}

	// length is the length of substring. need to calculate how many results can get between substring [i, i+length]
	for length := 1; length < len(input); length++ {
		for i := 0; i < len(input)-length; i++ {
			// when calculate the substring [i, i+length], j is the index of the operator.
			for j := i; j <= i+length; j++ {
				if input[j] == '+' || input[j] == '-' || input[j] == '*' {
					left := dp[i][j-1]
					right := dp[j+1][i+length]

					for _, leftV := range left {
						for _, rightV := range right {
							if input[j] == '+' {
								dp[i][i+length] = append(dp[i][i+length], leftV+rightV)
							} else if input[j] == '-' {
								dp[i][i+length] = append(dp[i][i+length], leftV-rightV)
							} else if input[j] == '*' {
								dp[i][i+length] = append(dp[i][i+length], leftV*rightV)
							}
						}
					}
				}
			}
		}
	}

	return dp[0][len(input)-1]
}

func main() {
	fmt.Println(diffWaysToComputeDP("2-1-1"))
}
