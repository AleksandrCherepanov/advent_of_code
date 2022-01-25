package advent2015_2

import "testing"

var testCasesCalcualteSquareFeet = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty case", "", 0},
	{"case #1", "2x3x4", 58},
	{"case #2", "1x1x10", 43},
}

func TestCalcSquareFeet(t *testing.T) {
	for _, testCase := range testCasesCalcualteSquareFeet {
		result := calcSquareFeet(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}

var testCasesAdvent2015_2_1 = []struct {
	caseName string
	input    []string
	expected int
}{
	{"empty case", []string{}, 0},
	{"case #1", []string{"2x3x4", "1x1x10"}, 101},
}

func TestAdvent2015_2_1(t *testing.T) {
	for _, testCase := range testCasesAdvent2015_2_1 {
		result := Advent2015_2_1(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}

var testCasesCalcRibbon = []struct {
	caseName string
	input    []int
	expected int
}{
	{"empty case", []int{}, 0},
	{"case #1", []int{2, 3, 4}, 24},
	{"case #2", []int{1, 1, 10}, 10},
}

func TestCalcRibbon(t *testing.T) {
	for _, testCase := range testCasesCalcRibbon {
		result := calcRibbon(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}

var testCasesFeetOfRibbon = []struct {
	caseName string
	input    string
	expected int
}{
	{"empty case", "", 0},
	{"case #1", "2x3x4", 34},
	{"case #2", "1x1x10", 14},
}

func TestFindFeetOfRibbon(t *testing.T) {
	for _, testCase := range testCasesFeetOfRibbon {
		result := calcFeetOfRibbon(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}

var testCasesAdvent2015_2_2 = []struct {
	caseName string
	input    []string
	expected int
}{
	{"empty case", []string{}, 0},
	{"case #1", []string{"2x3x4", "1x1x10"}, 48},
}

func TestAdvent2015_2_2(t *testing.T) {
	for _, testCase := range testCasesAdvent2015_2_2 {
		result := Advent2015_2_2(testCase.input)
		if result != testCase.expected {
			t.Fatalf(`%v. Expected: %v. Actual: %v`, testCase.caseName, testCase.expected, result)
		}
	}
}
