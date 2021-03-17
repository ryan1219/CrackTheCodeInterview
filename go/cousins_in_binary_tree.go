package main

// question: https://leetcode.com/problems/cousins-in-binary-tree/
func isCousins(root *TreeNode, x int, y int) bool {
	xDepth, xParent := dfs(root, x)
	yDepth, yParent := dfs(root, y)

	return xDepth == yDepth && xParent.Val != yParent.Val
}

func dfs(root *TreeNode, target int) (int, *TreeNode) {
	if root == nil {
		return -1, nil
	}

	if root.Val == target {
		return 0, nil
	}

	leftDepth, leftParent := dfs(root.Left, target)
	rightDepth, rightParent := dfs(root.Right, target)
	if leftDepth == 0 {
		return leftDepth + 1, root
	}
	if leftDepth > 0 {
		return leftDepth + 1, leftParent
	}
	if rightDepth == 0 {
		return rightDepth + 1, root
	}
	if rightDepth > 0 {
		return rightDepth + 1, rightParent
	}

	return -1, nil
}
