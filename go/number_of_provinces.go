package main

// question: https://leetcode.com/problems/number-of-provinces/
// bfs finds connected component in a graph
func findCircleNum(isConnected [][]int) int {
	connected := 0
	visited := make([]int, len(isConnected))
	for i := 0; i < len(visited); i++ {
		if visited[i] == 0 {
			bfs(isConnected, visited, i)
			connected++
		}
	}

	return connected
}

func bfs(isConnected [][]int, visited []int, start int) {
	visited[start] = 1
	list := make([]int, 0)
	list = append(list, start)

	for len(list) > 0 {
		current := list[0]
		visited[current] = 1
		list = list[1:]

		for j, val := range isConnected[current] {
			if val == 1 && visited[j] == 0 {
				list = append(list, j)
			}
		}
	}
}
