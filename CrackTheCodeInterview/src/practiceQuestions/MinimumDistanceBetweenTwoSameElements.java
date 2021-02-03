package practiceQuestions;

import java.util.HashMap;

/* questions: https://www.geeksforgeeks.org/maximum-distance-two-occurrences-element-array/
 * 
 * change maximum to minimum
 */
public class MinimumDistanceBetweenTwoSameElements {

	public static int minDistance(int[] positions) {
		HashMap<Integer, Integer> map = new HashMap<>();
		int minDistance = Integer.MAX_VALUE;
		for (int i = 0; i < positions.length; i++) {
			if (!map.containsKey(positions[i])) {
				map.put(positions[i], i);
			} else {
				minDistance = Math.min(minDistance, i - map.get(positions[i]));
			}
		}

		if (minDistance == Integer.MAX_VALUE) {
			minDistance = positions.length;
		}

		return minDistance;
	}
}
