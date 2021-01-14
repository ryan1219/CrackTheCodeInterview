package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {
	var max int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

		}
	}

	return max
}

func main() {
	// var a = [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	var a = [][]byte{{'1', '1', '1', '1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '1', '1', '0'}, {'1', '1', '1', '1', '1', '0', '0', '0'}, {'0', '1', '1', '1', '1', '0', '0', '0'}}
	fmt.Println(maximalRectangle(a))
}
