package leecode;

import java.util.Arrays;
import java.util.LinkedList;

/*
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
 */
public class AddTwoNumbers {
	
	public static void main(String[] args){
		LinkedList<Integer> a = new LinkedList<>(Arrays.asList(2));
		LinkedList<Integer> b = new LinkedList<>(Arrays.asList(5,6,4));
		System.out.println(addTwoNumbers(a,b));
	}
	
	public static LinkedList<Integer> addTwoNumbers(LinkedList<Integer> a, LinkedList<Integer> b){
		LinkedList<Integer> longNumber;
		LinkedList<Integer> shortNumber;
		LinkedList<Integer>	result = new LinkedList<>();
		if(a.size() > b.size()){
			longNumber = a;
			shortNumber = b;
		}else{
			longNumber = b;
			shortNumber = a;
		}
		int increment = 0;
		int i;
		for(i = 0; i<shortNumber.size(); i++){
			int sum = increment+shortNumber.get(i)+longNumber.get(i);
			result.add(sum%10);
			increment = sum/10;
		}
		while(i<longNumber.size()){
			int sum = increment+longNumber.get(i);
			result.add(sum%10);
			increment = sum/10;
			i++;
		}
		if(increment != 0){
			result.add(1);
		}
		
		return result;
	}

}
