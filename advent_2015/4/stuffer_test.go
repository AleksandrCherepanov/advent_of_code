package advent_2015_4

import "testing"

var testCasesStuffer = []struct {
	caseName string
	input    string
	compare  string
	expected int
}{
	{"empty case", "", "00000", 0},
	{"case #1", "abcdef", "00000", 609043},
	{"case #2", "pqrstuv", "00000", 1048970},
}

func TestStuffer(t *testing.T) {
	for _, testCase := range testCasesStuffer {
		result := Stuffer(testCase.input, testCase.compare)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}
