package main

// question: https://leetcode.com/problems/my-calendar-ii/
type MyCalendarTwo struct {
	intervals [][]int
	overlaps  [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{intervals: make([][]int, 0), overlaps: make([][]int, 0)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, overlap := range this.overlaps {
		if start < overlap[1] && end > overlap[0] {
			return false
		}
	}

	for _, interval := range this.intervals {
		if start < interval[1] && end > interval[0] {
			this.overlaps = append(this.overlaps, []int{max(start, interval[0]), min(end, interval[1])})
		}
	}
	this.intervals = append(this.intervals, []int{start, end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
