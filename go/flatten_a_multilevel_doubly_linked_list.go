package main

// question: https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	head := root
	cur := root

	for cur != nil {
		if cur.Child != nil {
			flattenSubLevel(cur)
		}
		cur = cur.Next
	}

	return head
}

func flattenSubLevel(root *Node) *Node {
	head := root
	next := root.Next
	cur := root.Child

	// for each sub level, change the head node
	head.Child = nil
	head.Next = cur
	cur.Prev = head

	for cur.Next != nil {
		if cur.Child != nil {
			flattenSubLevel(cur)
		}
		cur = cur.Next
	}

	// link the tail of sub level to the upper level
	cur.Next = next
	if next != nil {
		next.Prev = cur
	}
	return head
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
