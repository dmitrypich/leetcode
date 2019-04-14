package main

import "testing"

func TestConvert(t *testing.T) {

	var testData = []struct {
		InputText string
		RowCount  int
		Output    string
	}{
		{
			InputText: "PAYPALISHIRING",
			RowCount:  3,
			Output:    "PAHNAPLSIIGYIR",
		},
		{
			InputText: "PAYPALISHIRING",
			RowCount:  4,
			Output:    "PINALSIGYAHRPI",
		},
	}

	for _, value := range testData {
		res := convert(value.InputText, value.RowCount)

		if res != value.Output {
			t.Errorf("input %s %d convert: %s isn't equal %s", value.InputText, value.RowCount, res, value.Output)
		}
	}
}
