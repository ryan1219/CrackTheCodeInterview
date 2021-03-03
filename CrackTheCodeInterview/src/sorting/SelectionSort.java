package sorting;

import java.util.Arrays;

public class SelectionSort {

	public static void main(String[] args) {
		int[] a = { 1, 100, 20, 3 };
		System.out.println(Arrays.toString(a));
		selectionSort(a);
		System.out.println(Arrays.toString(a));
	}

	/*
	 * Find the smallest element in the array and exchange it with first element
	 * Then find the smallest element in the sub-array a[1] and swap it with a[1]
	 */
	public static void selectionSort(int[] a) {
		for (int i = 0; i < a.length; i++) {
			int min = i;
			for (int j = i; j < a.length; j++) {
				if (a[j] < a[i]) {
					min = j;
				}
			}
			swap(a, i, min);
		}
	}

	private static void swap(int[] a, int i, int j) {
		int tmp = a[i];
		a[i] = a[j];
		a[j] = tmp;
	}

}
