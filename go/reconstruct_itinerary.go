package main

import "sort"

/*
question: https://leetcode.com/problems/reconstruct-itinerary/

related question: Find Itinerary from a given list of tickets
question: https://www.geeksforgeeks.org/find-itinerary-from-a-given-list-of-tickets/
in order to find the starting point, compare the map with the reversed map,
if the key in map is not in the reversed map, this key is the starting point

Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge u v,
vertex u comes before v in the ordering. Topological Sorting for a graph is not possible if the graph is not a DAG.
*/

// Hierholzer’s algorithm to find a Eulerian path
func findItinerary(tickets [][]string) []string {
	path := make([]string, 0)
	stack := make([]string, 0)
	graph := make(map[string][]string)

	for _, ticket := range tickets {
		_, contains := graph[ticket[0]]
		if !contains {
			graph[ticket[0]] = []string{ticket[1]}
		} else {
			graph[ticket[0]] = append(graph[ticket[0]], ticket[1])
		}
	}
	// sort each value in lexical order
	for key, value := range graph {
		sort.Strings(value)
		graph[key] = value
	}

	stack = append(stack, "JFK")

	for len(stack) > 0 {
		val, exists := graph[stack[len(stack)-1]]
		for exists && len(val) > 0 {
			graph[stack[len(stack)-1]] = val[1:]
			stack = append(stack, val[0])
			val, exists = graph[stack[len(stack)-1]]
		}
		path = prepend(path, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return path
}

func prepend(src []string, value string) []string {
	if len(src) == 0 {
		return append(src, value)
	}
	src = append(src, "")
	copy(src[1:], src)
	src[0] = value
	return src
}
