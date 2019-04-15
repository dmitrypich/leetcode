package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("a", "ab"))
}

func isAnagram(s string, t string) bool {

	m1 := make(map[rune]int)

	for _, value := range s {
		m1[value]++
	}

	for _, value := range t {
		m1[value]--
	}

	for _, value := range m1 {
		if value != 0 {
			return false
		}
	}

	return true
}
