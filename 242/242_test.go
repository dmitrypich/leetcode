package main

import "testing"

func TestValidAnagram(t *testing.T) {
	var testData = []struct {
		Input1 string
		Input2 string
		Output bool
	}{
		{
			Input1: "anagram",
			Input2: "nagaram",
			Output: true,
		},
		{
			Input1: "rat",
			Input2: "car",
			Output: false,
		},
	}

	for _, value := range testData {
		res := isAnagram(value.Input1, value.Input2)
		if res != value.Output {
			t.Errorf("%s %s isAnagram %t isn't equal %t", value.Input1, value.Input2, res, value.Output)
		}
	}
}
