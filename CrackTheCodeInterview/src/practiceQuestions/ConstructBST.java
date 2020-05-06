package practiceQuestions;

/*
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 * 
 * question: 
 */
public class ConstructBST {

	public static void main(String[] args) {
		int[] inorder = { 9, 3, 15, 20, 7 };
		int[] postorder = { 9, 15, 7, 20, 3 };

		new ConstructBST().buildTree2(inorder, postorder);
	}

	/*
	 * solution1: two pointers for inorder array, two pointers for postorder array
	 */
	public TreeNode buildTree(int[] inorder, int[] postorder) {
		int start = 0;
		int end = inorder.length - 1;

		if (start > end) {
			return null;
		}

		return build(inorder, postorder, start, end, start, end);
	}

	TreeNode build(int[] inorder, int[] postorder, int inorderStart, int inorderEnd, int postStart, int postEnd) {
		if (inorderStart > inorderEnd) {
			return null;
		}
		TreeNode node = new TreeNode(postorder[postEnd]);
//		base case
		if (inorderStart == inorderEnd) {
			return node;
		}

		int index = search(inorder, postorder[postEnd]);
		node.left = build(inorder, postorder, inorderStart, index - 1, postStart, postStart + index - inorderStart - 1);
		node.right = build(inorder, postorder, index + 1, inorderEnd, postStart + index - inorderStart, postEnd - 1);
		return node;
	}

	/*
	 * solution2: two pointers for inorder array, one pointer for postorder array
	 * but always do right subtree first, the last element in postorder is the root
	 * of right subtree
	 */
	int postEnd;

	public TreeNode buildTree2(int[] inorder, int[] postorder) {
		int start = 0;
		int end = inorder.length - 1;
		postEnd = end;

		if (start > end) {
			return null;
		}

		return build(inorder, postorder, start, end);
	}

	TreeNode build(int[] inorder, int[] postorder, int inorderStart, int inorderEnd) {
		if (inorderStart > inorderEnd) {
			return null;
		}

		TreeNode node = new TreeNode(postorder[postEnd]);
		if (inorderStart == inorderEnd) {
			postEnd--;
			return node;
		}
		int index = search(inorder, postorder[postEnd]);
		postEnd--;
		node.right = build(inorder, postorder, index + 1, inorderEnd);
		node.left = build(inorder, postorder, inorderStart, index - 1);
		return node;
	}

	int search(int[] source, int target) {
		for (int i = 0; i < source.length; i++) {
			if (source[i] == target) {
				return i;
			}
		}
		return -1;
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
