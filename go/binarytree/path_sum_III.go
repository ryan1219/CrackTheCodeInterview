package binarytree

// question: https://leetcode.com/problems/path-sum-iii/
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return pathSumFrom(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumFrom(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	if sum == root.Val {
		return 1 + pathSumFrom(root.Left, sum-root.Val) + pathSumFrom(root.Right, sum-root.Val)
	}

	return pathSumFrom(root.Left, sum-root.Val) + pathSumFrom(root.Right, sum-root.Val)
}
