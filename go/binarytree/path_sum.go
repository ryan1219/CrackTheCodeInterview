package binarytree

func hasPathSum(root *TreeNode, sum int) bool {
	result := false
	find := &result

	inorder(root, 0, sum, find)
	return *find
}

func inorder(root *TreeNode, sum, target int, find *bool) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if sum+root.Val == target {
			*find = true
		}
	}

	inorder(root.Left, sum+root.Val, target, find)
	inorder(root.Right, sum+root.Val, target, find)
}
