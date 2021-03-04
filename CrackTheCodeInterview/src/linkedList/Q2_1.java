package linkedList;

public class Q2_1 {

	public static void main(String[] args) {
		linkedList a = new linkedList();
		a.add(8);
		a.add(8);
		a.add(8);
		a.add(8);
		a.add(8);
		a.add(8);
		a.add(8);
		a.add(8);
		// System.out.println(a.get(8));
		System.out.println(a);
		System.out.println(a.size());
		a = removeDuplicates(a);
		System.out.println(a);
		System.out.println(a.size());
	}

	public static linkedList removeDuplicates(linkedList a) {
		for (int i = 0; i < a.size(); i++) {
			for (int j = i + 1; j < a.size(); j++) {
				if (a.get(i).equals(a.get(j))) {
					a.remove(j);
					j = i;
				}
			}
		}
		return a;
	}

}
