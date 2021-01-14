package main

/*
https://leetcode.com/problems/minimum-window-substring/
*/
func minWindow(s string, t string) string {

	min := len(s) + 1
	min_i := len(s) + 1
	min_j := 0
	src := []rune(s)
	i := 0
	m := make(map[rune]int)
	counter := len(t)

	for _, c := range t {
		_, found := m[c]
		if found {
			m[c] = m[c] + 1
		} else {
			m[c] = 1
		}
	}

	for j := 0; j < len(src); j++ {
		_, found := m[src[j]]
		if found {
			if m[src[j]] > 0 {
				counter -= 1
			}
			m[src[j]] = m[src[j]] - 1
		}

		for counter == 0 {
			if j-i < min {
				min = j - i
				min_i = i
				min_j = j
			}

			_, found := m[src[i]]
			if found {
				m[src[i]] = m[src[i]] + 1
				if m[src[i]] > 0 {
					counter += 1
				}
			}

			i += 1
		}
	}

	if min == len(s)+1 {
		return ""
	} else {
		return s[min_i : min_j+1]
	}
}
