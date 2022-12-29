package advent_2022

import "testing"

func TestBeacon1(t *testing.T) {
	testData := []struct {
		name     string
		y        int
		input    []string
		expected int
	}{
		{
			"test #1",
			10,
			[]string{
				"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
				"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
				"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
				"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
				"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
				"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
				"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
				"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
				"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
				"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
				"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
				"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
				"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
				"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
			},
			26,
		},
	}

	for _, td := range testData {
		result := Beacon1(td.y, td.input)
		if result != td.expected {
			t.Fatalf("%s failed. Expected: %v, Actual: %v\n", td.name, td.expected, result)
		}
	}
}

func TestBeacon2(t *testing.T) {
	testData := []struct {
		name     string
		y        int
		input    []string
		expected int64
	}{
		{
			"test #1",
			20,
			[]string{
				"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
				"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
				"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
				"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
				"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
				"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
				"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
				"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
				"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
				"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
				"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
				"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
				"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
				"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
			},
			56000011,
		},
	}

	for _, td := range testData {
		result := Beacon2(td.y, td.input)
		if result != td.expected {
			t.Fatalf("%s failed.\nExpected:\n%v\nActual:\n%v\n", td.name, td.expected, result)
		}
	}
}
