package treesAndGraphs;

public class BalanceableBinarySearchTree extends BinarySearchTree {

	private void relink(Node parent, Node child, boolean makeLeftChild) {
		child.parent = parent;
		if (makeLeftChild) {
			parent.left = child;
		} else {
			parent.right = child;
		}
	}

	/*
	 * Rotates Node x above its parents
	 */
	public void rotate(Node x) {
		// x has no parent, nothing to rotate
		if (x.parent == null) {
			return;
		}
		Node y = x.parent;
		Node z = y.parent;

		if (z == null) {
			this.root = x;
			x.parent = null;
		} else {
			// replace x with y's position
			relink(z, x, y == z.left);
		}
		if (x == y.left) {
			relink(y, x.right, true);
			relink(x, y, false);
		} else {
			relink(y, x.left, false);
			relink(x, y, true);
		}
	}

	/*
	 * four types of rotation: left left, right right, left right, right left z z /
	 * / y y / \ x x
	 */
	public void restructure(Node x) {
		Node y = x.parent;
		Node z = y.parent;

		if ((x == y.left && y == z.left) || (x == y.right && y == z.right)) {
			rotate(y);
		} else {
			rotate(x);
			rotate(x);
		}
	}

}