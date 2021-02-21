package main

import (
	"sort"
	"strings"
	"unicode"
)

// https://leetcode.com/problems/reorder-data-in-log-files/
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		a := strings.Split(logs[i], " ")
		b := strings.Split(logs[j], " ")
		aIsDigit := isDigit(a[1])
		bIsDigit := isDigit(b[1])

		if !aIsDigit && bIsDigit {
			return true
		}

		if !aIsDigit && !bIsDigit {
			aContent := strings.Join(a[1:], " ")
			bContent := strings.Join(b[1:], " ")

			if aContent == bContent {
				return a[0] < b[0]
			}
			return aContent < bContent
		}

		return false
	})

	return logs
}

func isDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
