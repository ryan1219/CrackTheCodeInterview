package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")

	test := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(test))
}
