package main

// question: https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, max, min int) int {
	if root == nil {
		return max - min
	}

	max = maxInt(max, root.Val)
	min = minInt(min, root.Val)

	return maxInt(dfs(root.Left, max, min), dfs(root.Right, max, min))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
