package main

// question: https://leetcode.com/problems/frog-jump/
func canCross(stones []int) bool {
	stonesMap := make(map[int][]int)
	for _, stone := range stones {
		stonesMap[stone] = make([]int, 0)
	}

	stonesMap[0] = append(stonesMap[0], 1)

	for _, stone := range stones {
		steps := stonesMap[stone]
		for _, step := range steps {
			nextStone := stone + step
			if _, ok := stonesMap[nextStone]; ok {
				if nextStone == stones[len(stones)-1] {
					return true
				}
				if !contains(stonesMap[nextStone], step-1) {
					stonesMap[nextStone] = append(stonesMap[nextStone], step-1)
				}
				if !contains(stonesMap[nextStone], step) {
					stonesMap[nextStone] = append(stonesMap[nextStone], step)
				}
				if !contains(stonesMap[nextStone], step+1) {
					stonesMap[nextStone] = append(stonesMap[nextStone], step+1)
				}
			}
		}
	}

	return false
}

// contains checks if an int is preset in a slice
func contains(s []int, num int) bool {
	for _, v := range s {
		if v == num {
			return true
		}
	}

	return false
}
