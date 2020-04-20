package stackAndQueues;

public class Stack {
	Node top;
	
	public Object pop(){
		if(top != null){
			Object item = top.data;
			top = top.next;
			return item;
		}
		System.out.println("empty stack");
		return null;
	}
	
	public void push(Object item){
		Node t = new Node(item);
		t.next = top;
		top = t;
	}
	
	Object peek(){
		return top.data;
	}
}
