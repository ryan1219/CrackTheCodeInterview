package main

// question: https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	recursiveSum(candidates, make([]int, 0), 0, target, &result)
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
