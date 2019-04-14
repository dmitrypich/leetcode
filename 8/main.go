package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	fmt.Println(myAtoi("2147483648"))

	fmt.Println(unicode.IsMark('+'))
}

func myAtoi(str string) int {

	var beginParse bool

	var isNegative bool
	var result int

	runes := []rune(str)

	for key, value := range runes {
		if beginParse && !unicode.IsDigit(value) {
			break
		}

		if result > math.MaxInt32+1 {
			break
		}

		if unicode.IsSpace(value) {
			continue
		}

		if !beginParse && !(unicode.IsDigit(value) || isSing(value)) {
			return 0
		}

		if unicode.IsDigit(value) || isSing(value) {
			beginParse = true
		}

		if value == '-' {
			isNegative = true
			continue
		}

		if unicode.IsDigit(value) {
			result *= 10

			result += int(value - '0')
		}

		if beginParse && key > 0 && unicode.IsDigit(runes[key-1]) && !unicode.IsDigit(value) {
			break
		}

	}

	if isNegative {
		result *= -1
	}

	if result >= math.MaxInt32 {
		return math.MaxInt32
	}

	if result <= math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func isSing(r rune) bool {

	return r == '+' || r == '-'
}
