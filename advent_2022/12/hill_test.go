package advent_2022

import "testing"

func TestHill1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"Sabqponm",
				"abcryxxl",
				"accszExk",
				"acctuvwj",
				"abdefghi",
			},
			31,
		},
	}

	for _, td := range testData {
		result := Hill1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestHill2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"Sabqponm",
				"abcryxxl",
				"accszExk",
				"acctuvwj",
				"abdefghi",
			},
			29,
		},
	}

	for _, td := range testData {
		result := Hill2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed.\nExpected:\n%v\nActual:\n%v\n", td.name, td.expected, result)
		}
	}
}
