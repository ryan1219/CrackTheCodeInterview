package practiceQuestions;

import java.util.HashMap;

public class CopyListWithRandomPointer {

	public Node copyRandomList(Node head) {
		if (head == null) {
			return null;
		}

		HashMap<Node, Node> map = new HashMap<>();
		Node newHead = new Node(head.val);
		map.put(head, newHead);
		Node p1 = newHead;
		Node p2 = head.next;
		while (p2 != null) {
			Node temp = new Node(p2.val);
			map.put(p2, temp);
			p1.next = temp;
			p2 = p2.next;
			p1 = p1.next;
		}

		p2 = head;
		p1 = newHead;
		while (p2 != null) {
			if (p2.random == null) {
				p1.random = null;
			} else {
				p1.random = map.get(p2.random);
			}
			p2 = p2.next;
			p1 = p1.next;
		}

		return newHead;
	}

	class Node {
		int val;
		Node next;
		Node random;

		public Node(int val) {
			this.val = val;
			this.next = null;
			this.random = null;
		}
	}
}
