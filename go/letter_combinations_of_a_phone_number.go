package main

// question: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}

	chars := [][]string{[]string{}, []string{}, []string{"a", "b", "c"}, []string{"d", "e", "f"}, []string{"g", "h", "i"}, []string{"j", "k", "l"},
		[]string{"m", "n", "o"}, []string{"p", "q", "r", "s"}, []string{"t", "u", "v"}, []string{"w", "x", "y", "z"}}

	res = append(res, "")

	for i := 0; i < len(digits); i++ {
		temp := make([]string, 0)
		digit := chars[digits[i]-'0']
		for j := 0; j < len(digit); j++ {
			for k := 0; k < len(res); k++ {
				temp = append(temp, res[k]+digit[j])
			}
		}
		res = temp
	}

	return res
}
