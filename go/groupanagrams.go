package main

import (
	"fmt"
	"strings"
)

/*
question: https://leetcode.com/problems/group-anagrams/submissions/
*/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	result := make([][]string, 0)
	for _, str := range strs {
		key := encode(str)
		if _, ok := m[key]; !ok {
			m[key] = make([]string, 0)
		}
		m[key] = append(m[key], str)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

func encode(src string) string {
	var bitVector [26]int
	for _, r := range src {
		bitVector[r-97] += 1
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(bitVector)), ","), "[]")
}
