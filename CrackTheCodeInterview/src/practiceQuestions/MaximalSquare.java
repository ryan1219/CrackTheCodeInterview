package practiceQuestions;

public class MaximalSquare {

	/*
	 * DP
	 */
	public int maximalSquare(char[][] matrix) {
		if (matrix.length < 1) {
			return 0;
		}
		int[][] dp = new int[matrix.length][];
		int max = 0;
		for (int i = 0; i < dp.length; i++) {
			dp[i] = new int[matrix[i].length];
		}

		// base
		int i;
		int j;
		i = 0;
		for (j = 0; j < matrix[i].length; j++) {
			dp[i][j] = Character.getNumericValue(matrix[i][j]);
			max = Math.max(dp[i][j], max);
		}
		j = 0;
		for (i = 0; i < matrix.length; i++) {
			dp[i][j] = Character.getNumericValue(matrix[i][j]);
			max = Math.max(dp[i][j], max);
		}

		for (i = 1; i < matrix.length; i++) {
			for (j = 1; j < matrix[i].length; j++) {
				if (matrix[i][j] == '1') {
					dp[i][j] = Math.min(Math.min(dp[i - 1][j], dp[i][j - 1]), dp[i - 1][j - 1]) + 1;
					if (max < dp[i][j]) {
						max = dp[i][j];
					}
				} else {
					dp[i][j] = 0;
				}
			}
		}

		return max * max;
	}
}
