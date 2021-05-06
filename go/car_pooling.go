package main

// question: https://leetcode.com/problems/car-pooling/
func carPooling(trips [][]int, capacity int) bool {
	stops := make([]int, 1001)
	for _, trip := range trips {
		stops[trip[1]] += trip[0]
		stops[trip[2]] -= trip[0]
	}
	for i := 0; capacity >= 0 && i < len(stops); i++ {
		capacity -= stops[i]
	}

	return capacity >= 0
}
