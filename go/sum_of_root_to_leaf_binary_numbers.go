package main

import "strconv"

// question: https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
func sumRootToLeaf(root *TreeNode) int {
	res := 0
	traverse(root, "0", &res)
	return res
}

func traverse(node *TreeNode, prefix string, res *int) {
	if node.Left == nil && node.Right == nil {
		val, _ := strconv.ParseInt(prefix+strconv.Itoa(node.Val), 2, 0)
		*res += int(val)
		return
	}

	if node.Left != nil {
		traverse(node.Left, prefix+strconv.Itoa(node.Val), res)
	}
	if node.Right != nil {
		traverse(node.Right, prefix+strconv.Itoa(node.Val), res)
	}
}
