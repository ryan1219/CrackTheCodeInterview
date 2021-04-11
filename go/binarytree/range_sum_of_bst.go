package binarytree

// question: https://leetcode.com/problems/range-sum-of-bst/
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0

	preorder(root, low, high, &res)
	return res
}

func preorder(root *TreeNode, low int, high int, res *int) {
	if root == nil {
		return
	}

	if low <= root.Val && root.Val <= high {
		*res += root.Val
	}
	preorder(root.Left, low, high, res)
	preorder(root.Right, low, high, res)
}
