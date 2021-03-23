package main

// question: https://leetcode.com/problems/sum-of-left-leaves/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	return traverse(root, 0)
}

func traverse(root *TreeNode, left int) int {
	if root.Left == nil && root.Right == nil && left == 1 {
		return root.Val
	}

	if root.Left == nil && root.Right == nil && left == 0 {
		return 0
	}

	leftValue, rightValue := 0, 0
	if root.Left != nil {
		leftValue = traverse(root.Left, 1)
	}
	if root.Right != nil {
		rightValue = traverse(root.Right, 0)
	}

	return leftValue + rightValue
}
