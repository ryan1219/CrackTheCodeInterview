package practiceQuestions;

import java.util.TreeMap;

class SummaryRanges {

    TreeMap<Integer, int[]> tree;

    /** Initialize your data structure here. */
    public SummaryRanges() {
        tree = new TreeMap<>();
    }

    public void addNum(int val) {
        if (tree.containsKey(val))
            return;
        Integer l = tree.lowerKey(val);
        Integer h = tree.higherKey(val);

        // merge two intervals
        if (l != null && h != null && tree.get(l)[1] + 1 == val && h - 1 == val) {
            tree.get(l)[1] = tree.get(h)[1];
            tree.remove(h);
        } else if (l != null && tree.get(l)[1] + 1 >= val) {
            tree.get(l)[1] = Math.max(tree.get(l)[1], val);
        } else if (h != null && h - 1 == val) {
            tree.put(val, new int[] { val, tree.get(h)[1] });
            tree.remove(h);
        } else {
            tree.put(val, new int[] { val, val });
        }

    }

    public int[][] getIntervals() {
        int[][] res = new int[tree.keySet().size()][];
        int i = 0;
        for (int[] interval : tree.values()) {
            res[i] = interval;
            i++;
        }
        return res;
    }
}
