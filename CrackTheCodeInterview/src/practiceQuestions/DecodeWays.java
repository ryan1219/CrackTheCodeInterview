package practiceQuestions;

/*
 * question: https://leetcode.com/problems/decode-ways/
 */
public class DecodeWays {

	/*
	 * DP
	 * 
	 * "220"
	 * 
	 * if dp[i-1]dp[i] <= 26 dp[i] = dp[i - 2] + dp[i - 1] else dp[i] = dp[i-1]
	 * 
	 * 
	 */
	public int numDecodings(String s) {
		if (s.length() == 0) {
			return 0;
		}
		int[] values = s.chars().map(x -> x - '0').toArray();
		int[] dp = new int[s.length()];
		// base case
		dp[0] = s.charAt(0) == '0' ? 0 : 1;
		if (s.length() == 1) {
			return dp[0];
		}
		if (values[1] >= 1 && values[1] <= 9) {
			dp[1] = dp[0];
		}
		if (values[0] * 10 + values[1] >= 10 && values[0] * 10 + values[1] <= 26) {
			dp[1] += 1;
		}
		for (int i = 2; i < s.length(); i++) {
			if (1 <= values[i] && values[i] <= 9) {
				dp[i] += dp[i - 1];
			}
			if (10 <= (values[i - 1] * 10 + values[i]) && (values[i - 1] * 10 + values[i]) <= 26) {
				dp[i] += dp[i - 2];
			}
		}
		return dp[dp.length - 1];
	}

	public static void main(String[] args) {
		new DecodeWays().numDecodings("100");
	}
}
