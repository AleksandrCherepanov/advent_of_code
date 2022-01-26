package advent_2105_3

import "testing"

var testCasesAdvent2015_3_1 = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty case", "", 1},
	{"case #1", ">", 2},
	{"case #2", "^>v<", 4},
	{"case #3", "^v^v^v^v^v", 2},
}

func TestAdvent2015_3_1(t *testing.T) {
	for _, testCase := range testCasesAdvent2015_3_1 {
		result := Advent2015_3_1(testCase.input)
		if result != testCase.expected {
			t.Fatalf("%v. Expected: %v. Actual: %v", testCase.caseName, testCase.expected, result)
		}
	}
}

var testCasesAdvent2015_3_2 = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty case", "", 1},
	{"case #1", "^v", 3},
	{"case #2", "^>v<", 3},
	{"case #3", "^v^v^v^v^v", 11},
}

func TestAdvent2015_3_2(t *testing.T) {
	for _, testCase := range testCasesAdvent2015_3_2 {
		result := Advent2015_3_2(testCase.input)
		if result != testCase.expected {
			t.Fatalf("%v. Expected: %v. Actual: %v", testCase.caseName, testCase.expected, result)
		}
	}
}
