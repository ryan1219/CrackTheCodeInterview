package practiceQuestions;

public class RemoveNthNodeFromEndOfList {

	public ListNode removeNthFromEnd(ListNode head, int n) {
		ListNode start = head;
		// want to remove nth item, need to stop at n+1 item to change pointer
		int itemNumbers = remove(start, n + 1) - 1;
		// edge case: removed item is the head
		if (itemNumbers == n) {
			if (head.next == null) {
				head = null;
			} else {
				head = head.next;
			}
		}
		return head;
	}

	public int remove(ListNode node, int stop) {
		// base case
		if (node == null) {
			return 1;
		}

		int current = remove(node.next, stop);
		if (current == stop) {
			node.next = node.next.next;
		}

		return current + 1;
	}

	public class ListNode {
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
}
