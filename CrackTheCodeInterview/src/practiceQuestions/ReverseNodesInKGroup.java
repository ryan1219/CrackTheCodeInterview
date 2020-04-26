package practiceQuestions;

/*
 * question: https://leetcode.com/problems/reverse-nodes-in-k-group/
 */
public class ReverseNodesInKGroup {

	/*
	 * solution1: single list manipulation
	 * 
	 */

	public ListNode reverseKGroup(ListNode head, int k) {
		if (head == null || head.next == null || k <= 1) {
			return head;
		}
		ListNode start = new ListNode(-1);
		start.next = head;
		ListNode end = head;
		boolean firstReverse = true;

		int i = 0;
		while (end != null) {
			i++;
			if (i % k == 0) {
				ListNode previous = reverse(start, end.next);
				if (firstReverse) {
					firstReverse = false;
					head = start.next;
				}
				start = previous;
				end = previous.next;
			} else {
				end = end.next;
			}
		}

		return head;
	}
	
	/*
	 * reverse part of single linked list
	 * 
	 *    head            end
	 * 1 -> 2 -> 3 -> 4 -> 5
	 * 
	 *    head            end
	 * 1 -> 2 -> 4 -> 3 -> 5
	 * 
	 */

	public ListNode reverse(ListNode head, ListNode end) {
		ListNode previous = head;
		ListNode current = head.next;
		ListNode next = null;
		while (current != end) {
			next = current.next;
			current.next = previous;
			previous = current;
			current = next;
		}
		
		head.next.next = end;
		ListNode nodeBeforeEnd = head.next;
		head.next = previous;
		return nodeBeforeEnd;
	}

	class ListNode {
		int val;
		ListNode next;

		ListNode(int x) {
			this.val = x;
		}
	}
}
