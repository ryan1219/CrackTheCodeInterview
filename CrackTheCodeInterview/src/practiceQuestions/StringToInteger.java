package practiceQuestions;

/*
 * question: https://leetcode.com/problems/string-to-integer-atoi/
 */
public class StringToInteger {
	public int myAtoi(String str) {
		char[] items = str.trim().toCharArray();
		boolean first = true;
		int sign = 1;
		StringBuilder sb = new StringBuilder();
		for (int i = 0; i < items.length; i++) {
			if (first) {
				first = false;
				if (items[i] == '-') {
					sign = -1;
				} else if (items[i] == '+') {
					sign = 1;
				} else if (items[i] == '0') {
					continue;
				} else if ('0' < items[i] && items[i] <= '9') {
					sb.append(items[i]);
				} else {
					return 0;
				}
			} else {
				if ('0' <= items[i] && items[i] <= '9') {
					sb.append(items[i]);
				} else {
					break;
				}
			}
		}

		while (sb.length() > 0 && sb.charAt(0) == '0') {
			sb.delete(0, 1);
		}
		String stringValue = sb.toString();
		System.out.println(stringValue);
		int digit = 1;
		int value = 0;
		for (int i = stringValue.length() - 1; i >= 0; i--) {
			try {
				int temp = Math.multiplyExact(digit, Character.getNumericValue(stringValue.charAt(i)));
				value = Math.addExact(value, temp);
				System.out.println(i + ":" + value);
			} catch (ArithmeticException e) {
				if (sign == 1) {
					return Integer.MAX_VALUE;
				} else {
					return Integer.MIN_VALUE;
				}
			}
			try {
				digit = Math.multiplyExact(digit, 10);
			} catch (ArithmeticException e) {
				if (i != 0) {
					if (sign == 1) {
						return Integer.MAX_VALUE;
					} else {
						return Integer.MIN_VALUE;
					}
				}
			}
		}

		return value * sign;
	}

	public static void main(String[] args) {
		System.out.println(new StringToInteger().myAtoi("-6147483648"));
	}
}
