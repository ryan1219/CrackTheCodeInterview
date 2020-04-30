package practiceQuestions;

import java.util.HashMap;

public class LRUCache {

	HashMap<Integer, Node> map = new HashMap<>();

	Node head = new Node();
	Node tail = new Node();

	int size = 0;
	int capacity;

	public LRUCache(int capacity) {
		this.capacity = capacity;
		head.next = tail;
		tail.prev = head;
	}

	public int get(int key) {
		// node exists
		if (map.containsKey(key)) {
			Node temp = map.get(key);
			remove(temp);
			addHead(temp);
			return temp.value;
		} else {
			return -1;
		}
	}

	public void put(int key, int value) {
		// old node, update priority and value
		if (map.containsKey(key)) {
			Node temp = map.get(key);
			temp.value = value;
			remove(temp);
			addHead(temp);
			return;
		}

		// new node
		Node temp = new Node();
		temp.value = value;
		temp.key = key;
		addHead(temp);
		if (size < capacity) {
			size++;
		} else {
			map.remove(tail.prev.key);
			remove(tail.prev);
		}
		map.put(key, temp);
	}

	void addHead(Node node) {
		node.prev = head;
		node.next = head.next;
		head.next.prev = node;
		head.next = node;
	}

	void remove(Node node) {
		node.next.prev = node.prev;
		node.prev.next = node.next;
		node.prev = null;
		node.next = null;
	}

	class Node {
		int value, key;
		Node prev, next;
	}
}
