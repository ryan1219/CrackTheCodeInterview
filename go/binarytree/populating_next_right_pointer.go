package binarytree

// question: https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	var current *Node
	var head *Node
	var pre *Node

	head = root

	for head != nil {
		current = head
		pre = nil
		head = nil

		// on each level
		for current != nil {
			if current.Left != nil {
				if pre == nil {
					pre = current.Left
					head = current.Left
				} else {
					pre.Next = current.Left
					pre = pre.Next
				}
			}
			if current.Right != nil {
				if pre == nil {
					pre = current.Right
					head = current.Right
				} else {
					pre.Next = current.Right
					pre = pre.Next
				}
			}
			current = current.Next
		}
	}

	return root
}
