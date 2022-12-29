package advent_2015

import "testing"

func TestIntern1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"ugknbfddgicrmopn",
				"aaa",
				"jchzalrnumimnmhp",
				"haegwjzuvuyypxyu",
				"dvszwmarrgswjxmb",
			},
			2,
		},
	}

	for _, td := range testData {
		result := Intern1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestIntern2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"qjhvhtzxzqqjkmpb",
				"xxyxx",
				"uurcxstgmygtbstg",
				"ieodomkazucvgmuy",
			},
			2,
		},
	}

	for _, td := range testData {
		result := Intern2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
