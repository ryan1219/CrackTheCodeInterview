package main

// question: https://leetcode.com/problems/find-the-town-judge/
func findJudge(N int, trust [][]int) int {
	network := make([]int, N+1)

	for _, v := range trust {
		network[v[0]]--
		network[v[1]]++
	}

	for i := 1; i <= N; i++ {
		if network[i] == N-1 {
			return i
		}
	}
	return -1
}
