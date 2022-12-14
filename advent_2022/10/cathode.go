package advent_2022

import (
	"strconv"
	"strings"
)

var steps = map[int]int{
	20:  20,
	60:  60,
	100: 100,
	140: 140,
	180: 180,
	220: 220,
}

func addX(arg int, cycle int, x int, powers []int, add bool) (int, int, []int) {
	cycle++
	if add {
		x += arg
	}
	if v, ok := steps[cycle]; ok {
		powers = append(powers, v*x)
	}
	return cycle, x, powers
}

func noop(cycle int, x int, powers []int) (int, []int) {
	cycle++
	if v, ok := steps[cycle]; ok {
		powers = append(powers, v*x)
	}
	return cycle, powers
}

func Cathode1(input []string) int {
	x := 1
	cycle := 1
	powers := make([]int, 0, len(steps))
	for _, s := range input {
		parts := strings.Split(s, " ")
		cmd := parts[0]

		if cmd == "noop" {
			cycle, powers = noop(cycle, x, powers)
		}

		if cmd == "addx" {
			val, _ := strconv.Atoi(parts[1])
			cycle, x, powers = addX(val, cycle, x, powers, false)
			cycle, x, powers = addX(val, cycle, x, powers, true)
		}
	}

	sum := 0
	for _, v := range powers {
		sum += v
	}
	return sum
}

func pixels(template []int, idx, val int) []int {
	if idx-1 >= 0 && idx-1 < len(template) {
		template[idx-1] = val
	}
	if idx >= 0 && idx < len(template) {
		template[idx] = val
	}
	if idx+1 >= 0 && idx+1 < len(template) {
		template[idx+1] = val
	}

	return template
}

func Cathode2(input []string) string {
	template := []int{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	x := 1
	cycle := 0
	powers := make([]int, 0, len(steps))
	screen := make([]int, len(template)*6+1)
	screen[0] = 1
	for _, s := range input {
		parts := strings.Split(s, " ")
		cmd := parts[0]

		if cmd == "noop" {
			cycle, powers = noop(cycle, x, powers)
			index := cycle % 40
			screen[cycle] = template[index]
		}

		if cmd == "addx" {
			val, _ := strconv.Atoi(parts[1])
			prevX := x
			cycle, x, powers = addX(val, cycle, x, powers, false)
			template = pixels(template, prevX, 0)
			template = pixels(template, x, 1)
			index := cycle % 40
			screen[cycle] = template[index]

			cycle, x, powers = addX(val, cycle, x, powers, true)
			template = pixels(template, prevX, 0)
			template = pixels(template, x, 1)

			index = cycle % 40
			screen[cycle] = template[index]
		}
	}

	s := ""
	for i := 0; i < len(screen)-1; i++ {
		//if i%40 == 0 {
		//	s += "\n"
		//}
		if screen[i] == 1 {
			s += "#"
		} else {
			s += "."
		}
	}

	return s
}
