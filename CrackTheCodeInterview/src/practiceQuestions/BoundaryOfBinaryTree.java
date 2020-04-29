package practiceQuestions;

import java.util.ArrayList;
import java.util.List;

/*
 * question: https://xiaoguan.gitbooks.io/leetcode/content/LeetCode/545-boundary-of-binary-tree-medium.html
 * 
 */
public class BoundaryOfBinaryTree {
	class TreeNode {
		int val;
		TreeNode left;
		TreeNode right;

		TreeNode(int x) {
			this.val = x;
			left = null;
			right = null;
		}
	}

	List<Integer> res = new ArrayList<>();
	List<Integer> right = new ArrayList<>();

	List<Integer> boundaryOfBinaryTree(TreeNode root) {
		res = new ArrayList<>();
		if (root == null) {
			return res;
		}
		if (root.left != null || root.right != null) {
			res.add(root.val);
		}

		visitLeftBoundary(root);
		visitLeaf(root);
		visitRightBoundary(root);

		for (int i = right.size() - 1; i >= 0; i--) {
			res.add(right.get(i));
		}
		return res;
	}

	void visitLeftBoundary(TreeNode root) {
		TreeNode current = root.left;
		while (current != null) {
			if (current.left != null || current.right != null) {
				res.add(current.val);
			}
			if (current.left != null) {
				current = current.left;
			} else {
				current = current.right;
			}

		}
	}

	void visitLeaf(TreeNode root) {
		if (root == null) {
			return;
		}

		if (root.left == null && root.right == null) {
			res.add(root.val);
		}

		visitLeaf(root.left);
		visitLeaf(root.right);
	}

	void visitRightBoundary(TreeNode root) {
		TreeNode current = root.right;
		while (current != null) {
			if (current.right != null || current.left != null) {
				right.add(current.val);
			}
			if (current.right != null) {
				current = current.right;
			} else {
				current = current.left;
			}

		}
	}

}
