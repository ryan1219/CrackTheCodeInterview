package main

// question: https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		newArea := (right - left)
		if height[right] < height[left] {
			newArea = newArea * height[right]
		} else {
			newArea = newArea * height[left]
		}

		if newArea > maxArea {
			maxArea = newArea
		}

		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
