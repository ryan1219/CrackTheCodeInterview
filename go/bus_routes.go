package main

import "container/list"

// question: https://leetcode.com/problems/bus-routes/
func numBusesToDestination(routes [][]int, source int, target int) int {
	m := make(map[int][]int)
	visited := make(map[int]int)
	queue := list.New()
	buses := 0
	if source == target {
		return 0
	}

	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			val, contains := m[routes[i][j]]
			if contains {
				val = append(val, i)
			} else {
				val = []int{i}
			}
			m[routes[i][j]] = val
		}
	}

	queue.PushFront(source)
	for queue.Len() > 0 {
		length := queue.Len()
		buses++
		for i := 0; i < length; i++ {
			stop := queue.Front()
			queue.Remove(stop)
			for _, bus := range m[stop.Value.(int)] {
				_, contains := visited[bus]
				if contains {
					continue
				}
				visited[bus] = 1
				for j := 0; j < len(routes[bus]); j++ {
					if routes[bus][j] == target {
						return buses
					}
					queue.PushBack(routes[bus][j])
				}
			}
		}
	}
	return -1
}
