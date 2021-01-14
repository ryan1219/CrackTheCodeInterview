package binarytree

/**
 * question: https://leetcode.com/problems/minimum-depth-of-binary-tree
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	depth := 1
	for len(queue) > 0 {
		current := len(queue)
		for i := 0; i < current; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

// DFS
// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}

// 	leftDepth := minDepth(root.Left)
// 	rightDepth := minDepth(root.Right)

// 	if leftDepth == 0 && rightDepth != 0 {
// 		return rightDepth + 1
// 	}

// 	if leftDepth != 0 && rightDepth == 0 {
// 		return leftDepth + 1
// 	}

// 	return Min(leftDepth, rightDepth) + 1
// }

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
