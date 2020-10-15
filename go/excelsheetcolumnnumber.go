package main

import "math"

/*
question: https://leetcode.com/problems/excel-sheet-column-number/
*/
func titleToNumber(s string) int {
	src := []rune(s)
	if len(src) == 0 {
		return 0
	}
	value := int(src[len(src)-1]) - 64
	for i := len(src) - 2; i >= 0; i-- {
		value = value + (int(src[i])-64)*int(math.Pow(26, float64(len(src)-i-1)))
	}

	return value
}
