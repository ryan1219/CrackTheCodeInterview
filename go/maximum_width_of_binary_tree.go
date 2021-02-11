package main

// question: https://leetcode.com/problems/maximum-width-of-binary-tree/
/*
use binary tree index - array implementation indexing, to represent the position for node
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	TreeNode *TreeNode
	Pos      int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := make([]Node, 0)

	queue = append(queue, Node{TreeNode: root, Pos: 1})

	max := 0
	start := 0
	end := len(queue) - 1
	for start < len(queue) {
		first := queue[start]
		last := first

		for i := start; i <= end; i++ {
			last = queue[i]
			if last.TreeNode.Left != nil {
				queue = append(queue, Node{last.TreeNode.Left, last.Pos * 2})
			}
			if last.TreeNode.Right != nil {
				queue = append(queue, Node{last.TreeNode.Right, last.Pos*2 + 1})
			}
		}

		if last.Pos-first.Pos+1 > max {
			max = last.Pos - first.Pos + 1
		}

		start = end + 1
		end = len(queue) - 1
	}

	return max
}
