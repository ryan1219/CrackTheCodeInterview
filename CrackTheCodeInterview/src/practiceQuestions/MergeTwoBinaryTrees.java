package practiceQuestions;

/*
 * question: https://leetcode.com/problems/merge-two-binary-trees/
 */
public class MergeTwoBinaryTrees {

	public TreeNode mergeTrees(TreeNode t1, TreeNode t2) {
		if(t2 == null && t1 != null) {
			t2 = new TreeNode(0);
		}
		inorder(t1, t2);
		return t2;
	}
	
	void inorder(TreeNode t1, TreeNode t2) {
		if(t1 == null) {
			return;
		}
		
		t2.val = t2.val + t1.val;
		if(t1.left != null && t2.left == null) {
			t2.left = new TreeNode(0);
		}
		if(t1.right != null && t2.right == null) {
			t2.right = new TreeNode(0);
		}
		inorder(t1.left, t2.left);
		inorder(t1.right, t2.right);
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
