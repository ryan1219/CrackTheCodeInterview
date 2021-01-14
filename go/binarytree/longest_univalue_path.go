package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result int = 0

func longestUnivaluePath(root *TreeNode) int {
	traverse(root)
	return result
}

func traverse(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftPath := traverse(root.Left)
	rightPath := traverse(root.Right)

	if root.Left != nil && root.Left.Val == root.Val {
		leftPath += 1
	} else {
		leftPath = 0
	}

	if root.Right != nil && root.Right.Val == root.Val {
		rightPath += 1
	} else {
		rightPath = 0
	}

	result = max(leftPath+rightPath, result)

	return max(leftPath, rightPath)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
