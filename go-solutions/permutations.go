package main

/*
 	Given a collection of distinct integers, return all possible permutations.
 	question: https://leetcode.com/problems/permutations/
 	Input: [1,2,3]
	Output:
	[
	[1,2,3],
	[1,3,2],
	[2,1,3],
	[2,3,1],
	[3,1,2],
	[3,2,1]
	]
*/
func permute(nums []int) [][]int {
	results := make([][]int, 0)
	results = backtrack(nums, 0, results)

	return results
}

func backtrack(nums []int, start int, results [][]int) [][]int {
	if start == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		results = append(results, tmp)
		return results
	}

	for i := start; i < len(nums); i++ {
		swap(nums, start, i)
		results = backtrack(nums, start+1, results)
		swap(nums, i, start)
	}

	return results
}

func swap(nums []int, i int, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

/*
Given a collection of numbers that might contain duplicates, return all possible unique permutations.
Question: https://leetcode.com/problems/permutations-ii/
Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

explain:
Let's look at another example, permute[1,2,3,4,1].

At recursion level 0, we have [1] + permute[2,3,4,1], [2] + permute[1,3,4,1], [3] + permute[2,1,4,1], [4] + permute[2,3,1,1], [1] + permute[2,3,4,1].

1 has already been at the 1st index of current recursion level, so the last possibility is redundant. We can use a hash set to mark which elements have been at the 1st index of current recursion level, so that if we meet the element again, we can just skip it.
*/
func permuteUnique(nums []int) [][]int {
	results := make([][]int, 0)
	results = backtrack2(nums, 0, results)

	return results
}

func backtrack2(nums []int, start int, results [][]int) [][]int {
	if start == len(nums)-1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		results = append(results, tmp)
		return results
	}

	visited := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if _, ok := visited[nums[i]]; ok {
			continue
		}
		visited[nums[i]] = true
		swap(nums, start, i)
		results = backtrack2(nums, start+1, results)
		swap(nums, i, start)
	}

	return results
}
