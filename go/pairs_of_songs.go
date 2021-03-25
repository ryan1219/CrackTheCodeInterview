package main

// question: https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
/*
for each t in time []int, how many x, so that (t+x)%60=0
*/
func numPairsDivisibleBy60(time []int) int {
	remainder := make([]int, 60)
	res := 0

	for _, t := range time {
		res += remainder[(60-t%60)%60]
		remainder[t%60]++
	}

	return res
}
