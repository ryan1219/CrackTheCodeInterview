package practiceQuestions;

import java.util.Arrays;

public class ThreeSumClosest {

	/*
	 * solution1: find all 3 sum value, find the closest to the target
	 */
	
	/*
	 * solution2: sort; for each number find other two numbers can make the sum of three which is closes to target
	 */
	public int threeSumClosest(int[] nums, int target) {
		int closest = nums[0] + nums[1] + nums[2];

		Arrays.sort(nums);

		for (int i = 0; i < nums.length - 2; i++) {
			int j = i + 1;
			int k = nums.length - 1;
			while (j < k) {
				int sum = nums[i] + nums[j] + nums[k];
				if (Math.abs(target - sum) < Math.abs(target - closest)) {
					closest = sum;
				}
				if (sum > target) {
					k--;
				} else {
					j++;
				}
			}
		}
		return closest;
	}
	
	public static void main(String[] args) {
		int [] a = {1, 1, -1, -1, 3};
		
		System.out.println(new ThreeSumClosest().threeSumClosest(a, -1));
	}
}
