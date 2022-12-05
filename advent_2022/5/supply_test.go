package advent_2022

import "testing"

func TestSupply1(t *testing.T) {
	testData := []struct {
		name     string
		steps    []string
		cargoes  [][]string
		expected string
	}{
		{
			"test #1",
			[]string{
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			[][]string{
				{"Z", "N"},
				{"M", "C", "D"},
				{"P"},
			},
			"CMZ",
		},
	}

	for _, td := range testData {
		result := Supply1(td.steps, td.cargoes)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestSupply2(t *testing.T) {
	testData := []struct {
		name     string
		steps    []string
		cargoes  [][]string
		expected string
	}{
		{
			"test #1",
			[]string{
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			[][]string{
				{"Z", "N"},
				{"M", "C", "D"},
				{"P"},
			},
			"MCD",
		},
	}

	for _, td := range testData {
		result := Supply2(td.steps, td.cargoes)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}
