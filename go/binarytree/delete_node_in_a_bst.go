package binarytree

// question: https://leetcode.com/problems/delete-node-in-a-bst/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
 case1: node to be deleted is leaf: simply remove from the tree
 case2: node to be deleted has one child: copy the child to the node and delete the child
 case3: node has two childre: find predecessor in the inorder, copy the value of predecessor and delete the predecessor
 note: inorder predecessor = largest value in the left subtree = right most position of the left subtree
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// case1 and case2
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// case3: inorder predecessor -> largest in the left subtree
		root.Val = maxValue(*root.Left)
		root.Left = deleteNode(root.Left, root.Val)
	}

	return root
}

func maxValue(root TreeNode) int {
	max := root.Val
	for root.Right != nil {
		max = root.Right.Val
		root = *root.Right
	}
	return max
}
