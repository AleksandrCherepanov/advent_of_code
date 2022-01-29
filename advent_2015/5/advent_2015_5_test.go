package advent_2015_5

import "testing"

var testCasesAdvent2015_5_1 = []struct {
	caseName string
	input    string
	expected bool
}{
	{"empty case", "", false},
	{"case #1", "ugknbfddgicrmopn", true},
	{"case #2", "jchzalrnumimnmhp", false},
	{"case #3", "haegwjzuvuyypxyu", false},
	{"case #4", "dvszwmarrgswjxmb", false},
}

func TestAdvent2015_5_1(t *testing.T) {
	for _, testCase := range testCasesAdvent2015_5_1 {
		result := Advent2015_5_1(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}
