package practiceQuestions;

import java.util.HashMap;
import java.util.Map.Entry;

/*
 * question: https://leetcode.com/problems/majority-element/
 */
public class MajorityElement {

	/*
	 * solution1: hashmap counting
	 */
	public int majorityElement(int[] nums) {
		int threshold = nums.length / 2;
		HashMap<Integer, Integer> map = new HashMap<>();
		for (int i = 0; i < nums.length; i++) {
			if(map.containsKey(nums[i])) {
				map.put(nums[i], map.get(nums[i]) + 1);
			}
			else {
				map.put(nums[i], 1);
			}
		}
		
		for(Entry<Integer, Integer> entry: map.entrySet()) {
			if(entry.getValue() > threshold) {
				return entry.getKey();
			}
		}
		
		return -1;
	}
}
