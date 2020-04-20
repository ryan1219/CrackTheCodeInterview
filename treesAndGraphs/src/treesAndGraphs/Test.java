package treesAndGraphs;

public class Test {
	public static void main(String[] args){
		BinarySearchTree bst = new BinarySearchTree();
		bst.add(10);
		bst.add(2);
		bst.add(4);
		bst.add(3);
		bst.add(5);
		bst.add(15);
		bst.add(1);
		bst.print();
		System.out.println(bst.find(15));
		bst.reverse();
		bst.print();
	}
}
