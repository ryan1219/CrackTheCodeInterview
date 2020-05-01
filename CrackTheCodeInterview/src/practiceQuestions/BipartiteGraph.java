package practiceQuestions;

import java.util.Stack;

public class BipartiteGraph {
	public boolean isBipartite(int[][] graph) {
		if (graph.length < 1) {
			return false;
		}
		int[] seen = new int[graph.length];
		for (int j = 0; j < graph.length; j++) {
			if(seen[j] != 0) {
				continue;
			}
			// 0 not seen, 1 red, -1 black
			Stack<Integer> stack = new Stack<>();
			stack.push(j);
			seen[j] = 1;
			while (!stack.isEmpty()) {
				int current = stack.pop();
				int[] neighbors = graph[current];
				for (int i = 0; i < neighbors.length; i++) {
					if (seen[neighbors[i]] == 0) {
						seen[neighbors[i]] = -1 * seen[current];
						stack.push(neighbors[i]);
					} else if (seen[neighbors[i]] == seen[current]) {
						return false;
					}
				}
			}
		}

		return true;
	}
}
