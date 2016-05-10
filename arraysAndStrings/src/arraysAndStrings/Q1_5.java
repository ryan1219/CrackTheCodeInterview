package arraysAndStrings;
/*
 * string compression
 */
public class Q1_5 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String originalString = "aabcccccaaa";
		System.out.println(compress(originalString));
	}
	
	public static String compress(String str){
		char[] input = str.toCharArray();
		StringBuffer temp = new StringBuffer();
		temp.append(input[0]);
		
		int count = 1;
		for(int i = 1 ; i < input.length; i++){
			if (input[i]==input[i-1])
			{
				count++;
			}
			else
			{
				temp.append(count);
				temp.append(input[i]);
				count = 1;
			}
		}
		if(count > 1){
			temp.append(count);
		}
		
		String output = temp.toString();
		if(output.length() < str.length()){
			return output;
		}else{
			return str;
		}
	}

}
