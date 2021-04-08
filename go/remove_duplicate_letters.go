package main

import (
	"fmt"
	"math"
)

// question: https://leetcode.com/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
	array := []rune(s)
	m := make(map[rune]int)

	for i, c := range array {
		m[c] = i
	}

	result := make([]rune, len(m))
	start := 0
	for i := 0; i < len(result); i++ {
		end := findMinPos(m)
		min := 'z' + 1

		for j := start; j <= end; j++ {
			if _, contained := m[array[j]]; contained && array[j] < min {
				min = array[j]
				start = j + 1
			}
		}

		result[i] = min
		delete(m, min)
	}

	return string(result)
}

func findMinPos(m map[rune]int) int {
	pos := math.MaxInt64
	for _, v := range m {
		if v < pos {
			pos = v
		}
	}

	return pos
}

func main() {
	fmt.Println(removeDuplicateLetters("abcd"))
}
