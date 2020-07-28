package practiceQuestions;

public class JumpGameII {
	public static void main(String[] args) {
		int[] test = { 2, 3, 1, 1, 4 };
		System.out.println(new JumpGameII().jump(test));
	}

	public int jump(int[] nums) {
		int[] jumps = new int[nums.length];
		for (int i = 0; i < nums.length; i++) {
			int step = 1;
			while (step <= nums[i]) {
				if (i + step == jumps.length - 1) {
					return jumps[i] + 1;
				} else if (i + step < jumps.length) {
					if (jumps[i + step] == 0 || jumps[i] + 1 < jumps[i + step]) {
						jumps[i + step] = jumps[i] + 1;
					}
				} else {
					break;
				}
				step++;
			}
		}
//		System.out.println(Arrays.toString(jumps));

		return jumps[jumps.length - 1];
	}
	/*
	 * public int jump(int[] A) {
		int jumps = 0, curEnd = 0, curFarthest = 0;
		for (int i = 0; i < A.length - 1; i++) {
			curFarthest = Math.max(curFarthest, i + A[i]);
			if (i == curEnd) {
				jumps++;
				curEnd = curFarthest;

				if (curEnd >= A.length - 1) {
					break;
				}
			}
		}
		return jumps;
	}
	 */
}
