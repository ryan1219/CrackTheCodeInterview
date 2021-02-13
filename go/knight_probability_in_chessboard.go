package main

// question: https://leetcode.com/problems/knight-probability-in-chessboard/
// recursive solution
var directions = [][]int{{-2, -1}, {-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}}

func knightProbability(N int, K int, r int, c int) float64 {
	ways := make([][][]float64, N)

	for i := 0; i < len(ways); i++ {
		ways[i] = make([][]float64, N)
	}

	for i := 0; i < len(ways); i++ {
		for j := 0; j < len(ways[i]); j++ {
			ways[i][j] = make([]float64, K+1)
		}
	}

	ways[r][c][0] = 1
	for k := 1; k <= K; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for _, dir := range directions {
					rp := i - dir[0]
					rc := j - dir[1]

					if 0 <= rp && rp < N && 0 <= rc && rc < N {
						ways[i][j][k] += ways[rp][rc][k-1] / 8.0
					}
				}
			}
		}
	}

	var res float64
	for i := 0; i < len(ways); i++ {
		for j := 0; j < len(ways[i]); j++ {
			res += ways[i][j][K]
		}
	}

	return res
}

func knightProbabilityRecursive(N int, K int, r int, c int) float64 {
	if r < 0 || r > N-1 || c < 0 || c > N-1 {
		return 0
	}
	if K == 0 {
		return 1
	}

	var rate float64
	for _, dir := range directions {
		rate += 0.125 * knightProbabilityRecursive(N, K-1, r+dir[0], c+dir[1])
	}

	return rate
}
