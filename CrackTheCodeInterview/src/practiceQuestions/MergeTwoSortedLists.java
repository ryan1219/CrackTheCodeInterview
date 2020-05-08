package practiceQuestions;

/*
 * question: https://leetcode.com/problems/merge-two-sorted-lists/
 */
public class MergeTwoSortedLists {
	class ListNode {
		int val;
		ListNode next;

		ListNode() {
		}

		ListNode(int val) {
			this.val = val;
		}

		ListNode(int val, ListNode next) {
			this.val = val;
			this.next = next;
		}
	}

	public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
		// edge case
		if (l1 == null) {
			return l2;
		}
		if (l2 == null) {
			return l1;
		}
		ListNode p1 = l1;
		ListNode p2 = l2;
		ListNode head = null;
		// initial head
		if (p1.val < p2.val) {
			head = p1;
			p1 = p1.next;
			head.next = null;
		} else {
			head = p2;
			p2 = p2.next;
			head.next = null;
		}
		ListNode current = head;

		while (p1 != null && p2 != null) {
			if (p1.val < p2.val) {
				current.next = p1;
				p1 = p1.next;
				current = current.next;
				current.next = null;
			} else {
				current.next = p2;
				p2 = p2.next;
				current = current.next;
				current.next = null;
			}
		}

		while (p1 != null) {
			current.next = p1;
			p1 = p1.next;
			current = current.next;
			current.next = null;
		}

		while (p2 != null) {
			current.next = p2;
			p2 = p2.next;
			current = current.next;
			current.next = null;
		}

		return head;
	}
}
