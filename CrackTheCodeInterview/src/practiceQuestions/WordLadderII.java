package practiceQuestions;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

public class WordLadderII {

	/*
	 * build graph, get all path from source to end using DFS and find the shortest
	 * path, DFS exceeds time limit, use BFS
	 */
	HashMap<String, List<String>> graph;

	public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
		List<List<String>> ans = new ArrayList<>();
		if (!wordList.contains(endWord)) {
			return new ArrayList<>();
		}
		graph = buildGraph(beginWord, endWord, wordList);
		bfs(beginWord, endWord, wordList, ans);
		return ans;
	}

	public void bfs(String beginWord, String endWord, List<String> wordList, List<List<String>> ans) {
		// create a queue which stores the path
		Queue<List<String>> queue = new LinkedList<>();
		List<String> path = new ArrayList<>();
		path.add(beginWord);
		queue.offer(path);
		boolean isFound = false;
		Set<String> visited = new HashSet<>();
		visited.add(beginWord);
		while (!queue.isEmpty()) {
			int size = queue.size();
			Set<String> currentLevelVisited = new HashSet<>();
			// BFS one level
			for (int j = 0; j < size; j++) {
				List<String> p = queue.poll();
				String temp = p.get(p.size() - 1);
				for (String neighbor : graph.get(temp)) {
					// avoid go back to previous level
					if (!visited.contains(neighbor)) {
						if (neighbor.equals(endWord)) {
							isFound = true;
							p.add(neighbor);
							ans.add(new ArrayList<String>(p));
							// p is used in next neighbor cycle, restore
							p.remove(p.size() - 1);
						}
						p.add(neighbor);
						// add new path to the queue
						queue.offer(new ArrayList<String>(p));
						// p is used in next neighbor cycle, restore
						p.remove(p.size() - 1);
						currentLevelVisited.add(neighbor);
					}
				}
			}
			visited.addAll(currentLevelVisited);
			if (isFound) {
				break;
			}
		}
	}

	List<ArrayList<String>> allPath = new ArrayList<>();

	public List<List<String>> findLadders2(String beginWord, String endWord, List<String> wordList) {
		if (!wordList.contains(endWord)) {
			return new ArrayList<>();
		}
		graph = buildGraph(beginWord, endWord, wordList);
		Set<String> visited = new HashSet<>();
		ArrayList<String> currentPath = new ArrayList<>();
		currentPath.add(beginWord);
		dfs(beginWord, endWord, visited, currentPath);

		int min = Integer.MAX_VALUE;
		for (List<String> l : allPath) {
			if (l.size() < min) {
				min = l.size();
			}
		}
		List<List<String>> res = new ArrayList<>();
		for (List<String> l : allPath) {
			if (l.size() == min) {
				res.add(l);
			}
		}
		return res;
	}

	@SuppressWarnings("unchecked")
	private void dfs(String u, String d, Set<String> visited, ArrayList<String> currentPath) {
		visited.add(u); // to avoid cycle

		if (u.equals(d)) {
			allPath.add((ArrayList<String>) currentPath.clone());
			return;
		}

		for (String i : graph.get(u)) {
			if (!visited.contains(i)) {
				currentPath.add(i);
				dfs(i, d, visited, currentPath);
				visited.remove(i);
				currentPath.remove(i);
			}
		}

		visited.remove(u);
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

	public static void main(String[] args) {
		String beginWord = "hit";
		String endWord = "cog";
		List<String> wordList = Arrays.asList(new String[] { "hot", "dot", "dog", "lot", "log", "cog" });
		System.out.println(new WordLadderII().findLadders(beginWord, endWord, wordList));
	}
}
