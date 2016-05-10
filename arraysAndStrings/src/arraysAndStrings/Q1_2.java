package arraysAndStrings;
/* 
 * Written by Yan
 * reverse a string in java
 */
public class Q1_2 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String a = "abcdefg1234567890";
		System.out.println("original string "+ a);
		System.out.println("reverse string "+ reverse(a));
	}

	public static String reverse(String str){
		char[] temp = new char[str.length()];
		int index = 0;
		for(int i = str.length()-1; i>=0; i--){
			temp[index]=str.charAt(i);
			index++;
		}
		return new String(temp);
	}
}
