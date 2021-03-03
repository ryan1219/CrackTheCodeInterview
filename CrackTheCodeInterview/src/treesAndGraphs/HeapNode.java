package treesAndGraphs;

public class HeapNode {
	int value;
	HeapNode left;
	HeapNode right;
	HeapNode parent;

	HeapNode(int value) {
		this.value = value;
		this.left = null;
		this.right = null;
		this.parent = null;
	}

}
