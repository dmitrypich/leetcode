package main

import "testing"

func TestAtoi(t *testing.T) {

	var testData = []struct {
		Input  string
		Result int
	}{
		{
			Input:  "42",
			Result: 42,
		},
		{
			Input:  "           -42",
			Result: -42,
		},
		{
			Input:  "4193 with words",
			Result: 4193,
		},
		{
			Input:  "words and 987",
			Result: 0,
		},
		{
			Input:  "-91283472332",
			Result: -2147483648,
		},
		{
			Input:  "0",
			Result: 0,
		},
		{
			Input:  "-1 ",
			Result: -1,
		},
		{
			Input:  "+1 ",
			Result: 1,
		},
		{
			Input:  "+-1 ",
			Result: 0,
		},
		{
			Input:  "-+1 ",
			Result: 0,
		},
		{
			Input:  "9223372036854775808",
			Result: 2147483647,
		},
		{
			Input:  "2147483648",
			Result: 2147483647,
		},
		{
			Input:  "-2147483648",
			Result: -2147483648,
		},
		{
			Input:  "-2147483649",
			Result: -2147483648,
		},
		{
			Input:  "-2147483647",
			Result: -2147483647,
		},
	}

	for _, value := range testData {
		res := myAtoi(value.Input)

		if res != value.Result {
			t.Errorf("%s after myAtoi %d isn't equal %d", value.Input, res, value.Result)
		}
	}
}
