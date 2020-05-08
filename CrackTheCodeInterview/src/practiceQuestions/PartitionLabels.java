package practiceQuestions;

import java.util.ArrayList;
import java.util.List;

/*
 * question: https://leetcode.com/problems/partition-labels/
 */
public class PartitionLabels {
	public List<Integer> partitionLabels(String S) {
		if (S == null || S.length() <= 0) {
			return new ArrayList<>();
		}

		int[] map = new int[26];
		for (int i = 0; i < S.length(); i++) {
			map[S.charAt(i) - 'a'] = i;
		}

		int start = 0;
		int end = 0;
		List<Integer> res = new ArrayList<>();
		for (int i = 0; i < S.length(); i++) {
			end = Math.max(end, map[S.charAt(i) - 'a']);
			if(i == end) {
				res.add(end - start + 1);
				start = end + 1;
			}
		}
		return res;
	}
}
