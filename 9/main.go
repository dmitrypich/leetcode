package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	d := x

	result := 0

	for d != 0 {

		result *= 10
		result += d % 10

		d /= 10

	}

	return result == x
}
