package main

import "fmt"

// https://leetcode.com/problems/subarrays-with-k-different-integers/
func subarraysWithKDistinct(A []int, K int) int {
	seen := make(map[int]int)
	j, count, sameGroup := 0, 0, 0

	for i := 0; i < len(A); i++ {
		if val, contained := seen[A[i]]; contained {
			seen[A[i]] = val + 1
		} else {
			seen[A[i]] = 1
		}

		if len(seen) > K {
			delete(seen, A[j])
			j++
			sameGroup = 0
		}
		// count number of subarrays with the same number of distinct numbers
		for seen[A[j]] > 1 {
			seen[A[j]]--
			j++
			sameGroup++
		}

		if len(seen) == K {
			count += sameGroup + 1
		}
	}

	return count
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}
