package linkedList;

public class Node {
	public Node next;
	public Object data;
	
	public Node(){
		next= null;
		this.data = null;
	}
	public Node(Object data){
		next = null;
		this.data = data;
	}
	public Node(Object data, Node next){
		this.next = next;
		this.data = data;
	}
	
}
