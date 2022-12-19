package advent_2022

import "testing"

func TestRope1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			13,
		},
	}

	for _, td := range testData {
		result := Rope1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestRope2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{"test #1",
			[]string{
				"R 5",
				"U 8",
				"L 8",
				"D 3",
				"R 17",
				"D 10",
				"L 25",
				"U 20",
			},
			36,
		},
	}

	for _, td := range testData {
		result := Rope2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
