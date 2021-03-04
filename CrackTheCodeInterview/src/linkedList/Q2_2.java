package linkedList;

/*
 * 2.2 Implement an algorithm to find the kth to last element of a singly linked list.

 */
public class Q2_2 {

	public static void main(String[] args) {
		linkedList a = new linkedList();
		a.add(1);
		a.add(2);
		a.add(3);
		a.add(4);
		a.add(5);
		a.add(6);
		a.add(7);
		a.add(8);
		// System.out.println(a.get(8));
		System.out.println(a);
		Node b = kthToLast(a, 3);
		System.out.println(b.data);
	}

	// iterative solution
	public static Node kthToLast(linkedList a, int k) {
		Node p1 = a.head;
		Node p2 = a.head;

		for (int i = 0; i < k - 1; i++) {
			if (p2 == null)
				return null;
			p2 = p2.next;
		}
		if (p2 == null)
			return null;
		while (p2.next != null) {
			p1 = p1.next;
			p2 = p2.next;
		}
		return p1;
	}
}
