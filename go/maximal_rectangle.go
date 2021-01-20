package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	var max int
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])
	current := make([]int, col)
	for i := 0; i < len(matrix[0]); i++ {
		current[i] = int(matrix[0][i] - '0')
	}

	max = maxHistogram(current)
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				current[j] = current[j] + 1
			} else {
				current[j] = 0
			}
		}

		tmp := maxHistogram(current)
		if tmp > max {
			max = tmp
		}
	}

	return max
}

// Largest Rectangular Area in a Histogram
func maxHistogram(hist []int) int {
	stack := make([]int, 0)
	var maxArea int

	var i int
	for i < len(hist) {
		if len(stack) == 0 || hist[stack[len(stack)-1]] < hist[i] {
			stack = append(stack, i)
			i++
		} else {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop the top

			var area int
			if len(stack) == 0 {
				area = int(hist[topIndex]) * i
			} else {
				area = int(hist[topIndex]) * (i - stack[len(stack)-1] - 1)
			}

			if area > maxArea {
				maxArea = area
			}
		}
	}

	for len(stack) > 0 {
		topIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // pop the top

		var area int
		if len(stack) == 0 {
			area = int(hist[topIndex]) * i
		} else {
			area = int(hist[topIndex]) * (i - stack[len(stack)-1] - 1)
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func main() {
	var a = [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	// var a = [][]byte{{'1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '0', '0', '0'}, {'0', '1', '1', '1', '1', '0', '0', '0'}}
	fmt.Println(maximalRectangle(a))
}
