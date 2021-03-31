package main

// question: https://leetcode.com/problems/camelcase-matching/
func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, 0)

	for _, s := range queries {
		res = append(res, isMatch(s, pattern))
	}

	return res
}

func isMatch(src, pattern string) bool {
	chars := []rune(pattern)
	i := 0

	for _, c := range src {
		if i < len(chars) && c == chars[i] {
			i++
		} else if c < 'a' {
			return false
		}
	}

	return i == len(chars)
}
