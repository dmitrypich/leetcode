package main

import "fmt"

func main() {
	h := []byte("hello")
	h2 := []byte("he")

	fmt.Println(string(h))
	reverseString(h)
	fmt.Println(string(h))

	fmt.Println(string(h2))
	reverseString(h2)
	fmt.Println(string(h2))

}

func reverseString(s []byte) {

	lastIndex := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[lastIndex-i] = s[lastIndex-i], s[i]
	}
}
