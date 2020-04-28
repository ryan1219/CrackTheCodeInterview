package practiceQuestions;

import java.util.PriorityQueue;

/*
 * question: https://leetcode.com/problems/merge-k-sorted-lists/
 */
public class MergeKSortedLists {

	/*
	 * solution1: loop through the head of each list, find the min and remove from
	 * the list and add to the new list
	 * 
	 * if has n list, each list has m nodes, n * m nodes in total, find min takes n,
	 * total time n * (nm) = m * n^2
	 */

	/*
	 * solution2: using priority queue, add all nodes to priority queue, and dequeue
	 * O(nm(log(nm))
	 * 
	 */
	public ListNode mergeKLists(ListNode[] lists) {
		PriorityQueue<Integer> queue = new PriorityQueue<>();
		for (int i = 0; i < lists.length; i++) {
			ListNode head = lists[i];
			while (head != null) {
				queue.add(head.val);
				head = head.next;
			}
		}

		ListNode head = null;
		ListNode tail = null;
		while (!queue.isEmpty()) {
			int value = queue.poll();

			if (head == null) {
				head = new ListNode(value);
				tail = head;
			} else {
				tail.next = new ListNode(value);
				tail = tail.next;
			}
		}
		
		return head;
	}

	class ListNode {
		int val;
		ListNode next;

		ListNode(int x) {
			this.val = x;
		}
	}
}
