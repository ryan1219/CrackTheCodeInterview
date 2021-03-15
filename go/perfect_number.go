package main

// question: https://leetcode.com/problems/perfect-number/
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i != num/i {
				sum += num / i
			}
		}
	}

	return sum == num
}
