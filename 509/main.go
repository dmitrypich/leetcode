package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	first := 1
	second := 1
	for i := 2; i < n; i++ {
		first, second = second, first+second
	}
	return second

}

// FibRe calculates the Fibonacci number by recursion
func FibRe(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return FibRe(n-1) + FibRe(n-2)
}
