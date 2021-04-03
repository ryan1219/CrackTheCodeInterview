package practiceQuestions;

import java.util.Arrays;

public class LongestCommonPrefix {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) {
            return "";
        }
        Arrays.sort(strs);

        String first = strs[0];
        String last = strs[strs.length - 1];

        int i = 0;
        while (i < first.length() && i < last.length()) {
            if (first.charAt(i) != last.charAt(i)) {
                break;
            }
            i++;
        }

        if (first.length() < last.length()) {
            return first.substring(0, i);
        }

        return last.substring(0, i);
    }
}