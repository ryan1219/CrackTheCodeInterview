package treesAndGraphs;

public class BinarySearchTree {
	public static void main(String[] args){
		BinarySearchTree bst = new BinarySearchTree();
		bst.add(10);
		bst.add(2);
		bst.add(4);
		bst.add(3);
		bst.add(5);
		bst.add(15);
		bst.add(1);
		bst.root.print();
	}
	
	Node root;
	
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
	
	public void add(int value){
		root = addRecursive(this.root, value);
	}
	
}
