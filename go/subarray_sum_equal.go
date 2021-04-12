package main

// question: https://leetcode.com/problems/subarray-sum-equals-k/
func subarraySum(nums []int, k int) int {
	preFix := make(map[int]int)
	res := 0
	sum := 0
	preFix[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if value, contained := preFix[sum-k]; contained {
			res += value
		}

		if _, contained := preFix[sum]; contained {
			preFix[sum] += 1
		} else {
			preFix[sum] = 1
		}
	}
	return res
}
