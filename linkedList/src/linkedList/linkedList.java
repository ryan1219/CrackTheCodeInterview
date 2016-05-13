package linkedList;

public class linkedList {
	
	private int size;
	public Node head;
	//Default constructor
	public linkedList(){
		this.size = 0;
		this.head = null;
	}
	//
	public void add(Object data){
		
		if (head==null){
			head = new Node (data);
		}
		
		Node newNode = new Node(data);
		Node temp = this.head;
		
		if(temp != null){
			while(temp.next != null){
				temp = temp.next;
			}
			temp.next = newNode;
		}
		size++;
	}
	//
	public int size(){
		return size;
	}
	
}
