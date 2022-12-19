package advent_2022

import "testing"

func TestDistress1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"[1,1,3,1,1]",
				"[1,1,5,1,1]",
				"",
				"[[1],[2,3,4]]",
				"[[1],4]",
				"",
				"[9]",
				"[[8,7,6]]",
				"",
				"[[4,4],4,4]",
				"[[4,4],4,4,4]",
				"",
				"[7,7,7,7]",
				"[7,7,7]",
				"",
				"[]",
				"[3]",
				"",
				"[[[]]]",
				"[[]]",
				"",
				"[1,[2,[3,[4,[5,6,7]]]],8,9]",
				"[1,[2,[3,[4,[5,6,0]]]],8,9]",
			},
			13,
		},
	}

	for _, td := range testData {
		result := Distress1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestDistress2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"[1,1,3,1,1]",
				"[1,1,5,1,1]",
				"",
				"[[1],[2,3,4]]",
				"[[1],4]",
				"",
				"[9]",
				"[[8,7,6]]",
				"",
				"[[4,4],4,4]",
				"[[4,4],4,4,4]",
				"",
				"[7,7,7,7]",
				"[7,7,7]",
				"",
				"[]",
				"[3]",
				"",
				"[[[]]]",
				"[[]]",
				"",
				"[1,[2,[3,[4,[5,6,7]]]],8,9]",
				"[1,[2,[3,[4,[5,6,0]]]],8,9]",
			},
			140,
		},
	}

	for _, td := range testData {
		result := Distress2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed.\nExpected:\n%v\nActual:\n%v\n", td.name, td.expected, result)
		}
	}
}
