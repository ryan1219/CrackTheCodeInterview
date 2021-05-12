package main

import "sort"

// question: https://leetcode.com/problems/my-calendar-iii/
type MyCalendarThree struct {
	keys     []int
	interval map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{keys: make([]int, 0), interval: make(map[int]int)}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	if _, contained := this.interval[start]; contained {
		this.interval[start]++
	} else {
		this.interval[start] = 1
		this.keys = append(this.keys, start)
	}

	if _, contained := this.interval[end]; contained {
		this.interval[end]--
	} else {
		this.interval[end] = -1
		this.keys = append(this.keys, end)
	}

	sort.Ints(this.keys)
	booking := 0
	maxBooking := 0

	for _, key := range this.keys {
		booking += this.interval[key]

		if booking > maxBooking {
			maxBooking = booking
		}
	}

	return maxBooking
}
