package sorting;

import java.util.ArrayList;
import java.util.Arrays;

public class MergeSort {

	public static void main(String[] args) {
		int[] a = { 9, 800, 1, 6, 5, -4, 0 };
		System.out.println(Arrays.toString(a));
		System.out.println(Arrays.toString(mergeSort(a)));
	}

	public static int[] mergeSort(int[] a) {
		if (a.length == 1) {
			return a;
		}

		int[] a1 = Arrays.copyOfRange(a, 0, a.length / 2);
		int[] a2 = Arrays.copyOfRange(a, a.length / 2, a.length);

		a1 = mergeSort(a1);
		a2 = mergeSort(a2);

		return merge(a1, a2);
	}

	public static int[] merge(int[] a1, int[] a2) {
		ArrayList<Integer> result = new ArrayList<>();

		int i = 0;
		int j = 0;

		while (i < a1.length && j < a2.length) {
			if (a1[i] < a2[j]) {
				result.add(a1[i]);
				i++;
			} else {
				result.add(a2[j]);
				j++;
			}
		}

		// check the remaining in a1
		while (i < a1.length) {
			result.add(a1[i]);
			i++;
		}

		// check the remaining in a2
		while (j < a2.length) {
			result.add(a2[j]);
			j++;
		}

		return result.stream().mapToInt(Integer::valueOf).toArray();
	}
}
