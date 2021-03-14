package main

import (
	"strconv"
	"strings"
)

// question: https://leetcode.com/problems/sliding-puzzle/
func slidingPuzzle(board [][]int) int {
	target := "123450"
	start := ""

	// find the start status
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			start += strconv.Itoa(board[i][j])
		}
	}

	visited := make(map[string]int)

	dirs := [][]int{{1, 3}, {0, 2, 4},
		{1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

	queue := make([]string, 0)
	queue = append(queue, start)
	visited[start] = 1

	res := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return res
			}
			zeroIndex := strings.Index(cur, "0")
			// find all possible next status
			for _, dir := range dirs[zeroIndex] {
				next := swap(cur, zeroIndex, dir)
				if visited[next] == 1 {
					continue
				}
				visited[next] = 1
				queue = append(queue, next)
			}
		}
		res++
	}
	return -1
}

func swap(str string, i, j int) string {
	r := []rune(str)
	r[i], r[j] = r[j], r[i]
	return string(r)
}
