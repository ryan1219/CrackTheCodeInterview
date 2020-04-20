package linkedList;

public class Q2_5 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		linkedList a = new linkedList();
		linkedList b = new linkedList ();
		//a.add(9);
		a.add(9);
		a.add(9);
		//b.add(5);
		b.add(9);
		b.add(9);
		System.out.println("sum is "+ add(a,b));

	}
	public static linkedList add (linkedList a, linkedList b){
		linkedList sum = new linkedList();
		Node aTemp = a.head;
		Node bTemp = b.head;
		boolean carry = false;
		
		while(aTemp != null && bTemp != null){
			int tempSum = (int)aTemp.data + (int)bTemp.data;
			
			if(carry == true){
				tempSum++;
			}
			if(tempSum >= 10){
				carry = true;
				sum.add(tempSum - 10);
			}else{
				carry = false;
				sum.add(tempSum);
			}
			aTemp = aTemp.next;
			bTemp = bTemp.next;
		}
		if(aTemp == null && bTemp != null){
			while(bTemp != null){
				int tempSum = (int)bTemp.data;
				if(carry == true){
					tempSum++;
				}
				if(tempSum >= 10){
					carry = true;
					sum.add(tempSum - 10);
				}else{
					carry = false;
					sum.add(tempSum);
				}
				bTemp = bTemp.next;
			}
		}else if(bTemp == null && aTemp != null){
			while(aTemp != null){
				int tempSum = (int)aTemp.data;
				if(carry == true){
					tempSum++;
				}
				if(tempSum >= 10){
					carry = true;
					sum.add(tempSum - 10);
				}else{
					carry = false;
					sum.add(tempSum);
				}
				aTemp = aTemp.next;
			}
		}
		if(carry == true){
			sum.add(1);
		}
		return sum;		
	}

}
