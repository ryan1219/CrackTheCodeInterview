package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//iterative approach
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	var previous *ListNode = nil
	for current.Next != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	current.Next = previous
	return current
}
