package binarytree

import "strconv"

func isSymmetric(root *TreeNode) bool {
	pre := ""
	post := ""

	preorder(root, &pre)
	postorder(root, &post)

	if pre == post {
		return true
	}

	return false
}

func preorder(root *TreeNode, result *string) {
	if root == nil {
		(*result) = (*result) + "null" + ","
		return
	}

	(*result) = (*result) + strconv.Itoa(root.Val) + ","
	preorder(root.Left, result)
	preorder(root.Right, result)
}

func postorder(root *TreeNode, result *string) {
	if root == nil {
		(*result) = (*result) + "null" + ","
		return
	}

	(*result) = (*result) + strconv.Itoa(root.Val) + ","
	postorder(root.Right, result)
	postorder(root.Left, result)
}
