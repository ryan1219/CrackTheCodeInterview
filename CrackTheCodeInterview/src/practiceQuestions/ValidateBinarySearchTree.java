package practiceQuestions;

import java.util.ArrayList;

/*
 * question: https://leetcode.com/problems/validate-binary-search-tree/
 */
public class ValidateBinarySearchTree {

	/*
	 * solution1: if a BST is valid, inorder traversal gives a ascending order
	 */
	ArrayList<Integer> order = new ArrayList<>();

	public boolean isValidBST(TreeNode root) {
		inorder(root);
		if(order.isEmpty()) {
			return true;
		}
		int previous = order.get(0);
		for(int i = 1; i < order.size(); i++) {
			if(order.get(i)> previous) {
				previous = order.get(i);
			}else {
				return false;
			}
		}
		return true;
	}

	public void inorder(TreeNode root) {
		if (root == null) {
			return;
		}
		inorder(root.left);
		order.add(root.val);
		inorder(root.right);
	}

	class TreeNode {
		int val;
		TreeNode left;
		TreeNode right;

		public TreeNode(int value) {
			this.val = value;
		}
	}
}
