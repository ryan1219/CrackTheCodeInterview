package main

// https://leetcode.com/problems/evaluate-division/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	for i := 0; i < len(values); i++ {
		_, contains := graph[equations[i][0]]
		if !contains {
			graph[equations[i][0]] = make(map[string]float64)
		}
		_, contains = graph[equations[i][1]]
		if !contains {
			graph[equations[i][1]] = make(map[string]float64)
		}

		graph[equations[i][0]][equations[i][1]] = float64(values[i])
		graph[equations[i][1]][equations[i][0]] = float64(1) / float64(values[i])
	}

	result := make([]float64, 0)
	for i := 0; i < len(queries); i++ {
		result = append(result, calcEquationDFS(queries[i][0], queries[i][1], 1, graph, make(map[string]int)))
	}

	return result
}

func calcEquationDFS(src, target string, weight float64, graph map[string]map[string]float64, seen map[string]int) float64 {
	if _, contained := graph[src]; !contained {
		return -1
	}
	if _, contained := seen[src]; contained {
		return -1
	}

	if src == target {
		return weight
	}

	seen[src] = 1
	for key := range graph[src] {
		res := calcEquationDFS(key, target, weight*graph[src][key], graph, seen)
		if res != -1 {
			return res
		}
	}

	return -1
}
