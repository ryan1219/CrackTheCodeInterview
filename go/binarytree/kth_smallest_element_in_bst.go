package binarytree

// question:
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	list := make([]int, 0)
	inorder(root, &list, k)

	return list[len(list)-1]
}

func inorder(root *TreeNode, list *[]int, k int) {
	if root == nil {
		return
	}

	inorder(root.Left, list, k)
	if len(*list) == k {
		return
	}
	*list = append(*list, root.Val)
	inorder(root.Right, list, k)
}
