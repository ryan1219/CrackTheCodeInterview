package practiceQuestions;

import java.util.LinkedList;
import java.util.Queue;

/*
 * question: https://leetcode.com/problems/maximum-depth-of-binary-tree/
 * 
 * 
 */

public class MaximumDepthOfBinaryTree {

	class TreeNode {
		int val;
		TreeNode left;
		TreeNode right;

		TreeNode(int x) {
			this.val = x;
		}
	}

	// solution1: BFS
	public int maxDepth(TreeNode root) {
		if(root == null) {
			return 0;
		}
		Queue<TreeNode> queue = new LinkedList<>();
		int depth = 0;
		int currentLevel = 0;
		queue.add(root);
		while (!queue.isEmpty()) {
			currentLevel = queue.size();
			while (currentLevel > 0) {
				TreeNode temp = queue.remove();
				if (temp.left != null) {
					queue.add(temp.left);
				}
				if (temp.right != null) {
					queue.add(temp.right);
				}
				currentLevel--;
			}
			depth++;
		}
		return depth;
	}
	
	/* solution2: recursive approach
	 * 
	 * base case: when node == null; depth = 0
	 * recursive case: node height = 1 + max(height of left subtree, height of right subtree)
	*/
	
//	public int maxDepth(TreeNode root) {
//		if(root == null) {
//			return 0;
//		}
//		return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
//	}

}
