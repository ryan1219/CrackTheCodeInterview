package treesAndGraphs;

public class LinkedHeap implements Heap {

	private HeapNode root;
	private int size;

	private HeapNode getLastNode() {
		String binary = Integer.toBinaryString(size);
		HeapNode current = root;
		for (int i = 1; i < binary.length(); i++) {
			if (binary.charAt(i) == '0') {
				current = current.left;
			} else {
				current = current.right;
			}
		}
		return current;
	}

	private void insertLast(int value) {
		String binary = Integer.toBinaryString(size + 1);
		HeapNode current = root;
		for (int i = 1; i < binary.length() - 1; i++) {
			if (binary.charAt(i) == '0') {
				current = current.left;
			} else {
				current = current.right;
			}
		}

		if (binary.charAt(binary.length() - 1) == '0') {
			current.left = new HeapNode(value);
		} else {
			current.right = new HeapNode(value);
		}

		size++;
	}

	private void removeLast() {
		HeapNode last = getLastNode();
		if(last.parent.left == last) {
			last.parent.left = null;
		}else {
			last.parent.right = null;
		}
		last.parent = null;
	}

	/*
	 * swap the value inside two nodes, i and j
	 */
	private void swap(HeapNode i, HeapNode j) {
		int temp = i.value;
		i.value = j.value;
		j.value = temp;
	}

	private void heapifyUp(HeapNode j) {
		while (j != root) {
			HeapNode p = j.parent;
			if (j.value > p.value) {
				break;
			}
			swap(j, p);
			j = p;
		}
	}

	private void heapifyDown(HeapNode j) {
		while (j.left != null) {
			HeapNode min = j.left;
			if (j.right != null) {
				if (j.right.value < j.left.value) {
					min = j.right;
				}
			}

			if (j.value < min.value) {
				break;
			}
			swap(j, min);
			j = min;
		}
	}

	@Override
	public void insert(int i) {
		insertLast(i);
		heapifyUp(getLastNode());
	}

	@Override
	public int remove() throws Exception {
		int result = root.value;
		swap(root, getLastNode());
		removeLast();
		heapifyDown(root);
		return result;
	}

}
