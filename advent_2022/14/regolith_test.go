package advent_2022

import "testing"

func TestRegolith1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"498,4 -> 498,6 -> 496,6",
				"503,4 -> 502,4 -> 502,9 -> 494,9",
			},
			24,
		},
	}

	for _, td := range testData {
		result := Regolith1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestRegolith2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"498,4 -> 498,6 -> 496,6",
				"503,4 -> 502,4 -> 502,9 -> 494,9",
			},
			93,
		},
	}

	for _, td := range testData {
		result := Regolith2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed.\nExpected:\n%v\nActual:\n%v\n", td.name, td.expected, result)
		}
	}
}
