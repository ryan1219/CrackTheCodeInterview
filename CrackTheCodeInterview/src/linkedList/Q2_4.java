package linkedList;

/*
 * Write code to partition a linked list around a value x, such that all nodes less than
 * x come before all nodes greater than or equal to x.
 */
public class Q2_4 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
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
		// linkedList b = a.clone();
		a.head.data = 9;
		// b.head.data = 9;
		// System.out.println("b is " +b);
		System.out.println("a is " + a);
		System.out.println("partition a with 4 " + partition(a, 4));
	}

	public static linkedList partition(linkedList input, int value) {
		linkedList p1 = new linkedList();
		linkedList p2 = new linkedList();
		Node temp = input.head;
		while (temp != null) {
			if ((int) temp.data >= value) {
				p2.add(temp.data);
			} else {
				p1.add(temp.data);
			}
			temp = temp.next;
		}

		return merge(p1, p2);
	}

	public static linkedList merge(linkedList a, linkedList b) {

		// System.out.println("temp data is" + temp.data);
		if (b.head == null) {
			return a;
		}

		if (a.head == null) {
			a = b.clone();
			return a;
		}
		Node temp = a.head;
		// System.out.println("a.head is "+ a.head.data);
		// System.out.println("temp data is" + temp.data);
		while (temp.next != null) {
			temp = temp.next;
		}
		temp.next = b.head;
		return a;
	}

}
