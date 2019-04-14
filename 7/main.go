package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {

	d := x

	result := 0

	for d != 0 {

		result *= 10
		result += d % 10

		d /= 10

	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
