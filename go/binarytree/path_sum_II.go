package binarytree

// question: https://leetcode.com/problems/path-sum-ii/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)

	search(root, sum, make([]int, 0), &result)

	return result
}

func search(root *TreeNode, sum int, path []int, result *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if sum-root.Val == 0 {
			cpy := make([]int, len(path))
			copy(cpy, path)
			cpy = append(cpy, root.Val)
			(*result) = append(*result, cpy)
		}
		return
	}

	path = append(path, root.Val)
	search(root.Left, sum-root.Val, path, result)
	search(root.Right, sum-root.Val, path, result)
}
