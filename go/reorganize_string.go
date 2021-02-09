package main

import (
	"container/heap"
)

// question: https://leetcode.com/problems/reorganize-string/
func reorganizeString(S string) string {
	m := make(map[rune]int)
	res := ""

	for _, ch := range S {
		val, contains := m[ch]
		if contains {
			m[ch] = val + 1
		} else {
			m[ch] = 1
		}

		if m[ch] > (len(S)+1)/2 {
			return res
		}
	}

	pq := &PriorityQueue{}

	for k, v := range m {
		heap.Push(pq, Element{v, k})
	}

	for pq.Len() != 0 {
		top := heap.Pop(pq).(Element)

		if len(res) == 0 || res[len(res)-1] != byte(top.value) {
			res += string(top.value)
			if top.priority-1 > 0 {
				top.priority--
				heap.Push(pq, top)
			}
		} else {
			second := heap.Pop(pq).(Element)
			res += string(second.value)
			if second.priority-1 > 0 {
				second.priority--
				heap.Push(pq, second)
			}

			heap.Push(pq, top)
		}
	}

	return res
}

type Element struct {
	priority int
	value    rune
}

type PriorityQueue []Element

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Element)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}
