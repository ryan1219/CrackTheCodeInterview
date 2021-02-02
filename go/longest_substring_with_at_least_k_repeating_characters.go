package main

// question: https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/
func longestSubstring(s string, k int) int {
	res := 0
	// substring can contains unique chars from 1 to 26
	for i := 1; i <= 26; i++ {
		tmp := longestSubstringWithKCharacters(s, k, i)
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func longestSubstringWithKCharacters(s string, k, uniqChars int) int {
	var m [26]int
	var i, j, unique, charMoreThank, res int
	for j < len(s) {
		char := s[j] - 'a'
		m[char]++
		if m[char] == 1 {
			unique++
		}
		if m[char] == k {
			charMoreThank++
		}
		j++
		for unique > uniqChars {
			char = s[i] - 'a'
			m[char]--
			if m[char] == 0 {
				unique--
			}
			if m[char] == k-1 {
				charMoreThank--
			}
			i++
		}
		if unique == charMoreThank && j-i > res {
			res = j - i
		}
	}

	return res
}
