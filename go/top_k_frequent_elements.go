package main

import (
	"sort"
)

type pair struct {
	Key   int
	Value int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Less(i, j int) bool {
	return (*p)[i].Value < (*p)[j].Value
}

func (p *pairList) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	p := make(pairList, 0)
	result := make([]int, 0)

	for _, number := range nums {
		val, ok := m[number]
		if ok {
			m[number] = val + 1
		} else {
			m[number] = 1
		}
	}

	for k, v := range m {
		p = append(p, pair{k, v})
	}
	sort.Sort(sort.Reverse(&p))
	for k > 0 {
		result = append(result, p[k-1].Key)
		k--
	}

	return result
}
