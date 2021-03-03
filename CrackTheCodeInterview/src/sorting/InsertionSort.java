package sorting;

import java.util.Arrays;

public class InsertionSort {

	public static void main(String[] args) {
		int[] a = { 1000, 100, 20, 3 };
		System.out.println(Arrays.toString(a));
		insertionSort(a);
		System.out.println(Arrays.toString(a));
	}

	public static void insertionSort(int[] a) {
		if (a.length <= 1) {
			return;
		}

		for (int i = 1; i < a.length; i++) {
			int pivit = i;

			for (int j = i; j >= 0; j--) {
				if (a[pivit] < a[j]) {
					swap(a, pivit, j);
					pivit = j;
				}
			}
		}
	}

	public static void swap(int[] a, int i, int j) {
		int tmp = a[i];
		a[i] = a[j];
		a[j] = tmp;
	}

}
