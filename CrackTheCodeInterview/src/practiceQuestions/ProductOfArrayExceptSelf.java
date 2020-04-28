package practiceQuestions;

/*
 * question: https://leetcode.com/problems/product-of-array-except-self/
 */
public class ProductOfArrayExceptSelf {

	/*
	 * 1 2 3 4
	 */
	public int[] productExceptSelf(int[] nums) {
		int[] res = new int[nums.length];
		res[0] = 1;
		int left = 1;
		for (int i = 1; i < nums.length; i++) {
			res[i] = nums[i - 1] * left;
			left = res[i];
		}

		int right = 1;
		for (int i = nums.length - 2; i >= 0; i--) {
			right = nums[i + 1] * right;
			res[i] = res[i] * right;
		}
		return res;
	}
}
