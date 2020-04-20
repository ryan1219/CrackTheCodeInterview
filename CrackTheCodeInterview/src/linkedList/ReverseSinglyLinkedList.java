package linkedList;

public class ReverseSinglyLinkedList {

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
		System.out.println(a);
		System.out.println(a.head.data);
		recursiveReverse(a);
		System.out.println(a);
		System.out.println(a.head.data);
	}
	
	public static void iterativeReverse(linkedList a){
		Node current = a.head;
		Node previous = null;
		Node next = null;
		while(current != null){
			next = current.next;
			current.next = previous;
			previous = current;
			current = next;
		}
		a.head = previous;
	}
	
	public static void recursiveReverse(linkedList a){
		reverse(a, a.head, null);
		
	}
	
	public static void reverse(linkedList a, Node current, Node previous){
		// base case
		if(current.next == null){
			a.head = current;
			current.next = previous;
			return;
		}
		// recursive case
		Node next = current.next;
		current.next = previous;
		reverse(a, next, current);
	}

}
