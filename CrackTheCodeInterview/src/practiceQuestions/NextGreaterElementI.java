package practiceQuestions;

import java.util.HashMap;

/*
 * question: https://leetcode.com/problems/next-greater-element-i/
 * 
 * explanation:
 * Input: nums1 = [2,4], nums2 = [1,2,3,4]
 * Output: [3,-1]
 * 
 * 2 in nums1 is index 1 in nums2, starts from index 1 in nums2, the next value greater than 2
 * 4 in nums1 is index 4 in nums2, starts from index 4 in nums2, the next value greater than 4
 */
public class NextGreaterElementI {

	public int[] nextGreaterElement(int[] nums1, int[] nums2) {
		int[] re = new int[nums1.length];
		HashMap<Integer, Integer> map = new HashMap<>();
		for (int i = 0; i < nums2.length; i++) {
			map.put(nums2[i], i);
		}
		for (int i = 0; i < re.length; i++) {
			re[i] = -1;
			for (int j = map.get(nums1[i]) + 1; j < nums2.length; j++) {
				if (nums2[j] > nums1[i]) {
					re[i] = nums2[j];
					break;
				}
			}
		}
		return re;
	}
}
