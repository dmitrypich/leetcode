package main

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	var testData = []struct {
		Input  string
		Output []string
	}{
		{
			Input:  "cbbd",
			Output: []string{"bb"},
		},
		{
			Input:  "abcbdc",
			Output: []string{"cbc", "bcb"},
		},
		{
			Input:  "abb",
			Output: []string{"bb"},
		},
		{
			Input:  "",
			Output: []string{""},
		},
		{
			Input:  "a",
			Output: []string{"a"},
		},
		{
			Input:  "ccc",
			Output: []string{"ccc"},
		},
		{
			Input:  "cccc",
			Output: []string{"cccc"},
		},
		{
			Input:  "azwdzwmwcqzgcobeeiphemqbjtxzwkhiqpbrprocbppbxrnsxnwgikiaqutwpftbiinlnpyqstkiqzbggcsdzzjbrkfmhgtnbujzszxsycmvipjtktpebaafycngqasbbhxaeawwmkjcziybxowkaibqnndcjbsoehtamhspnidjylyisiaewmypfyiqtwlmejkpzlieolfdjnxntonnzfgcqlcfpoxcwqctalwrgwhvqvtrpwemxhirpgizjffqgntsmvzldpjfijdncexbwtxnmbnoykxshkqbounzrewkpqjxocvaufnhunsmsazgibxedtopnccriwcfzeomsrrangufkjfzipkmwfbmkarnyyrgdsooosgqlkzvorrrsaveuoxjeajvbdpgxlcrtqomliphnlehgrzgwujogxteyulphhuhwyoyvcxqatfkboahfqhjgujcaapoyqtsdqfwnijlkknuralezqmcryvkankszmzpgqutojoyzsnyfwsyeqqzrlhzbc",
			Output: []string{"sooos"},
		},
		{
			Input:  "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg",
			Output: []string{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"},
		},
	}

	for _, v := range testData {
		result := longestPalindrome(v.Input)

		found := false
		for _, value := range v.Output {
			if value == result {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("in row %s  %s isn't equal %s", v.Input, result, v.Output)
		}
	}
}
