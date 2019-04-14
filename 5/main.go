package main

import (
	"fmt"
)

func main() {

	fmt.Println(longestPalindrome("ccfc"))
}

func longestPalindrome(s string) string {

	s1, size1 := checkEvenPalindrome(s)

	s2, size2 := checkNotEvenPalindrome(s)

	if size1 > size2 {
		return s1
	}

	return s2
}

func checkEvenPalindrome(s string) (string, int) {

	runes := []rune(s)

	size := len(runes)

	if size <= 1 {
		return s, len(runes)
	}

	isPalindrome := true
	for i := 0; i < size/2; i++ {
		if runes[i] != runes[size-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		return s, len(runes)
	}

	var palindromeSize, k, j int
	for i := 1; i < size; i++ {
		var right, left int

		checkForEvenPalindrome := runes[i-1] == runes[i]

		if checkForEvenPalindrome {

			left, right = i-1, i
		}

		if !checkForEvenPalindrome {
			continue
		}

		for right < size && left >= 0 && runes[right] == runes[left] {

			if right-left > palindromeSize {
				k = left
				j = right

				palindromeSize = right - left
			}

			right++
			left--

		}

	}

	return string(runes[k : j+1]), j - k
}

func checkNotEvenPalindrome(s string) (string, int) {

	runes := []rune(s)

	size := len(runes)

	if size <= 1 {
		return s, len(runes)
	}

	isPalindrome := true
	for i := 0; i < size/2; i++ {
		if runes[i] != runes[size-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		return s, len(runes)
	}

	//isEven := (size % 2) == 0

	var palindromeSize, k, j int
	for i := 1; i < size; i++ {
		var right, left int

		checkForNotEvenPalindrome := i < size-1 && runes[i-1] == runes[i+1]

		if checkForNotEvenPalindrome {

			left, right = i-1, i+1
		}

		if !checkForNotEvenPalindrome {
			continue
		}

		for right < size && left >= 0 && runes[right] == runes[left] {

			if right-left > palindromeSize {
				k = left
				j = right

				palindromeSize = right - left
			}

			right++
			left--

		}

	}

	return string(runes[k : j+1]), j - k
}
