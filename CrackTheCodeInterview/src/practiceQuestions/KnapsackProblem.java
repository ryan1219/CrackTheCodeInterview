package practiceQuestions;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

/*
 * 
 */
public class KnapsackProblem {

	/*
	 * solution1: DP, top-down approach - recursive with memorization
	 */
	/**
	 * @param items  number of remaining items
	 * @param weight available weight capacity
	 * @return value
	 */

	List<Integer> values = new ArrayList<>();
	List<Integer> weights = new ArrayList<>();
	HashMap<String, Integer> map = new HashMap<>();

	int knapsack(int items, int weight) {
		if (map.containsKey(items + ":" + weight)) {
			return map.get(items + ":" + weight);
		}
		// base case
		int result = 0;
		if (items == 0 || weight == 0) {
			result = 0;
		} else if (weight < values.get(items)) {
			result = knapsack(items - 1, weight);
		} else {
			result = Math.max(knapsack(items - 1, weight), values.get(items) + knapsack(items - 1, weights.get(items)));
		}
		map.put(items + ":" + weight, result);
		return result;
	}

	/*
	 * solution2: DP, bottom-up approach
	 * 
	 * i = item w = maxWeight
	 * 
	 * v[i][w] = max(v[i-1][w], value i + v[i-1][w - weight i]) if w > weight i
	 *      or = v[i-1][w]
	 * 
	 * base case if i = 0 or w = 0 v[i][w] = 0
	 */

}
