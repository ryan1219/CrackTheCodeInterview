package main

import "strconv"

// question: https://leetcode.com/problems/different-ways-to-add-parentheses/
func diffWaysToCompute(input string) []int {
	return diffWaysToComputeRecursive(input)
}

/*
recursive

break the equation base on operator, get left part and right part, each part is another subproblem.
after get values from each part, try different combination from each part
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

}
