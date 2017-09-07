package stackAndQueues;

public class Queue {
	Node head, tail;

	void enqueue(Object item) {
		if (head == null) {
			tail = new Node(item);
			head = tail;
		} else {
			head.next = new Node(item);
			head = head.next;
		}
	}

	Object dequeue() {
		if (tail != null) {
			Object item = tail.data;
			tail = tail.next;
			return item;
		}
		return null;
	}
}
