package advent_2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_5_1() {
	input := utils.GetFile()
	field := make(map[string]int, 0)
	for _, line := range input {
		coord := strings.Split(strings.ReplaceAll(line, " -> ", ","), ",")
		points := utils.ConvertSliceStringToInt(coord)
		if points[0] != points[2] && points[1] != points[3] {
			continue
		}
		if points[0] == points[2] && points[1] == points[3] {
			key := getKey(points[0], points[1])
			field[key]++
		}

		if points[1] == points[3] {
			writeLineToField(field, points[0], points[2], points[1], true)
		}

		if points[0] == points[2] {
			writeLineToField(field, points[1], points[3], points[0], false)
		}
	}

	result := 0
	for _, v := range field {
		if v > 1 {
			result++
		}
	}

	fmt.Println(result)
}

func getKey(x, y int) string {
	xStr := strconv.Itoa(x)
	yStr := strconv.Itoa(y)

	return xStr + ":" + yStr
}

func writeLineToField(field map[string]int, p1, p2, p3 int, x bool) {
	start := p2
	end := p1
	if p1 < p2 {
		start = p1
		end = p2
	}

	for i := start; i <= end; i++ {
		key := getKey(i, p3)
		if !x {
			key = getKey(p3, i)
		}
		field[key]++
	}
}

func writeDiagonalLineToField(field map[string]int, x1, x2, y1, y2 int) {
	if x1 > x2 && y1 < y2 {
		for x1 >= x2 && y1 <= y2 {
			key := getKey(x1, y1)
			field[key]++
			x1--
			y1++
		}
		return
	}

	if x1 < x2 && y1 < y2 {
		for x1 <= x2 && y1 <= y2 {
			key := getKey(x1, y1)
			field[key]++
			x1++
			y1++
		}
		return
	}

	if x1 < x2 && y1 > y2 {
		for x1 <= x2 && y1 >= y2 {
			key := getKey(x1, y1)
			field[key]++
			x1++
			y1--
		}
		return
	}

	if x1 > x2 && y1 > y2 {
		for x1 >= x2 && y1 >= y2 {
			key := getKey(x1, y1)
			field[key]++
			x1--
			y1--
		}
		return
	}
}

func Advent_5_2() {
	input := utils.GetFile()
	field := make(map[string]int, 0)
	for _, line := range input {
		coord := strings.Split(strings.ReplaceAll(line, " -> ", ","), ",")
		points := utils.ConvertSliceStringToInt(coord)
		if points[0] != points[2] && points[1] != points[3] {
			writeDiagonalLineToField(field, points[0], points[2], points[1], points[3])
			continue
		}

		if points[0] == points[2] && points[1] == points[3] {
			key := getKey(points[0], points[1])
			field[key]++
		}

		if points[1] == points[3] {
			writeLineToField(field, points[0], points[2], points[1], true)
		}

		if points[0] == points[2] {
			writeLineToField(field, points[1], points[3], points[0], false)
		}

	}

	result := 0
	for _, v := range field {
		if v > 1 {
			result++
		}
	}

	fmt.Println(result)
}
