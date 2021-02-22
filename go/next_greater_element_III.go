package main

import "strconv"

// question: https://leetcode.com/problems/next-greater-element-iii/
func nextGreaterElement(n int) int {
	r := []rune(strconv.Itoa(n))
	i := len(r) - 2
	for i >= 0 && r[i] >= r[i+1] {
		i--
	}

	if i < 0 {
		return -1
	}

	j := len(r) - 1
	for r[j] <= r[i] {
		j--
	}

	swap(r, i, j)
	i++
	j = len(r) - 1

	for i < j {
		swap(r, i, j)
		i++
		j--
	}
	v, e := strconv.ParseInt(string(r), 10, 32)
	if e != nil {
		return -1
	}

	return int(v)
}

func swap(n []rune, i, j int) {
	tmp := n[i]
	n[i] = n[j]
	n[j] = tmp
}
