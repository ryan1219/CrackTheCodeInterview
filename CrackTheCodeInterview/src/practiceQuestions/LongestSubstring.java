package practiceQuestions;
import java.util.HashMap;

/*
 * Question: give a string, find the length of the longest substring without repeating characters
 */
public class LongestSubstring {

	public static void main(String[] args) {
		System.out.println(lengthOfLongestSubstring("abcabcbb"));
		System.out.println(lengthOfLongestSubstring("bbbbbbb"));
		System.out.println(lengthOfLongestSubstring("pwwkew"));
	}
	
	public static int lengthOfLongestSubstring(String s) {
		int left = 0, right = 0, maxLength = 0;
		HashMap<Character, Integer> seen = new HashMap<>();
		while(right < s.length()) {
			char rightChar = s.charAt(right);
			if(seen.containsKey(rightChar)) {
				seen.put(s.charAt(right), seen.get(rightChar)+ 1);
			}else {
				seen.put(s.charAt(right), 1);
			}
			right += 1;
			while(seen.get(rightChar) > 1) {
				char leftChar = s.charAt(left);
				if(leftChar == rightChar) {
					seen.put(rightChar, seen.get(rightChar) - 1);
				}else {
					seen.remove(leftChar);
				}
				left += 1;
			}
			
			maxLength = Math.max(right - left, maxLength);
		}
		return maxLength;
	}

}
