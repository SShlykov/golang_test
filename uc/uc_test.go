package uc

import "testing"

type ucTest struct {
	input,output string
}

var testCases = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"Go", "GO"},
	ucTest{"Antwerp", "ANTWERP"},
}

func TestUppercase(t *testing.T) {
	for _,expected := range testCases {
		result := UpperCase(expected.input)
		if result != expected.output {
			t.Errorf("\nThe Result of testing failed:\n " +
				"Function Uppercase(%s) returned \"%s\", but must be %s", expected.input, result, expected.output)
		}
	}
}