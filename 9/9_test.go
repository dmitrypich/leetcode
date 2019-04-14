package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	var testData = []struct {
		Input  int
		Output bool
	}{
		{
			Input:  121,
			Output: true,
		},
		{
			Input:  -121,
			Output: false,
		},
		{
			Input:  10,
			Output: false,
		},
		{
			Input:  0121,
			Output: false,
		},
	}

	for _, value := range testData {
		res := isPalindrome(value.Input)

		if res != value.Output {
			t.Errorf("%d IsPalindrome %t isn't equal %t", value.Input, res, value.Output)
		}
	}
}
