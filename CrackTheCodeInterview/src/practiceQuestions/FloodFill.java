package practiceQuestions;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

/*
 * question:https://leetcode.com/problems/flood-fill/
 */
public class FloodFill {
	public static void main(String[] args) {
		int[][] a = {
				{0,0,0},
				{0,1,1}
		};
		System.out.println(Arrays.deepToString(new FloodFill().floodFill(a, 1, 1, 1)));
	}
	public int[][] floodFill(int[][] image, int sr, int sc, int newColor) {
		Queue<Point> queue = new LinkedList<>();
		queue.add(new Point(sr, sc));
		int startColor = image[sr][sc];
		// if new color is same as original color, no need to perform flood fill
		// also visited tracking is not needed since all color same as original color will be changed
		if (startColor == newColor) {
			return image;
		}
		while (!queue.isEmpty()) {
			int round = queue.size();
			while (round > 0) {
				Point p = queue.remove();
				// set color on p
				image[p.x][p.y] = newColor;
				// check four direction of point p
				if (p.x + 1 < image.length && image[p.x + 1][p.y] == startColor) {
					queue.add(new Point(p.x + 1, p.y));
				}
				if (p.y + 1 < image[p.x].length && image[p.x][p.y + 1] == startColor) {
					queue.add(new Point(p.x, p.y + 1));
				}
				if (p.x - 1 >= 0 && image[p.x - 1][p.y] == startColor) {
					queue.add(new Point(p.x - 1, p.y));
				}
				if (p.y - 1 >= 0 && image[p.x][p.y - 1] == startColor) {
					queue.add(new Point(p.x, p.y - 1));
				}
				round--;
			}
		}
		return image;
	}

	class Point {
		int x;
		int y;

		public Point(int x, int y) {
			this.x = x;
			this.y = y;
		}
	}
}
