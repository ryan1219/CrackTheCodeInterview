package main

// question: https://leetcode.com/problems/single-number-iii/
func singleNumber(nums []int) []int {
	single := make(map[int]int)
	res := make([]int, 0)

	for _, num := range nums {
		if _, containes := single[num]; containes {
			delete(single, num)
		} else {
			single[num] = 1
		}
	}

	for k := range single {
		res = append(res, k)
	}

	return res
}
