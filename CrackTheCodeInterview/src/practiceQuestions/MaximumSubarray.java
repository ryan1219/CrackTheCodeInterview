package practiceQuestions;

/*
 * question: https://leetcode.com/problems/maximum-subarray/
 */
public class MaximumSubarray {

	/*
	 * DP
	 */
	public int maxSubArrayDP(int[] nums) {
		if(nums.length == 0) {
			return 0;
		}
		int max = nums[0];
		int currentMax = max;
		for(int i = 1; i < nums.length; i++) {
			currentMax = Math.max(nums[i], nums[i] + currentMax);
			max = Math.max(currentMax, max);
		}
		
		return max;
	}
	
	/*
	 * divide and conquer
	 * 
	 * T(n) = 2 * T(n/2) + O(n)
	 * https://www.youtube.com/watch?v=yBCzO0FpsVc
	 */
	public int maxSubArray(int[] nums) {
		return maxSumSubArray(nums, 0, nums.length - 1);
	}
	
	public int maxSumSubArray(int[] nums, int low, int high) {
		if(high == low) {
			return nums[high];
		}
		
		int mid = (high + low) / 2;
		int leftMax = maxSumSubArray(nums, low, mid);
		int rightMax = maxSumSubArray(nums, mid+1, high);
		int crossMax = crossSumSubArray(nums, low, mid, high);
		
		return Math.max(Math.max(leftMax, rightMax), crossMax);
	}
	
	public int crossSumSubArray(int[] nums, int low, int mid, int high) {
		int leftMax = Integer.MIN_VALUE;
		int localSum = 0;
		for(int i = mid; i >= low; i--) {
			localSum += nums[i];
			if(localSum > leftMax) {
				leftMax = localSum;
			}
		}
		
		int rightMax = Integer.MIN_VALUE;
		localSum = 0;
		for(int i = mid + 1; i <= high; i++) {
			localSum += nums[i];
			if(localSum > rightMax) {
				rightMax = localSum;
			}
		}
		
		return leftMax + rightMax;
	}
}
