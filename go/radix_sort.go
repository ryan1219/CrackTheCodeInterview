package main

import "fmt"

func radixSort(input []int) {
	max := input[0]
	// find the largest value in input
	for _, val := range input {
		if val > max {
			max = val
		}
	}
	// counting sort for each digit, instead of passing digit number as array, pass exp, use exp to calculate the digit
	for exp := 1; max/exp > 0; exp = exp * 10 {
		countingSort(input, exp)
	}
}

func countingSort(input []int, exp int) {
	freq := make([]int, 10)
	tmpOutput := make([]int, len(input))
	// build frequency array
	for i := 0; i < len(input); i++ {
		freq[(input[i]/exp)%10]++
	}

	// build auxiliary array
	for i := 1; i < len(freq); i++ {
		freq[i] += freq[i-1]
	}

	for i := len(input) - 1; i >= 0; i-- {
		tmpOutput[freq[(input[i]/exp)%10]-1] = input[i]
		freq[(input[i]/exp)%10]--
	}

	for i := 0; i < len(input); i++ {
		input[i] = tmpOutput[i]
	}
}

func main() {
	test := []int{1, 20, 40, 10, 100}
	radixSort(test)
	fmt.Println(test)
}
