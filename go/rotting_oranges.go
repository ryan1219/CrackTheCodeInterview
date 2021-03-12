package main

import "fmt"

// question: https://leetcode.com/problems/rotting-oranges/
func orangesRotting(grid [][]int) int {
	breath := 0
	queue := make([]point, 0)
	seen := make([][]int, len(grid))
	for i := 0; i < len(seen); i++ {
		seen[i] = make([]int, len(grid[i]))
	}

	// find all rotten oranges at the beginning
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 2 {
				seen[x][y] = 1
				if x+1 < len(grid) && grid[x+1][y] == 1 {
					queue = append(queue, point{x + 1, y})
					seen[x+1][y] = 1
				}
				if x-1 >= 0 && grid[x-1][y] == 1 {
					queue = append(queue, point{x - 1, y})
					seen[x-1][y] = 1
				}
				if y+1 < len(grid[x]) && grid[x][y+1] == 1 {
					queue = append(queue, point{x, y + 1})
					seen[x][y+1] = 1
				}
				if y-1 >= 0 && grid[x][y-1] == 1 {
					queue = append(queue, point{x, y - 1})
					seen[x][y-1] = 1
				}
			}
		}
	}
	fmt.Println(queue)

	for len(queue) > 0 {
		level := len(queue)
		for level > 0 {
			p := queue[0]
			queue = queue[1:]
			grid[p.x][p.y] = 2
			if p.x+1 < len(grid) && grid[p.x+1][p.y] == 1 && seen[p.x+1][p.y] == 0 {
				queue = append(queue, point{p.x + 1, p.y})
				seen[p.x+1][p.y] = 1
			}
			if p.x-1 >= 0 && grid[p.x-1][p.y] == 1 && seen[p.x-1][p.y] == 0 {
				queue = append(queue, point{p.x - 1, p.y})
				seen[p.x-1][p.y] = 1
			}
			if p.y+1 < len(grid[p.x]) && grid[p.x][p.y+1] == 1 && seen[p.x][p.y+1] == 0 {
				queue = append(queue, point{p.x, p.y + 1})
				seen[p.x][p.y+1] = 1
			}
			if p.y-1 >= 0 && grid[p.x][p.y-1] == 1 && seen[p.x][p.y-1] == 0 {
				queue = append(queue, point{p.x, p.y - 1})
				seen[p.x][p.y-1] = 1
			}
			level--
		}
		breath++
	}

	// in the end, search if there is fresh orange left
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return breath
}

type point struct {
	x int
	y int
}

func main() {
	// fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	// fmt.Println(orangesRotting([][]int{{0, 2}}))
	// fmt.Println(orangesRotting([][]int{{0}}))
	fmt.Println(orangesRotting([][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}}))
}
