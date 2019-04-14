package main

import "testing"

func TestReverse(t *testing.T) {

	var testData = []struct {
		Input  int
		Output int
	}{
		{
			Input:  123,
			Output: 321,
		},
		{
			Input:  -123,
			Output: -321,
		},
		{
			Input:  120,
			Output: 21,
		},
		{
			Input:  25564,
			Output: 46552,
		},
		{
			Input:  1534236469,
			Output: 0,
		},
	}

	for _, value := range testData {
		res := reverse(value.Input)

		if res != value.Output {
			t.Errorf("%d after reverse %d isn't equal %d", value.Input, res, value.Output)
		}
	}
}
