package practiceQuestions;

import java.util.Iterator;
import java.util.List;

/* 
Given two iterators that return SORTED integers, build an iterator with a `next` method that
iterates over the intersection of it1 and it2, i.e. only the elements that appear in both.
Example:
Iterator<Integer> i1 = testIter(1, 2, 3, 3, 4, 5, 6, 10)
Iterator<Integer> i2 = testIter(2, 3, 3, 10, 11)
Output:
IntersectionIterator iter = IntersectionIterator(i1, i2)
iter.next() -> 2
iter.next() -> 3
iter.next() -> 10
*/
public class IntersectionIterator {
	private Iterator<Integer> it1;
	private Iterator<Integer> it2;
	private Integer next;
	private List<Integer> valuesReturned;

	public IntersectionIterator(Iterator<Integer> it1, Iterator<Integer> it2) {
		this.it1 = it1;
		this.it2 = it2;
	}

	public Integer next() throws Exception {
		if (next != null) {
			Integer value = this.next;
			next = null;
			valuesReturned.add(value);
			return value;
		}

		if (!it1.hasNext() && !it2.hasNext()) {
			this.next = null;
			throw new Exception();
		}

		Integer v1 = it1.next();
		Integer v2 = it2.next();

		while (it1.hasNext() || it2.hasNext()) {
			if (v1 < v2) {
				Integer v1Past = v1;
				v1 = it1.next();
			}
			if (v2 < v1) {
				v2 = it2.next();

			}

			if (v1 == v2) {
				valuesReturned.add(v1);
				if (valuesReturned.contains(v1)) {
					return next();
				}
				return v1;
			}
		}

		this.next = null;
		throw new Exception();
	}

	public boolean hasNext() {
		if (next != null) {
			return true;
		}

		try {
			next = next();
		} catch (Exception e) {
			return false;
		}
		return true;
	}
}
