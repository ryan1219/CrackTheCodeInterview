package practiceQuestions;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

public class WordLadder {

	HashMap<String, List<String>> graph;

	public int ladderLength(String beginWord, String endWord, List<String> wordList) {
		if (!wordList.contains(endWord)) {
			return 0;
		}
		graph = buildGraph(beginWord, endWord, wordList);
		return bfs(beginWord, endWord);
	}

	public int bfs(String beginWord, String endWord) {
		Queue<String> queue = new LinkedList<>();
		queue.add(beginWord);
		int level = 1;
		Set<String> visited = new HashSet<>();
		visited.add(beginWord);
		while (!queue.isEmpty()) {
//			number of nodes in current level
			int size = queue.size();
			level++;
			for (int i = 0; i < size; i++) {
				String tmp = queue.remove();
				List<String> neighbors = graph.get(tmp);
				for(String neighbor: neighbors) {
					if(endWord.equals(neighbor)) {
						return level;
					}
					if(!visited.contains(neighbor)) {
						visited.add(neighbor);
						queue.add(neighbor);
					}
				}
			}
		}
		
		// no path from beginWord to endWord
		return 0;
	}

	HashMap<String, List<String>> buildGraph(String beginWord, String endWord, List<String> wordList) {
		HashMap<String, List<String>> graph = new HashMap<>();
		Set<String> wordSet = new HashSet<>(wordList);
		wordSet.add(beginWord);
		wordSet.add(endWord);
		for (String word : wordSet) {
			graph.put(word, getNeighbors(word, wordSet));
		}

		return graph;
	}

	private List<String> getNeighbors(String cur, Set<String> wordSet) {
		List<String> res = new ArrayList<>();
		char[] chars = cur.toCharArray();
		for (int i = 0; i < chars.length; i++) {
			char old = chars[i];
			for (char c = 'a'; c <= 'z'; c++) {
				if (c == old) {
					continue;
				}
				chars[i] = c;
				String next = new String(chars);
				if (wordSet.contains(next)) {
					res.add(next);
				}
			}
			chars[i] = old;
		}
		return res;
	}
}
