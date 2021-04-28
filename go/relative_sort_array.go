package main

import "sort"

// question: https://leetcode.com/problems/relative-sort-array/
func relativeSortArray(arr1 []int, arr2 []int) []int {
	count := make(map[int]int)
	arr2Map := make(map[int]int)
	extra := make([]int, 0)
	result := make([]int, 0)

	for _, num := range arr2 {
		arr2Map[num] = 1
	}

	for _, num := range arr1 {
		if _, contained := arr2Map[num]; !contained {
			extra = append(extra, num)
			continue
		}

		if value, contained := count[num]; contained {
			count[num] = value + 1
		} else {
			count[num] = 1
		}
	}

	for _, num := range arr2 {
		times := count[num]

		for times > 0 {
			result = append(result, num)
			times--
		}
	}

	sort.Ints(extra)
	result = append(result, extra...)

	return result
}
