package main

import "sort"

type Point struct {
	x   int
	y   int
	dis int
}

// question: https://leetcode.com/problems/k-closest-points-to-origin/
func kClosest(points [][]int, K int) [][]int {
	pointsList := make([]Point, 0)
	res := make([][]int, 0)

	for _, point := range points {
		pointsList = append(pointsList, Point{point[0], point[1], point[0]*point[0] + point[1]*point[1]})
	}

	sort.Slice(pointsList, func(i, j int) bool {
		return pointsList[i].dis < pointsList[j].dis
	})

	count := 0
	for count < K {
		res = append(res, []int{pointsList[count].x, pointsList[count].y})
		count++
	}

	return res
}
