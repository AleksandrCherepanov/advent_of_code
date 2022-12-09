package advent_2022

import "testing"

func TestTreetop1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			21,
		},
	}

	for _, td := range testData {
		result := Treetop1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestTreetop2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"30373",
				"25512",
				"65332",
				"33549",
				"35390",
			},
			8,
		},
	}

	for _, td := range testData {
		result := Treetop2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
