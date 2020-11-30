package practiceQuestions;

import java.util.Arrays;

/*
 * https://leetcode.com/problems/squares-of-a-sorted-array/
 * Given an array of integers A sorted in non-decreasing order, 
 * return an array of the squares of each number, also in sorted non-decreasing order.
 * 
 * Example 1:

	Input: [-4,-1,0,3,10]
	Output: [0,1,9,16,100]
	
   Example 2:
	
	Input: [-7,-3,2,3,11]
	Output: [4,9,9,49,121]
 */
public class SquareOfSortedArray {
	public static void main(String[] args) {
		int[] input = { -4, -1, 0, 3, 10 };
		System.out.println(Arrays.toString(solution2(input)));
	}

	public static int[] solution1(int[] src) {
		int[] out = new int[src.length];
		for (int i = 0; i < src.length; i++) {
			out[i] = src[i] * src[i];
		}
		Arrays.sort(out);
		return out;
	}

	public static int[] solution2(int[] src) {
		int[] out = new int[src.length];
		int count = out.length - 1;
		int i = 0;
		int j = src.length - 1;
		while (i <= j) {
			if (src[i] * src[i] > src[j] * src[j]) {
				out[count] = src[i] * src[i];
				count -= 1;
				i += 1;
			} else {
				out[count] = src[j] * src[j];
				count -= 1;
				j -= 1;
			}
		}
		return out;
	}
}
