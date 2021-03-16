package main

import (
	"math/rand"
	"sort"
)

//question: https://leetcode.com/problems/random-pick-with-blacklist/
type Solution struct {
	upperBound int
	exclude    []int
}

func Constructor(N int, blacklist []int) Solution {
	sort.Ints(blacklist)
	return Solution{N, blacklist}
}

func (this *Solution) Pick() int {
	random := rand.Intn(this.upperBound - len(this.exclude))
	for _, v := range this.exclude {
		if random < v {
			break
		}
		random++
	}

	return random
}
