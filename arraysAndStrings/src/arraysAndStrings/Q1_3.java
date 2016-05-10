package arraysAndStrings;

/*
 * solution from the Internet
 * print out the list of permutation strings of an input string x 
 */
public class Q1_3 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		String a = "dog";
		printAllPermutation(a);

	}
	
	public static void printAllPermutation(String input){
		int inputLength = input.length();
		boolean[ ] used = new boolean[ inputLength ];
		StringBuffer outputString = new StringBuffer();
		char[ ] in = input.toCharArray( );
		  
		doPermute ( in, outputString, used, inputLength, 0 );
	}
	public static void doPermute( char[ ] in, StringBuffer outputString, 
            boolean[ ] used, int inputLength, int level){
		if( level == inputLength) {
			System.out.println ( outputString.toString()); 
			return;
		}

		for( int i = 0; i < inputLength; ++i )
		{       

			if( used[i] ) continue;

			outputString.append( in[i] );      
			used[i] = true;       
			doPermute( in,   outputString, used, inputLength, level + 1 );       
			used[i] = false;       
			outputString.setLength(   outputString.length() - 1 );   
		}
	}

}
