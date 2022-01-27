package advent_2015_4

import "testing"

var testCasesAdvent2015_4 = []struct {
	caseName string
	input    string
	compare  string
	expected int
}{
	{"empty case", "", "00000", 0},
	{"case #1", "abcdef", "00000", 609043},
	{"case #2", "pqrstuv", "00000", 1048970},
}

func TestAdvent2015_4(t *testing.T) {
	for _, testCase := range testCasesAdvent2015_4 {
		result := Advent2015_4(testCase.input, testCase.compare)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}
