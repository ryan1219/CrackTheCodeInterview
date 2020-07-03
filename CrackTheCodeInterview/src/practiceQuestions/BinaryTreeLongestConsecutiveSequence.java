package practiceQuestions;

import java.util.LinkedList;
import java.util.Queue;

public class BinaryTreeLongestConsecutiveSequence {
	public int longestConsective(TreeNode root) {
		if (root == null) {
			return 0;
		}

		Queue<TreeNode> nodes = new LinkedList<>();
		Queue<Integer> length = new LinkedList<>();

		nodes.add(root);
		length.add(1);
		int max = 1;

		while (nodes.size() != 0) {
			TreeNode current = nodes.poll();
			int currentLength = length.poll();

			if (current.left != null) {
				int leftLength = currentLength;
				if (current.left.val == current.val + 1) {
					leftLength += 1;
					max = Math.max(max, leftLength);
				} else {
					leftLength = 1;
				}
				nodes.add(current.left);
				length.add(leftLength);
			}

			if (current.right != null) {
				int rightLength = currentLength;
				if (current.right.val == current.val + 1) {
					rightLength += 1;
					max = Math.max(max, rightLength);
				} else {
					rightLength = 1;
				}
				nodes.add(current.right);
				length.add(rightLength);
			}
		}
		return max;
	}

	class TreeNode {
		int val;
		TreeNode left;
		TreeNode right;

		TreeNode() {
		}

		TreeNode(int val) {
			this.val = val;
		}

		TreeNode(int val, TreeNode left, TreeNode right) {
			this.val = val;
			this.left = left;
			this.right = right;
		}
	}
}
