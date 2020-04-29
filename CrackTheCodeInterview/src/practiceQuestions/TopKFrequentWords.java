package practiceQuestions;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashMap;
import java.util.List;
import java.util.Map.Entry;

/*
 * question: https://leetcode.com/problems/top-k-frequent-words/
 */
public class TopKFrequentWords {
	public List<String> topKFrequent(String[] words, int k) {
		HashMap<String, Integer> freq = new HashMap<>();
		for (int i = 0; i < words.length; i++) {
			if (freq.containsKey(words[i])) {
				freq.put(words[i], freq.get(words[i]) + 1);
			} else {
				freq.put(words[i], 1);
			}
		}

		List<Entry<String, Integer>> list = new ArrayList<>(freq.entrySet());

		list.sort(new Comparator<Entry<String, Integer>>() {

			@Override
			public int compare(Entry<String, Integer> o1, Entry<String, Integer> o2) {
				if (o1.getValue() != o2.getValue()) {
					return Integer.compare(o2.getValue(), o1.getValue());
				} else {
					return o1.getKey().compareTo(o2.getKey());
				}
			}

		});
		List<String> res = new ArrayList<>();
		for (int i = 0; i < k; i++) {
			res.add(list.get(i).getKey());
		}
		return res;
	}
}
