package linkedList;

/*
 * Given a circular linked list, implement an algorithm which returns the node at
 *	the beginning of the loop
 * Similar to comp250 assignment
 * 
 * base on Floyd’s loop detection algorithm. Below are steps to find the first node of the loop.
 * 1. If a loop is found, initialize a slow pointer to head, let fast pointer be at its position. 
 * 2. Move both slow and fast pointers one node at a time. 
 * 3. The point at which they meet is the start of the loop.
 */
public class Q2_6 {

	public static void main(String[] args) {
		Node head = new Node(50);
		head.next = new Node(20);
		head.next.next = new Node(15);
		head.next.next.next = new Node(4);
		head.next.next.next.next = new Node(10);

		// Create a loop for testing
		head.next.next.next.next.next = head.next.next;

		Node res = detectFirstNodeInLoop(head);
		// expect node 15
		System.out.println(res.data);
	}

	public static Node detectFirstNodeInLoop(Node head) {
		if (head == null || head.next == null)
			return null;

		Node slow = head, fast = head;

		// Move slow and fast 1
		// and 2 steps ahead
		// respectively.
		slow = slow.next;
		fast = fast.next.next;

		// Search for loop using
		// slow and fast pointers
		while (fast != null && fast.next != null) {
			if (slow == fast)
				break;
			slow = slow.next;
			fast = fast.next.next;
		}

		// If loop does not exist
		if (slow != fast)
			return null;

		// If loop exists. Start slow from
		// head and fast from meeting point.
		slow = head;
		while (slow != fast) {
			slow = slow.next;
			fast = fast.next;
		}

		return slow;
	}

}
/*
 * how to detect loop in a linked list
 * 
 * 1. hash: traverse the list one by one and keep putting the node in a
 * hashtable, if null is reach then no loop, ifnext of current node points to
 * any of the previously stored nodes return true
 * 
 * 2. looping and record the visited node
 * 
 * 3. Floyd’s Cycle-Finding Algorithm: traverse linked list by two pointers,
 * slow pointer and fast pointer slow pointer moves one node at a time, fast
 * pointer moves two nodes at a time, if two pointers meet at the same node then
 * there is a loop
 * 
 */