package main

import (
	"fmt"
)

func main() {

	d := []int{2, 3, 4, 5, 18, 17, 6}

	fmt.Println(maxArea(d))

	d1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(d1))
}

func maxArea(height []int) int {

	if len(height) <= 1 {
		return 0
	}

	left := 0
	right := len(height) - 1

	resultMaxArea := 0

	for left != right {
		minHigh := minInt(height[left], height[right])
		area := minHigh * (right - left)

		if area > resultMaxArea {
			resultMaxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return resultMaxArea
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}
