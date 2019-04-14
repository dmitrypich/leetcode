package main

import "testing"

func TestMaxArea(t *testing.T) {
	var testData = []struct {
		Array []int
		Area  int
	}{
		{
			Array: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Area:  49,
		},
	}

	for _, value := range testData {
		res := maxArea(value.Array)

		if res != value.Area {
			t.Errorf("max area of %d after maxArea %d isn't equald %d", value.Array, res, value.Area)
		}
	}
}
