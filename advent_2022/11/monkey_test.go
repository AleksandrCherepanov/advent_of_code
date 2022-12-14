package advent_2022

import "testing"

func TestMonkey1(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"Monkey 0:",
				"Starting items: 79, 98",
				"Operation: new = old * 19",
				"Test: divisible by 23",
				"If true: throw to monkey 2",
				"If false: throw to monkey 3",
				"Monkey 1:",
				"Starting items: 54, 65, 75, 74",
				"Operation: new = old + 6",
				"Test: divisible by 19",
				"If true: throw to monkey 2",
				"If false: throw to monkey 0",
				"Monkey 2:",
				"Starting items: 79, 60, 97",
				"Operation: new = old * old",
				"Test: divisible by 13",
				"If true: throw to monkey 1",
				"If false: throw to monkey 3",
				"Monkey 3:",
				"Starting items: 74",
				"Operation: new = old + 3",
				"Test: divisible by 17",
				"If true: throw to monkey 0",
				"If false: throw to monkey 1",
			},
			10605,
		},
	}

	for _, td := range testData {
		result := Monkey1(td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestMonkey2(t *testing.T) {
	testData := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"test #1",
			[]string{
				"Monkey 0:",
				"Starting items: 79, 98",
				"Operation: new = old * 19",
				"Test: divisible by 23",
				"If true: throw to monkey 2",
				"If false: throw to monkey 3",
				"Monkey 1:",
				"Starting items: 54, 65, 75, 74",
				"Operation: new = old + 6",
				"Test: divisible by 19",
				"If true: throw to monkey 2",
				"If false: throw to monkey 0",
				"Monkey 2:",
				"Starting items: 79, 60, 97",
				"Operation: new = old * old",
				"Test: divisible by 13",
				"If true: throw to monkey 1",
				"If false: throw to monkey 3",
				"Monkey 3:",
				"Starting items: 74",
				"Operation: new = old + 3",
				"Test: divisible by 17",
				"If true: throw to monkey 0",
				"If false: throw to monkey 1",
			},
			2713310158,
		},
	}

	for _, td := range testData {
		result := Monkey2(td.input)
		if result != td.expected {
			t.Fatalf("%s failed.\nExpected:\n%v\nActual:\n%v\n", td.name, td.expected, result)
		}
	}
}
