package advent_2022

import "testing"

func TestTuning1(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"test #1",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
		},
		{
			"test #2",
			"nppdvjthqldpwncqszvftbrmjlhg",
			6,
		},
		{
			"test #3",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
		},
		{
			"test #4",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
		},
		{
			"test #5",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			7,
		},
	}

	for _, td := range testData {
		result := Tuning1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestTuning2(t *testing.T) {
	testData := []struct {
		name     string
		input    string
		expected int
	}{
		{
			"test #1",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			23,
		},
		{
			"test #2",
			"nppdvjthqldpwncqszvftbrmjlhg",
			23,
		},
		{
			"test #3",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			29,
		},
		{
			"test #4",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			26,
		},
		{
			"test #5",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			19,
		},
	}

	for _, td := range testData {
		result := Tuning2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
