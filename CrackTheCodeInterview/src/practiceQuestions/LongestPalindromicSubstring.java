package practiceQuestions;

public class LongestPalindromicSubstring {

	/*
	 * brute force
	 */

	/*
	 * DP
	 */

	public static void main(String[] args) {
		String s = "aaaa";
		new LongestPalindromicSubstring().longestPalindrome(s);
	}

	public String longestPalindrome(String s) {
		if (s.isEmpty()) {
			return s;
		}
		int max = 0;
		int max_i = 0;
		int max_j = 0;

		int[][] dp = new int[s.length()][s.length()];

		for (int i = 0; i < s.length(); i++) {
			dp[i][i] = 1;
		}

		for (int k = 1; k < s.length(); k++) {
			for (int i = 0; i + k < s.length(); i++) {
				if (s.charAt(i) == s.charAt(k + i)) {
					if (k == 1 || dp[i + 1][k + i - 1] == 1) {
						dp[i][k + i] = 1;
						if (k > max) {
							max = k;
							max_i = i;
							max_j = k + i;
						}
					}

				} else {
					dp[i][k + i] = 0;
				}

			}
		}
		if (max == 0) {
			return s.substring(0, 1);
		}

		return s.substring(max_i, max_j + 1);
	}
}
