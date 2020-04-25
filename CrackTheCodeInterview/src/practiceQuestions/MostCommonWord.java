package practiceQuestions;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map.Entry;
import java.util.stream.Collectors;

/*
 * question: https://leetcode.com/problems/most-common-word/
 * 
 * time: O(n + m) - n is the size of paragraph, m is the size of banned
 */
public class MostCommonWord {

	List<String> retrieveMostFrequentlyUsedWords(String helpText, List<String> wordsToExclude) {
		List<String> words = new ArrayList<>();
		List<String> wordsToExcludeLowercase = wordsToExclude.stream().map(String::toLowerCase)
				.collect(Collectors.toList());
		int highestFrq = 0;
		// create a list of words from helpText, exclude the words in excluded list
		String word = "";
		for (char c : helpText.toCharArray()) {
			if (Character.isLetter(c)) {
				word += c;
			} else if (word.length() > 0) {
				if (!wordsToExcludeLowercase.contains(word.toLowerCase())) {
					words.add(word.toLowerCase());
				}
				word = "";
			}
		}
		// add the last word
		words.add(word.toLowerCase());

		// use map to count the frequency of each word
		HashMap<String, Integer> count = new HashMap<>();
		for (String s : words) {
			count.put(s, count.getOrDefault(s, 0) + 1);
		}

		// convert map entry to a list, sort the list on entry value.
		List<Entry<String, Integer>> sorted = new ArrayList<>(count.entrySet());
		sorted.sort(Entry.comparingByValue());
		// last element in list is the highest frequency word
		List<String> result = new ArrayList<>();
		highestFrq = sorted.get(sorted.size() - 1).getValue();
		result.add(sorted.get(sorted.size() - 1).getKey());
		// any word has the same frequency as the highest frequency will also be added
		// to result
		for (int i = sorted.size() - 2; i >= 0; i--) {
			if (sorted.get(i).getValue() == highestFrq) {
				result.add(sorted.get(i).getKey());
			}
		}
		System.out.println(sorted);

		return result;
	}
}
