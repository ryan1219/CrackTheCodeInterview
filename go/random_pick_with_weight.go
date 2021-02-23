package main

import (
	"math/rand"
	"time"
)

// question: https://leetcode.com/problems/random-pick-with-weight/
type Solution struct {
	slots []int
}

func Constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	slots := make([]int, len(w))
	slots[0] = w[0]
	for i := 1; i < len(slots); i++ {
		slots[i] = slots[i-1] + w[i]
	}

	return Solution{slots}
}

func (this *Solution) PickIndex() int {
	max := this.slots[len(this.slots)-1]
	num := rand.Intn(max) // [0, max)

	l, r := 0, len(this.slots)-1
	for l <= r {
		m := (l + r) >> 1
		if this.slots[m] > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
