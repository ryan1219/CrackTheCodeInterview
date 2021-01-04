package main

import "math"

/**
* question: https://leetcode.com/problems/binary-tree-maximum-path-sum/
* path is from one node to another node in the tree, max sum is the sum of all nodes along the path which is the max
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var max int = math.MinInt64
	var maxPointer = &max
	recursion(root, maxPointer)
	return max
}

func recursion(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := maxOf(recursion(root.Left, max), 0)
	right := maxOf(recursion(root.Right, max), 0)

	*max = maxOf(*max, left+right+root.Val)

	return maxOf(left+root.Val, right+root.Val)
}

func maxOf(a, b int) int {
	if a < b {
		return b
	}
	return a
}
