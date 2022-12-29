package advent_2105_3

import "testing"

var testCasesVacuum1 = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty case", "", 1},
	{"case #1", ">", 2},
	{"case #2", "^>v<", 4},
	{"case #3", "^v^v^v^v^v", 2},
}

func TestVacuum1(t *testing.T) {
	for _, testCase := range testCasesVacuum1 {
		result := Vacuum1(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}

var testCasesVacuum2 = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty case", "", 1},
	{"case #1", "^v", 3},
	{"case #2", "^>v<", 3},
	{"case #3", "^v^v^v^v^v", 11},
}

func TestVacuum2(t *testing.T) {
	for _, testCase := range testCasesVacuum2 {
		result := Vacuum2(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}
