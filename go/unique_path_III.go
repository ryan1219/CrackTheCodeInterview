package main

/*
question: https://leetcode.com/problems/unique-paths-iii/
1 represents the starting square.  There is exactly one starting square.
2 represents the ending square.  There is exactly one ending square.
0 represents empty squares we can walk over.
-1 represents obstacles that we cannot walk over.
*/
func uniquePathsIII(grid [][]int) int {
	startX, startY, empty, res := 0, 0, 1, 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				empty++
			}
			if grid[i][j] == 1 {
				startX = i
				startY = j
			}
		}
	}

	dfs(startX, startY, empty, grid, &res)

	return res
}

func dfs(x, y, empty int, grid [][]int, res *int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] < 0 {
		return
	}
	if grid[x][y] == 2 {
		if empty == 0 {
			*res++
		}
		return
	}
	grid[x][y] = -2
	empty--
	dfs(x+1, y, empty, grid, res)
	dfs(x-1, y, empty, grid, res)
	dfs(x, y+1, empty, grid, res)
	dfs(x, y-1, empty, grid, res)
	empty++
	grid[x][y] = 0
}
