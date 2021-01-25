package main

// question: https://leetcode.com/problems/candy/
func candy(ratings []int) int {
	sum := 1
	if len(ratings) == 0 {
		return 0
	}

	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(candies); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
			sum += candies[i-1] + 1
		} else {
			candies[i] = 1
			sum++
		}
	}

	for i := len(candies) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candies[i-1] < candies[i]+1 {
			sum = sum + candies[i] + 1 - candies[i-1]
			candies[i-1] = candies[i] + 1
		}
	}

	return sum
}
