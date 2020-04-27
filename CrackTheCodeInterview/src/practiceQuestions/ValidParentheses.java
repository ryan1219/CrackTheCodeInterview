package practiceQuestions;

import java.util.Stack;

/*
 * question: https://leetcode.com/problems/valid-parentheses/
 * 
 * matching parentheses
 */
public class ValidParentheses {
	public boolean isValid(String s) {
		Stack<Character> stack = new Stack<Character>();
		for (int i = 0; i < s.length(); i++) {
			if (!stack.isEmpty() && isPair(stack.peek(), s.charAt(i))) {
				stack.pop();
			} else {
				stack.push(s.charAt(i));
			}
		}
		return stack.isEmpty();
	}

	public boolean isPair(char s, char t) {
		if ((s == '(' && t == ')') || (s == '[' && t == ']') || (s == '{' && t == '}')) {
			return true;
		}
		return false;
	}
}
