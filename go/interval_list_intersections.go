package main

// question: https://leetcode.com/problems/interval-list-intersections/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)

	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		if start <= end {
			res = append(res, []int{start, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
