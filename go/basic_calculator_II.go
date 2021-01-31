package main

import (
	"strconv"
	"strings"
	"unicode"
)

// question: https://leetcode.com/problems/basic-calculator-ii/
func calculate(s string) int {
	s = strings.Join(strings.Fields(s), "")
	stack := make([]string, 0)
	result := 0

	i := 0
	for i < len(s) {
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			stack = append(stack, string(s[i]))
			i++
		} else {
			j := i + 1
			for j < len(s) && unicode.IsDigit(int32(s[j])) {
				j++
			}
			if len(stack) == 0 {
				stack = append(stack, s[i:j])
			} else if stack[len(stack)-1] == "*" || stack[len(stack)-1] == "/" {
				v1, _ := strconv.Atoi(stack[len(stack)-2])
				v2, _ := strconv.Atoi(s[i:j])
				res := 0
				if stack[len(stack)-1] == "*" {
					res = v1 * v2
				} else {
					res = v1 / v2
				}
				// pop last two values and push res
				stack = stack[:len(stack)-2]
				stack = append(stack, strconv.Itoa(res))
			} else {
				stack = append(stack, s[i:j])
			}
			i = j
		}
	}

	operator := ""
	for _, v := range stack {
		if v == "+" || v == "-" {
			operator = v
			continue
		}
		num, _ := strconv.Atoi(v)
		if operator == "" || operator == "+" {
			result += num
		} else {
			result -= num
		}
	}

	return result
}
