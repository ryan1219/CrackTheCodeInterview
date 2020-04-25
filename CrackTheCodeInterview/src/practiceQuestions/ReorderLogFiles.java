package practiceQuestions;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/*
 * question: https://leetcode.com/articles/reorder-log-files/
 * 
 * time: number of line is n, sorting 0(nlog(n))
 */
public class ReorderLogFiles {
	
	public List<String> reorderLines(int logFileSize, List<String> logLines) {
		// split letter log and number log
		List<String> letterLog = new ArrayList<>();
		List<String> numberLog = new ArrayList<>();

		for (String log : logLines) {
			String[] items = log.split(" ");
			if (Character.isLetter(items[1].charAt(0))) {
				letterLog.add(log);
			} else {
				numberLog.add(log);
			}
		}

		// create a custom comparator
		Comparator<String> customeComparator = new Comparator<String>() {
			@Override
			public int compare(String s1, String s2) {
				String[] list1 = s1.split(" ");
				String[] list2 = s2.split(" ");
				String id1 = list1[0];
				String id2 = list2[0];

				// compare each word by alphanumeric order
				int i;
				for (i = 1; i < Math.min(list1.length, list2.length); i++) {
					if (!list1[i].equals(list2[i])) {
						return list1[i].compareTo(list2[i]);
					}
				}
				// tie case
				if (list1.length < list2.length) {
					return -1;
				}

				return id1.compareTo(id2);
			}
		};

		letterLog.sort(customeComparator);
		ArrayList<String> result = new ArrayList<>();
		for (String log : letterLog) {
			result.add(log);
		}

		for (String log : numberLog) {
			result.add(log);
		}

		return result;
	}
}
