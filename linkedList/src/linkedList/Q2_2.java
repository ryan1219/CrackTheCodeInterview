package linkedList;

/*
 * 2.2 Implement an algorithm to find the kth to last element of a singly linked list.

 */
public class Q2_2 {

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
		//System.out.println(a.get(8));
		System.out.println(a);
		Node b = kthToLast(a, 3);
		System.out.println(b);
	}

	public static Node kthToLast(linkedList a, int k){
		return null;
	}
}
