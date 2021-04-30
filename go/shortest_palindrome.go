package main

// question: https://leetcode.com/problems/shortest-palindrome/
/*
brute force: find the longest palindrome starts from index 0
*/
func shortestPalindrome(s string) string {
	start := 0
	end := len(s) - 1

	for start < end {
		if isPalindrome(s[start : end+1]) {
			return reverse(s[end+1:]) + s
		}
		end--
	}

	return reverse(s[end+1:]) + s
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
