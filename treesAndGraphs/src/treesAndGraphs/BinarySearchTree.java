package treesAndGraphs;

public class BinarySearchTree {

	private Node root = null;

	private Node addRecursive(Node current, int value) {
		if (current == null) {
			return new Node(value);
		}

		if (value < current.value) {
			current.left = addRecursive(current.left, value);
		} else if (value > current.value) {
			current.right = addRecursive(current.right, value);
		} else {
			// value already exists
			return current;
		}

		return current;
	}

	/*
	 * return the node contains the value, return null if not exists
	 */
	private Node findRecursive(Node current, int value) {
		if (current == null) {
			return current;
		}
		if (current.value == value) {
			return current;
		}
		if (current.value > value) {
			return findRecursive(current.left, value);
		}
		return findRecursive(current.right, value);
	}

	private void reverseRecursive(Node current) {
		if(current == null) {
			return;
		}
		
		Node temp = current.left;
		current.left = current.right;
		current.right = temp;
		
		reverseRecursive(current.left);
		reverseRecursive(current.right);
	}
	
	public void add(int value) {
		if (root == null) {
			root = new Node(value);
		} else {
			root = addRecursive(this.root, value);
		}
	}
	
	public Node find(int value) {
		return findRecursive(this.root, value);
	}
	
	public void reverse() {
		reverseRecursive(this.root);
	}
	
	public void print() {
		this.root.print();
	}
}
