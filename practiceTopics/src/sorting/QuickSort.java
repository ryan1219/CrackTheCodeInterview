package sorting;

import java.util.Arrays;

public class QuickSort {

	public static void main(String[] args) {
		// test1
//		int[] a = { 5, 1, 20, 4, 6 };
//		System.out.println(Arrays.toString(a));
//		quickSort1(a, 0, a.length - 1);
//		System.out.println(Arrays.toString(a));

		// test2
//		int[] b = { 4, 7, 8, 1, 6, 5 };
		int[] b = { 5, 1, 20, 4, 6 };
		System.out.println(Arrays.toString(b));
		quickSort2(b, 0, b.length - 1);
		System.out.println(Arrays.toString(b));

	}

	/*
	 * Lomuto partition scheme chooses a pivot that is typically the last
	 * element in the array Cormen et al. in their book Introduction to
	 * Algorithms.
	 */
	public static void quickSort1(int[] array, int lo, int hi) {
		if (lo < hi) {
			int p = partition1(array, lo, hi);
			quickSort1(array, lo, p - 1);
			quickSort1(array, p + 1, hi);
		}
	}

	public static int partition1(int[] array, int lo, int hi) {
		int pivot = array[hi];
		int i = lo - 1;
		for (int j = lo; j < hi; j++) {
			if (array[j] < pivot) {
				i++;
				swap(array, j, i);
			}
		}
		swap(array, i + 1, hi);
		return i + 1;
	}

	/*
	 * Hoare partition scheme the pivot's final location is not necessarily at
	 * the index that was returned
	 */
	public static void quickSort2(int[] array, int lo, int hi) {
		if (lo >= hi) {
			return;
		}
		int p = partition2(array, lo, hi);
		quickSort2(array, lo, p);
		quickSort2(array, p + 1, hi);
	}

	public static int partition2(int[] array, int lo, int hi) {
		int pivot = array[(lo + hi) / 2];
//		System.out.println(pivot);
		int i = lo;
		int j = hi;
		while (true) {
			while (array[i] < pivot) {
//				System.out.println(array[i]);
				i++;
			}
			while (array[j] > pivot) {
				j--;
			}
			if (i >= j) {
				return j;
			}
			swap(array, i, j);
		}
	}

	public static void swap(int[] array, int a, int b) {
		int c = array[a];
		array[a] = array[b];
		array[b] = c;
		System.out.println("SWAP RESULT");
		System.out.println(Arrays.toString(array));
	}
}
