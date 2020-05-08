package practiceQuestions;

/*
 * question: https://leetcode.com/problems/maximum-product-subarray/
 * 
 * Kadane's algorithm
 * https://www.youtube.com/watch?v=86CQq3pKSUw&t=511s
 */
public class MaximumProductSubarray {
	
	/*
	 * max/min product of array ends at index i
	 */
	public int maxProduct(int[] a) {
		if (a == null || a.length == 0)
			return 0;

		int ans = a[0], min = ans, max = ans;

		for (int i = 1; i < a.length; i++) {
			if (a[i] >= 0) {
				max = Math.max(a[i], max * a[i]);
				min = Math.min(a[i], min * a[i]);
			} else {
				int tmp = max;
				max = Math.max(a[i], min * a[i]);
				min = Math.min(a[i], tmp * a[i]);
			}

			ans = Math.max(ans, max);
		}

		return ans;
	}
}
