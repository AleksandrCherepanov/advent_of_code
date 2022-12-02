package advent_2022

import "testing"

func TestRps1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"A Y",
				"B X",
				"C Z",
			},
			15,
		},
	}

	for _, td := range testData {
		result := Rps1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestRps2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"A Y",
				"B X",
				"C Z",
			},
			12,
		},
	}

	for _, td := range testData {
		result := Rps2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
