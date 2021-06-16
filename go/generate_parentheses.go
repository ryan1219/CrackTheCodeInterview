package main

// question: https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	buildParenthesis("", n, n, &res)
	return res
}

func buildParenthesis(subString string, left int, right int, res *[]string) {
	if left > right {
		return
	}

	if left > 0 {
		buildParenthesis(subString+"(", left-1, right, res)
	}

	if right > 0 {
		buildParenthesis(subString+")", left, right-1, res)
	}

	if left == 0 && right == 0 {
		*res = append(*res, subString)
		return
	}
}
