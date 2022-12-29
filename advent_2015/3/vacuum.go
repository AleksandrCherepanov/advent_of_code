package advent_2105_3

import "strconv"

var instructions = map[string][]int{
	"^": {0, 1},
	">": {1, 0},
	"v": {0, -1},
	"<": {-1, 0},
}

func Vacuum1(input string) int {
	x := 0
	y := 0
	visited := make(map[string]int, 0)
	visited["00"] = 1
	for _, runeValue := range input {
		stringValue := string(runeValue)
		xy := instructions[stringValue]
		x += xy[0]
		y += xy[1]
		visited[strconv.Itoa(x)+strconv.Itoa(y)]++
	}
	return len(visited)
}

func Vacuum2(input string) int {
	x := 0
	y := 0
	xr := 0
	yr := 0
	visited := make(map[string]int, 0)
	visited["00"] = 1
	for i, runeValue := range input {
		stringValue := string(runeValue)
		xy := instructions[stringValue]
		if i%2 == 0 {
			x += xy[0]
			y += xy[1]
			visited[strconv.Itoa(x)+strconv.Itoa(y)]++
		} else {
			xr += xy[0]
			yr += xy[1]
			visited[strconv.Itoa(xr)+strconv.Itoa(yr)]++
		}
	}
	return len(visited)
}
