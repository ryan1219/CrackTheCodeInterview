package main

import "fmt"

func countingSort(input []int) []int {
	max := input[0]

	// find the largest value in input
	for _, val := range input {
		if val > max {
			max = val
		}
	}
	freq := make([]int, max+1)

	// build frequency array
	for _, val := range input {
		freq[val]++
	}

	// build auxiliary array
	for i := 1; i < len(freq); i++ {
		freq[i] += freq[i-1]
	}

	// sort
	// reverse, because of the application to radix sorting, it is important for counting sort to be a stable sort:
	// repeated elements in the same order that they appear in the input
	sorted := make([]int, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		sorted[freq[input[i]]-1] = input[i]
		freq[input[i]]--
	}

	return sorted
}

func main() {
	test := []int{2, 3, 1, 1, 4, 0}
	fmt.Println(countingSort(test))
}
