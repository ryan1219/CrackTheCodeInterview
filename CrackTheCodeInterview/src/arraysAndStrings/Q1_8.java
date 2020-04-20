package arraysAndStrings;
/*
 * Check string rotation 
 */
public class Q1_8 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String s1 = "waterbottle";
		String s2 = "aterbottlew";
		System.out.println(isRotation(s1,s2));
	}
	
	public static boolean isRotation(String a, String b){
		if(a.length()!=b.length()||a.length()==0||b.length()==0){
			return false;
		}
		for(int i = 0; i <a.length(); i++){
			String temp = a.substring(i) + a.substring(0, i);
			if(temp.equals(b)){
				return true;
			}
		}
		return false;
	}

}
