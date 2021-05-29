package main

// question: https://leetcode.com/problems/my-calendar-i/
/*
same approach as my calendar II
*/
type MyCalendar struct {
	intervals [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{intervals: make([][]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, interval := range this.intervals {
		if start < interval[1] && end > interval[0] {
			return false
		}
	}

	this.intervals = append(this.intervals, []int{start, end})
	return true
}
