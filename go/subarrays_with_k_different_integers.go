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

func subarraysWithKDistinct(A []int, K int) int {
	window1 := make(map[int]int)
	window2 := make(map[int]int)
	w1, w2, count := 0, 0, 0

	for i := 0; i < len(A); i++ {
		if val, contained := window1[A[i]]; contained {
			window1[A[i]] = val + 1
		} else {
			window1[A[i]] = 1
		}

		if val, contained := window2[A[i]]; contained {
			window2[A[i]] = val + 1
		} else {
			window2[A[i]] = 1
		}

		for len(window1) > K {
			if window1[A[w1]] == 1 {
				delete(window1, A[w1])
			} else {
				window1[A[w1]]--
			}
			w1++
		}

		for len(window2) >= K {
			if window2[A[w2]] == 1 {
				delete(window2, A[w2])
			} else {
				window2[A[w2]]--
			}
			w2++
		}

		if len(window1) == K {
			count += w2 - w1
		}
	}

	return count
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
}
