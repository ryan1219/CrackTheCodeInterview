package linkedList;

public class linkedList{
	
	private int size;
	public Node head;
	//Default constructor
	public linkedList(){
		this.size = 0;
		this.head = null;
	}
	
	public linkedList clone(){
		linkedList copy = new linkedList();
		Node temp = this.head;
		
		while(temp!= null){
			copy.add(temp.data);
			temp= temp.next;
		}
		return copy;
	}
	//
	public void add(Object data){
		
		if (head==null){
			head = new Node (data);
			size++;
			return;
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
	// removes the element sat the specified position
	public boolean remove(int index){
		//if the index is out of range, exit
		if(this.head == null){
			System.out.println("empty linked list");
			return false;
		}
		else if(index <0 || index >= size){
			System.out.println("index out of boundary");
			return false;
		}
		
		Node currentNode = head;
		//remove head
		if(index==0){
				this.head = this.head.next;
				currentNode.next = null;
				size--;
				return true;		
		}
		
		if(head != null){
			for(int i =0; i<index-1; i++){
				if(currentNode.next == null){
					System.out.println("linked list break at index" + i + 1);
					return false;
				}
				currentNode = currentNode.next;
			}
			Node temp = currentNode.next;		
			currentNode.next = 	currentNode.next.next;
			temp.next = null;
			size--;
			return true;
		}
		return false;
	}
	//
	public Object get(int index){
		if(index < 0 || index >= size){
			System.out.println("index out of boundary");
			return null;
		}
		Node currentNode = head;
		if(currentNode != null){
			for(int i = 0 ; i < index; i++){
				currentNode = currentNode.next;
			}
			return currentNode.data;
		}
		System.out.println("empty linked list");
		return null;
	}
	//
	public String toString(){
		String output = "";
		if (head != null){
			Node currentNode = head;
			while(currentNode != null){
				output += "[" +currentNode.data+"]"+"->";
				currentNode = currentNode.next;
			}
		}
		return output;
	}
}
