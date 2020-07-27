package practiceQuestions;

import java.util.LinkedList;

/*
 * https://leetcode.com/problems/permutation-sequence/
 */
public class PermutationSequence {
	public String getPermutation(int n, int k) {
		String res = "";
		LinkedList<Integer> order = new LinkedList<>();
		for (int i = 1; i <= n; i++) {
			order.add(i);
		}
		k -= 1;
		while(n > 0) {
			int divisor = factorial(n - 1);
			int quotient = k / divisor;
			k = k % divisor;
			n -= 1;
			res += order.get(quotient);
			order.remove(quotient);
		}

		return res;
	}
	
	public int factorial(int n) {
		int res = 1;
		while(n > 0) {
			res = res  * n;
			n -= 1;
		}
		return res;
	}
}
