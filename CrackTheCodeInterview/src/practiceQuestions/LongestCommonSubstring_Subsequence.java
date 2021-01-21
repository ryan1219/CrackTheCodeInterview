package practiceQuestions;

public class LongestCommonSubstring_Subsequence {

	public static int LCSubStr(char x[], char y[], int m, int n) {
		// Create a table to store lengths of longest common suffixes of substrings.
		// Note that dp[i][j] contains length of longest common suffix of x[0..i-1]
		// and y[0..j-1].
		// The first row and first column entries have no logical meaning, they are used
		// only for simplicity of program
		int dp[][] = new int[m + 1][n + 1];

		// To store length of the longest common substring
		int result = 0;

		// Following steps build
		// dp[m+1][n+1] in bottom up fashion
		for (int i = 0; i <= m; i++) {
			for (int j = 0; j <= n; j++) {
				if (i == 0 || j == 0)
					dp[i][j] = 0;
				else if (x[i - 1] == y[j - 1]) {
					dp[i][j] = dp[i - 1][j - 1] + 1;
					result = Integer.max(result, dp[i][j]);
				} else
					dp[i][j] = 0;
			}
		}
		return result;
	}

	public static int LCSubSeq(char x[], char y[], int m, int n) {
		// Create a table to store lengths of longest common suffixes of substrings.
		// Note that dp[i][j] contains length of longest common suffix of x[0..i-1]
		// and y[0..j-1].
		// The first row and first column entries have no logical meaning, they are used
		// only for simplicity of program
		int dp[][] = new int[m + 1][n + 1];

		// To store length of the longest common substring
		int result = 0;

		// Following steps build
		// dp[m+1][n+1] in bottom up fashion
		for (int i = 0; i <= m; i++) {
			for (int j = 0; j <= n; j++) {
				if (i == 0 || j == 0)
					dp[i][j] = 0;
				else if (x[i - 1] == y[j - 1]) {
					dp[i][j] = dp[i - 1][j - 1] + 1;
				} else {
					dp[i][j] = Integer.max(dp[i][j - 1], dp[i - 1][j]);
				}
				result = Integer.max(result, dp[i][j]);
			}
		}
		return result;
	}
}
