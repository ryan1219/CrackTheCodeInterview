package practiceQuestions;

import java.util.LinkedList;
import java.util.Queue;

/*
 * question: https://leetcode.com/problems/number-of-islands/
 */
public class NumberOfIslands {

	int[][] seen;

	public int numIslands(char[][] grid) {
		int islands = 0;
		seen = new int[grid.length][];
		for (int i = 0; i < grid.length; i++) {
			seen[i] = new int[grid[i].length];
		}

		for (int i = 0; i < grid.length; i++) {
			for (int j = 0; j < grid[i].length; j++) {
				if (grid[i][j] == '1' && seen[i][j] == 0) {
					islands += 1;
					bfs(grid, i, j);
				}
			}
		}
		return islands;
	}

	public void bfs(char[][] grid, int x, int y) {
		int[][] positions = { { 0, 1 }, { 1, 0 }, { -1, 0 }, { 0, -1 } };
		Queue<Point> queue = new LinkedList<>();
		queue.add(new Point(x, y));
		while (!queue.isEmpty()) {
			Point current = queue.remove();
			for (int i = 0; i < positions.length; i++) {
				int xNext = current.x + positions[i][0];
				int yNext = current.y + positions[i][1];
				if(xNext < grid.length && xNext >= 0 && yNext < grid[x].length && yNext >= 0) {
//					next position is island and not visited yet
					if(grid[xNext][yNext] == '1' && seen[xNext][yNext] == 0) {
						seen[xNext][yNext] = 1;
						queue.add(new Point(xNext, yNext));
					}
				}
			}
		}
	}

	class Point {
		int x;
		int y;

		Point(int x, int y) {
			this.x = x;
			this.y = y;
		}
	}
}
