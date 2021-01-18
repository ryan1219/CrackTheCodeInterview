package main

// question: https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func findSubstring(s string, words []string) []int {
	standard := make(map[string]int)
	wordLength := len(words[0])
	result := make([]int, 0)
	for _, word := range words {
		val, find := standard[word]
		if !find {
			standard[word] = 1
		} else {
			standard[word] = val + 1
		}
	}

	for i := 0; i <= len(s)-len(words)*wordLength; i++ {
		if checkSubstring(s[i:i+len(words)*wordLength], standard, wordLength) {
			result = append(result, i)
		}
	}

	return result
}

func checkSubstring(sub string, standard map[string]int, wordLength int) bool {
	m := make(map[string]int)

	for i := 0; i <= len(sub)-wordLength; i = i + wordLength {
		key := sub[i : i+wordLength]
		val, find := m[key]
		if find {
			if val >= standard[key] {
				return false
			}
			m[key] = val + 1
		} else {
			if standard[key] > 0 {
				m[key] = 1
			} else {
				return false
			}
		}
	}

	return true
}
