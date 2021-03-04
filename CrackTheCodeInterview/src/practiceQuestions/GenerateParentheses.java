package practiceQuestions;

import java.util.ArrayList;
import java.util.List;

/*
 * question: https://leetcode.com/problems/generate-parentheses/
 * 
 * explanation:
 * https://algorithms.tutorialhorizon.com/generate-all-valid-parenthesis-strings-of-length-2n-of-given-n/
 * https://www.youtube.com/watch?v=sz1qaKt0KGQ
 */
public class GenerateParentheses {

	List<String> res;

	public List<String> generateParenthesis(int n) {
		res = new ArrayList<>();
		addParenthesis(n, n, "");

		return res;
	}

	void addParenthesis(int open, int close, String string) {
		if (open == 0 && close == 0) {
			res.add(string);
		}

		if (close < open) {
			return;
		}

		if (open > 0) {
			addParenthesis(open - 1, close, string + "(");
		}

		if (close > 0) {
			addParenthesis(open, close - 1, string + ")");
		}
	}
}
