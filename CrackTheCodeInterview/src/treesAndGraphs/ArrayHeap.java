package treesAndGraphs;

import java.util.ArrayList;

public class ArrayHeap implements Heap {

	private ArrayList<Integer> heap = new ArrayList<>();

	private int parent(int j) {
		return (j - 1) / 2;
	}

	private int left(int j) {
		return 2 * j + 1;
	}

	private boolean hasLeft(int j) {
		return left(j) < heap.size();
	}

	private int right(int j) {
		return 2 * j + 2;
	}

	private boolean hasRight(int j) {
		return right(j) < heap.size();
	}

	private void swap(int i, int j) {
		int temp = heap.get(i);
		heap.set(i, heap.get(j));
		heap.set(j, temp);
	}

	private void heapifyUp(int j) {
		while (j > 0) {
			int p = parent(j);
			if (heap.get(j) > heap.get(p)) {
				break;
			}
			swap(j, p);
			j = p;
		}
	}

	private void heapifyDown(int j) {
		while (hasLeft(j)) {
			int minIndex = left(j);
			if (hasRight(j)) {
				if (heap.get(left(j)) > heap.get(right(j))) {
					minIndex = right(j);
				}
			}

			if (heap.get(j) < heap.get(minIndex)) {
				break;
			}

			swap(j, minIndex);
			j = minIndex;
		}
	}

	@Override
	public void insert(int value) {
		heap.add(value);
		heapifyUp(heap.size() - 1);
	}

	@Override
	public int remove() throws Exception {
		if (heap.isEmpty()) {
			throw new Exception();
		}
		int result = heap.get(0);
		swap(0, heap.size() - 1);
		heap.remove(heap.size() - 1);
		heapifyDown(0);
		return result;

	}

}
