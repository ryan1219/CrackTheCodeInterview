package main

// question: https://leetcode.com/problems/shortest-path-in-binary-matrix/
/*
why dp won't work? because for each position there are 8 possible directions to reach it
*/
func shortestPathBinaryMatrix(grid [][]int) int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, -1}, {-1, 1}, {-1, -1}, {1, 1}}

	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}

	visited := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[0]))
	}
	visited[0][0] = 1
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0})
	res := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == len(grid)-1 && cur[1] == len(grid[0])-1 {
				return res + 1
			}

			for j := 0; j < len(directions); j++ {
				nextX := cur[0] + directions[j][0]
				nextY := cur[1] + directions[j][1]

				if nextX >= 0 && nextX < len(grid) && nextY >= 0 && nextY < len(grid[0]) && visited[nextX][nextY] == 0 && grid[nextX][nextY] == 0 {
					queue = append(queue, []int{nextX, nextY})
					visited[nextX][nextY] = 1
				}
			}
		}
		res++
	}

	return -1
}
