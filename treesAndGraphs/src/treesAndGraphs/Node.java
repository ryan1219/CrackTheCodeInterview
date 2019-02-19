package treesAndGraphs;

public class Node {
	int value;
	Node left;
	Node right;

	Node(int value) {
		this.value = value;
		right = null;
		left = null;
	}

	public void print() {
		print("", this, false);
	}

	public void print(String prefix, Node n, boolean isLeft) {
		if (n != null) {
			System.out.println(prefix + (isLeft ? "|-- " : "\\-- ") + n.value);
			print(prefix + (isLeft ? "|   " : "    "), n.left, true);
			print(prefix + (isLeft ? "|   " : "    "), n.right, false);
		}
	}
}
