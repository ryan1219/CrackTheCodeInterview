package binarytree

// question: https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
/*
binary search tree inorder traversal give acccending order, 1->2->3->4
if change the inorder traversal to right->current->left, get 4->3->2->1, accumulate the value from left and add to the right
*/
func bstToGst(root *TreeNode) *TreeNode {
	var sum int = 0
	reversedInorder(root, &sum)
	return root
}

func reversedInorder(root *TreeNode, sum *int) {
	if root.Right != nil {
		reversedInorder(root.Right, sum)
	}
	*sum += root.Val
	root.Val = *sum
	if root.Left != nil {
		reversedInorder(root.Left, sum)
	}
}
