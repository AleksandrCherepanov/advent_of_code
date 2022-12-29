package advent_2015

import "testing"

func TestSticks1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				`""`,
				`"abc"`,
				`"aaa\"aaa"`,
				`"\x27"`,
			},
			12,
		},
	}

	for _, td := range testData {
		result := Sticks1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestSticks2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				`""`,
				`"abc"`,
				`"aaa\"aaa"`,
				`"\x27"`,
			},
			19,
		},
	}

	for _, td := range testData {
		result := Sticks2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
