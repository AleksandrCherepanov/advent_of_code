package advent_2022

import "testing"

func TestCamp1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,3-7",
				"6-6,4-6",
				"2-6,4-8",
			},
			2,
		},
	}

	for _, td := range testData {
		result := Camp1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestCamp2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,3-7",
				"6-6,4-6",
				"2-6,4-8",
			},
			4,
		},
	}

	for _, td := range testData {
		result := Camp2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
