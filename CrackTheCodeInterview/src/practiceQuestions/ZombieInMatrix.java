package practiceQuestions;

import java.util.LinkedList;
import java.util.Queue;

/*
 * https://leetcode.com/discuss/interview-question/411357/
 * 
 * Given a 2D grid, each cell is either a zombie 1 or a human 0. Zombies can turn adjacent (up/down/left/right) human beings into zombies every hour. Find out how many hours does it take to infect all humans?

Example:

Input:
[[0, 1, 1, 0, 1],
 [0, 1, 0, 1, 0],
 [0, 0, 0, 0, 1],
 [0, 1, 0, 0, 0]]

Output: 2

Explanation:
At the end of the 1st hour, the status of the grid:
[[1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1],
 [0, 1, 0, 1, 1],
 [1, 1, 1, 0, 1]]

At the end of the 2nd hour, the status of the grid:
[[1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1],
 [1, 1, 1, 1, 1]]
 
 */
public class ZombieInMatrix {

	public static void main(String[] args) {
		int[][] grid = { { 0, 1, 1, 0, 1 }, { 0, 1, 0, 1, 0 }, { 0, 0, 0, 0, 1 }, { 0, 1, 0, 0, 0 } };
		System.out.println(minDays(grid));
	}

	public static int minDays(int[][] grid) {
		int days = 0;
		int all = 0;
		int zombie = 0;
		Queue<Point> queue = new LinkedList<>();
		int[][] positions = { { 0, 1 }, { 0, -1 }, { 1, 0 }, { -1, 0 } };

		for (int x = 0; x < grid.length; x++) {
			for (int y = 0; y < grid[x].length; y++) {
				if (grid[x][y] == 1) {
					queue.add(new Point(x, y));
					zombie++;
				}
				all++;
			}
		}

		while (!queue.isEmpty()) {
			int oneDay = queue.size();
			if (zombie == all) {
				return days;
			}
			while (oneDay > 0) {
				Point p = queue.remove();
				for (int[] dir : positions) {
					int x = p.x + dir[0];
					int y = p.y + dir[1];
					if (x >= 0 && x < grid.length && y >= 0 && y < grid[0].length && grid[x][y] == 0) {
						zombie++;
						grid[x][y] = 1;
						queue.add(new Point(x, y));
					}
				}
				oneDay--;
			}
			days++;
		}
		return -1;
	}
}

class Point {
	int x;
	int y;

	public Point(int x, int y) {
		this.x = x;
		this.y = y;
	}
}