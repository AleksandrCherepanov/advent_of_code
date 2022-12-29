package advent2015_1

import "testing"

var testCases1 = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty input", "", 0},
	{"case #1", "(())", 0},
	{"case #2", "()()", 0},
	{"case #3", "(((", 3},
	{"case #4", "(()(()(", 3},
	{"case #5", "))(((((", 3},
	{"case #6", "())", -1},
	{"case #7", "))(", -1},
	{"case #8", ")())())", -3},
	{"case #9", ")))", -3},
}

func TestLisp1(t *testing.T) {
	for _, testCase := range testCases1 {
		result := Lisp1(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}

var testCases2 = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty input", "", 0},
	{"case #1", ")", 1},
	{"case #2", "()())", 5},
	{"case #3", "(((", 3},
	{"case #4", "(()(()(", 7},
	{"case #5", "))(((((", 1},
	{"case #6", "())", 3},
	{"case #7", "))(", 1},
	{"case #8", ")())())", 1},
	{"case #9", ")))", 1},
}

func TestLisp2(t *testing.T) {
	for _, testCase := range testCases2 {
		result := Lisp2(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}
