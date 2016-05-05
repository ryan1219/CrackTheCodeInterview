package arraysAndStrings;

public class Q1_1 {
	//sort the array in ASCII order
	public static boolean isUniqueChar(String str){
		//sort
		
		//
		for(int i =0; i <str.length()-1; i++){
			if(str.charAt(i)==str.charAt(i+1)) return false;
		}
		return true;
	}
	//Build an ASCII table
	public static boolean isUniqueChars2(String str) {
		if (str.length() > 128) {
			return false;
		}
		boolean[] char_set = new boolean[128];
		for (int i = 0; i < str.length(); i++) {
			int val = str.charAt(i);
			if (char_set[val]) return false;
			char_set[val] = true;
		}
		return true;
	}
	
	public static void main(String[] args) {
		String[] words = {"abcde", "helblo", "apple", "kite", "padle"};
		for (String word : words) {
			System.out.println(word + ": " + isUniqueChars2(word));
		}
		//System.out.println((1<<-3));
	}

}
