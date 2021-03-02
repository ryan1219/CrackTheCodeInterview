package linkedList;

/*
 * Implement a function to check if a linked list is a palindrome.
 * Implement in C, http://www.geeksforgeeks.org/function-to-check-if-a-singly-linked-list-is-palindrome/
 */
public class Q2_7 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		linkedList test = new linkedList();
		test.add('R');
		test.add('A');
		test.add('D');
		test.add('A');
		test.add('R');
		System.out.println("test linkedList is " + test);
		System.out.println("test is palindrome " + isPal(test));
	}

	public static boolean isPal(linkedList input) {
		Wrapper wrapper = new Wrapper();
		wrapper.left = input.head;
		return recursion(wrapper, input.head);
	}

	public static void test(Node left) {
		System.out.println("inside test method " + left.data);
	}

	public static boolean recursion(Wrapper wrapper, Node right) {

		if (right == null) {
			return true;
		}

		boolean isp = recursion(wrapper, right.next);
		if (isp == false) {
			return false;
		}

		boolean isp1 = (wrapper.left.data == right.data);

		wrapper.left = wrapper.left.next;
		return isp1;
	}

}

class Wrapper {
	Node left;
}