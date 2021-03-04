package main

import "math"

// question: https://leetcode.com/problems/squares-of-a-sorted-array/
func sortedSquares(A []int) []int {
	positive := 0
	res := make([]int, 0)

	for positive < len(A) {
		if A[positive] > 0 {
			break
		}
		positive++
	}

	negative := positive - 1

	for negative >= 0 && positive < len(A) {
		if math.Abs(float64(A[negative])) < math.Abs(float64(A[positive])) {
			res = append(res, A[negative]*A[negative])
			negative--
		} else {
			res = append(res, A[positive]*A[positive])
			positive++
		}
	}

	for negative >= 0 {
		res = append(res, A[negative]*A[negative])
		negative--
	}

	for positive < len(A) {
		res = append(res, A[positive]*A[positive])
		positive++
	}

	return res
}
