package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}

func lengthOfLongestSubstring(s string) int {

	array := []rune(s)
	hashMap := make(map[rune]int)
	result := 0
	size := len(array)

	for i, j := 0, 0; i < size && j < size; {

		_, exist := hashMap[array[j]]
		if !exist {
			hashMap[array[j]]++
			j++
			result = max(result, j-i)
		} else {
			delete(hashMap, array[i])
			i++
		}

	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
