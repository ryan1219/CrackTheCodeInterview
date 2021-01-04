package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func makeData(length int) []uint8 {
	a := make([]uint8, length)
	for i := 0; i < length; i++ {
		a[i] = uint8(rand.Int() & 0xff)
	}
	return a
}

/*
sorts the input array of integers such that all the odd elements
appear before all the even elements, for example [1,2,3,4,5] would
be sorted to [1,3,5,2,4]  (answer of [3,1,5,4,2] is also acceptable, the
only requirement is that all the odd numbers appear before any even number in the list).
*/
func sortByOddEven(a []uint8) {
	sort.Slice(a, func(i, j int) bool {
		if a[i]%2 == 1 {
			return true
		}
		return false
	})
}

func insert(a []uint8, index int, value uint8) []uint8 {
	if len(a) == 0 {
		a = append(a, value)
	} else {
		a = append(a[:index+1], a[index:]...)
		a[index] = value
	}
	return a
}

/*
returns the value that occurs most often (with highest frequency)
in the input array (the statistical "mode")
For example, if the input is [1,2,1,2,2,3], the mode is 2, because it occurs
three times, which is more than any other element.
*/
func findMostFrequentElement(a []uint8) uint8 {
	dic := make(map[uint8]int)
	for _, value := range a {
		if _, ok := dic[value]; ok {
			dic[value] = int(value) + 1
		} else {
			dic[value] = 1
		}
	}
	var largestKey uint8 = 0
	var largestValue int = 0
	for key, value := range dic {
		if value > largestValue {
			largestValue = value
			largestKey = key
		}
	}

	return largestKey
}

/*
returns the value of the second most frequently occurring element.
Similar to the "mode" function above, but for the input [1,2,1,2,2,3]
the result would be 1, since it occurs twice (with 2 being the most
frequently occurring element with a total of three occurrences)
*/
func findSecondMostFrequentElement(a []uint8) uint8 {
	dic := make(map[uint8]int)
	for _, value := range a {
		if _, ok := dic[value]; ok {
			dic[value] = int(value) + 1
		} else {
			dic[value] = 1
		}
	}
	var secondLargestKey uint8 = 0
	var secondLargestValue int = 0
	var largestKey uint8 = 0
	var largestValue int = 0
	fmt.Println(dic)
	for key, value := range dic {
		if value > largestValue {
			secondLargestValue = largestValue
			secondLargestKey = largestKey
			largestValue = value
			largestKey = key
			fmt.Println(value)
		} else if value > secondLargestValue {
			secondLargestValue = value
			secondLargestKey = key
			fmt.Println(key)
		}
	}

	return secondLargestKey
}

func main() {
	length := 6
	a := makeData(length)
	sortByOddEven(a)
	test := []uint8{1, 2, 1, 2, 2, 3}
	x := findMostFrequentElement(test)
	y := findSecondMostFrequentElement(test)

	for i := 0; i < length; i++ {
		fmt.Printf("%d,", int(a[i]))
	}
	fmt.Printf("\n\n %d %d \n\n", x, y)
}
