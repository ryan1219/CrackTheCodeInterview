package priorityQueue

import (
	"container/heap"
	"fmt"
)

type Element struct {
	priority int
	value    string
}

type PriorityQueue []Element

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
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

func main() {
	pq := &PriorityQueue{
		{3, "Vacuum"},
		{2, "Feed cat"},
		{4, "Arrange play date"},
		{1, "Pack for trip"}}

	// heapify
	heap.Init(pq)

	// enqueue
	heap.Push(pq, Element{2, "Iron"})

	for pq.Len() != 0 {
		// dequeue
		fmt.Println(heap.Pop(pq))
		fmt.Println(pq)
	}
}
