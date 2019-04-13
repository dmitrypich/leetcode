package main

import "testing"

func TestLength(t *testing.T) {
	testData := []struct {
		Row   string
		Count int
	}{
		{
			Row:   "abcabcbb",
			Count: 3,
		},
		{
			Row:   "bbbbb",
			Count: 1,
		},
		{
			Row:   "pwwkew",
			Count: 3,
		},
	}

	for _, value := range testData {
		result := lengthOfLongestSubstring(value.Row)
		if result != value.Count {
			t.Errorf("row %s result %d isn't equal %d", value.Row, result, value.Count)
		}
	}
}
