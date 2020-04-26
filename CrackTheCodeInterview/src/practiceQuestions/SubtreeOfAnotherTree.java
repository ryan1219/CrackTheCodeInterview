package practiceQuestions;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/*
 * question: https://leetcode.com/problems/subtree-of-another-tree/
 */
public class SubtreeOfAnotherTree {

	/*
	 * solution1: postorder traversal for both BST s, t, and get two list a and b,
	 * check if b is sublist of a
	 * 
	 * preorder traversal also works
	 */

	public boolean isSubtree(TreeNode s, TreeNode t) {
		List<Integer> source = postorderTraversal(s);
		List<Integer> target = postorderTraversal(t);
		int index = Collections.indexOfSubList(source, target);
		return index >= 0;
	}

	public List<Integer> postorderTraversal(TreeNode root) {
		List<Integer> result = new ArrayList<>();
		postorder(root, result);
		return result;
	}

	public void postorder(TreeNode root, List<Integer> list) {
		if (root == null) {
			list.add(null);
			return;
		}
		postorder(root.left, list);
		postorder(root.right, list);
		list.add(root.val);
	}

	class TreeNode {
		int val;
		TreeNode left;
		TreeNode right;

		TreeNode(int x) {
			this.val = x;
		}
	}
}
