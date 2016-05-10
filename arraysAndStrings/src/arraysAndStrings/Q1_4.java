package arraysAndStrings;
/*
 * Written by Yan
 */
public class Q1_4 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String test = "ms jhon smith  ";
		System.out.println("test string: "+ test + "output string: "+ replace(test));
	}
	
	public static String replace(String str){
		StringBuffer outputString = new StringBuffer();
		char[] in = str.toCharArray();
		for(int i = 0; i < in.length; i++){
			// 
			if(in[i] == ' '&& i+1 < in.length && in[i+1] == ' '){
				break;
			}
			if(in[i] == ' '&& i+1 == in.length){
				break;
			}
			if(in[i] == ' '){
				outputString.append("%20");
			}
			else{
				outputString.append(in[i]);
			}
		}
		return outputString.toString();
	}

}
