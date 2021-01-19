package main

// question: https://leetcode.com/problems/diagonal-traverse/
func findDiagonalOrder(matrix [][]int) []int {
	result := make([]int, 0)
	if len(matrix) == 0 {
		return result
	}
	m := len(matrix)
	n := len(matrix[0])

	row := 0
	col := 0
	dir := 1

	for i := 0; i < m*n; i++ {
		result = append(result, matrix[row][col])
		row -= dir
		col += dir

		if row >= m {
			row = m - 1
			col = col + 2
			dir = -1 * dir
		}
		if col >= n {
			col = n - 1
			row = row + 2
			dir = -1 * dir
		}
		if row < 0 {
			row = 0
			dir = -1 * dir
		}
		if col < 0 {
			col = 0
			dir = -1 * dir
		}
	}

	return result
}
