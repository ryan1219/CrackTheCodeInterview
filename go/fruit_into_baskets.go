package main

/*
question: https://leetcode.com/problems/fruit-into-baskets/

find the longest subarray with at most 2 different numbers
*/
func totalFruit(tree []int) int {
	seen := make(map[int]int)
	max := 0

	i, j := 0, 0
	for ; j < len(tree); j++ {
		if val, contained := seen[tree[j]]; contained {
			seen[tree[j]] = val + 1
		} else {
			seen[tree[j]] = 1
		}
		for len(seen) > 2 {
			if seen[tree[i]] > 1 {
				seen[tree[i]]--
			} else {
				delete(seen, tree[i])
			}
			i++
		}
		if j-i+1 > max {
			max = j - i + 1
		}
	}

	return max
}
