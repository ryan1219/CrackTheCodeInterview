package main

import (
	"fmt"
	"strings"
)

// question: https://leetcode.com/problems/prison-cells-after-n-days/
func prisonAfterNDays(cells []int, N int) []int {
	// use map to track the path, find the cycle
	seen := make(map[string]int)
	totalDays := N

	for N > 0 {
		tmp := make([]int, len(cells))

		for i := 1; i < len(cells)-1; i++ {
			if cells[i-1] == cells[i+1] {
				tmp[i] = 1
			} else {
				tmp[i] = 0
			}
		}
		key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tmp)), ""), "[]")
		if _, contained := seen[key]; contained {
			break
		}
		seen[key] = 1
		cells = tmp

		N--
	}

	// hit cycle
	if N > 0 {
		totalDays = totalDays % len(seen)
		for totalDays > 0 {
			tmp := make([]int, len(cells))
			for i := 1; i < len(cells)-1; i++ {
				if cells[i-1] == cells[i+1] {
					tmp[i] = 1
				} else {
					tmp[i] = 0
				}
			}
			cells = tmp
			totalDays--
		}
	}

	return cells
}

func main() {
	test := []int{0, 1, 0, 1, 1, 0, 0, 1}
	fmt.Println(prisonAfterNDays(test, 15))
}
