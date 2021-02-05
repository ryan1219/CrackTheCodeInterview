package main

import "sort"

// question: https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	recursiveSumII(candidates, make([]int, 0), 0, target, &result)
	return result
}

func recursiveSum(candidates []int, current []int, start, remain int, result *[][]int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		cpy := make([]int, len(current))
		copy(cpy, current)
		*result = append(*result, cpy)
	}

	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		recursiveSum(candidates, current, i, remain-candidates[i], result)
		current = current[:len(current)-1]
	}
}

// question: https://leetcode.com/problems/combination-sum-ii/
func recursiveSumII(candidates []int, current []int, start, remain int, result *[][]int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		cpy := make([]int, len(current))
		copy(cpy, current)
		*result = append(*result, cpy)
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] { // skip duplicate, require list to be sorted
			continue
		}
		current = append(current, candidates[i])
		recursiveSumII(candidates, current, i+1, remain-candidates[i], result) // Each number in candidates may only be used once in the combination.
		current = current[:len(current)-1]
	}
}
