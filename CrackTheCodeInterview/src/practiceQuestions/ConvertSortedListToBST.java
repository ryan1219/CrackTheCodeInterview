package practiceQuestions;

import java.util.ArrayList;

/*
 * question: https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree/
 * 
 * Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
 */
public class ConvertSortedListToBST {
	/*
	 * solution1: create AVL tree, iterate through list, add each node to AVL tree
	 */

	/*
	 * solution2: middle value is the root of BST, left part of list is the left
	 * subtree, right part of list is the right subtree.
	 * 
	 * 2.1 find middle value in list using two pointers: fast pointer and slow
	 * pointer
	 * 
	 * 2.2 find middle value by index by converting list to array.
	 * 
	 */
	private TreeNode convertListToBST(int[] a, int left, int right) {
		// Invalid case
		if (left > right) {
			return null;
		}

		// Middle element forms the root.
		int mid = (left + right) / 2;
		TreeNode node = new TreeNode(a[mid]);

		// Base case for when there is only one element left in the array
		if (left == right) {
			return node;
		}

		// Recursively form BST on the two halves
		node.left = convertListToBST(a, left, mid - 1);
		node.right = convertListToBST(a, mid + 1, right);
		
		// subtree root at node is built return
		return node;
	}

	public TreeNode sortedListToBST(ListNode head) {

		// Form an array out of the given linked list and then
		// use the array to form the BST.
		int[] a = mapListToValues(head);

		// Convert the array to BST, return the root
		return convertListToBST(a, 0, a.length - 1);
	}

	private int[] mapListToValues(ListNode head) {
		ArrayList<Integer> values = new ArrayList<>();
		while (head != null) {
			values.add(head.val);
			head = head.next;
		}
		return values.stream().mapToInt(i -> i).toArray();
	}
}

class TreeNode {
	int value;
	TreeNode left;
	TreeNode right;

	public TreeNode(int value) {
		this.value = value;
	}
}

class ListNode {
	int val;
	ListNode next;
}
