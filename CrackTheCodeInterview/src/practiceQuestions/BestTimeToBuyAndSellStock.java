package practiceQuestions;

public class BestTimeToBuyAndSellStock {

	/*
	 * iterate each price, one a lower price find, this may be a buy price, if a
	 * higher price find, this may be a sell price.
	 */
	public int maxProfit(int[] prices) {
		if (prices == null || prices.length <= 1)
			return 0;
		int ans = 0, low = Integer.MAX_VALUE, high = Integer.MAX_VALUE;
		for (int i = 0; i < prices.length; i++) {
			if (prices[i] < low)
				low = high = prices[i];// if we have a lower price, update both low and high
			else if (prices[i] > high)
				high = prices[i];// if price is higher than high, update only high
			ans = Math.max(high - low, ans);
		}
		return ans;
	}
}
