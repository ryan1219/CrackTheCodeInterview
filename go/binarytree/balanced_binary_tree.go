package binarytree

import "math"

// question: https://leetcode.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := height(root.Left)
	r := height(root.Right)

	if int(math.Abs(float64(l-r))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(height(node.Left), height(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
