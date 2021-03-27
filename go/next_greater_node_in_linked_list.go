package main

// https://leetcode.com/problems/next-greater-node-in-linked-list/
func nextLargerNodes(head *ListNode) []int {
	res := make([]int, 0)

	current := head
	for current != nil {
		next_larger := 0
		start := current.Next

		for start != nil {
			if start.Val > current.Val {
				next_larger = start.Val
				break
			}
			start = start.Next
		}

		res = append(res, next_larger)
		current = current.Next
	}

	return res
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
