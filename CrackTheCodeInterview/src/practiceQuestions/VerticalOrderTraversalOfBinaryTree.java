package practiceQuestions;

import java.util.ArrayList;
import java.util.List;
import java.util.Map.Entry;
import java.util.PriorityQueue;
import java.util.TreeMap;

/*
 * question: https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
 * 
 */
public class VerticalOrderTraversalOfBinaryTree {

	/*
	 * solution1: treeMap
	 */

	TreeMap<Integer, TreeMap<Integer, PriorityQueue<Integer>>> order = new TreeMap<>();

	public List<List<Integer>> verticalTraversal(TreeNode root) {

		int x = 0;
		int y = 0;

		inorder(root, x, y);

		List<List<Integer>> result = new ArrayList<>();
		for (Entry<Integer, TreeMap<Integer, PriorityQueue<Integer>>> ex : order.entrySet()) {
			TreeMap<Integer, PriorityQueue<Integer>> ey = ex.getValue();
			ArrayList<Integer> yValues = new ArrayList<>();
			for (Entry<Integer, PriorityQueue<Integer>> e : ey.entrySet()) {
				yValues.addAll(queueToList(e.getValue()));
			}
			result.add(yValues);
		}

		return result;
	}

	public void inorder(TreeNode root, int x, int y) {
		if (root == null) {
			return;
		}
		inorder(root.left, x - 1, y + 1);

		if (!order.containsKey(x)) {
			order.put(x, new TreeMap<>());
		}
		if (!order.get(x).containsKey(y)) {
			order.get(x).put(y, new PriorityQueue<>());
		}

		order.get(x).get(y).add(root.val);
		inorder(root.right, x + 1, y + 1);
	}

	public ArrayList<Integer> queueToList(PriorityQueue<Integer> queue) {
		ArrayList<Integer> result = new ArrayList<>();
		while (!queue.isEmpty()) {
			result.add(queue.poll());
		}
		return result;
	}

	/*
	 * solution2: traverse the tree and build a list of Position, sort the list by
	 * using implemented Comparable in Position
	 */
	class Position implements Comparable<Position> {
		int x;
		int y;
		int val;

		Position(int x, int y, int val) {
			this.x = x;
			this.y = y;
			this.val = val;
		}

		public int compareTo(Position that) {
			if (this.x != that.x)
				return Integer.compare(this.x, that.x);
			else if (this.y != that.y)
				return Integer.compare(this.y, that.y);
			else
				return Integer.compare(this.val, that.val);
		}
	}

	class TreeNode {
		int val;
		TreeNode left;
		TreeNode right;

		TreeNode(int x) {
			val = x;
		}
	}
}
