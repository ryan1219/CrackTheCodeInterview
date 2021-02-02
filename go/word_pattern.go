package main

import "strings"

// quesiton: https://leetcode.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	patternList := strings.Split(pattern, "")
	sList := strings.Split(s, " ")

	m1 := make(map[string]string)
	m2 := make(map[string]string)

	if len(patternList) != len(sList) {
		return false
	}

	for i := 0; i < len(patternList); i++ {
		val, contains := m1[patternList[i]]
		if contains && sList[i] != val {
			return false
		} else if !contains {
			// add new key, make sure value is not duplicate
			_, contains = m2[sList[i]]
			if contains {
				return false
			}
			m1[patternList[i]] = sList[i]
			m2[sList[i]] = patternList[i]
		}
	}

	return true
}
