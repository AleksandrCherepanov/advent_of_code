package advent_2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_17_1() {
	input := utils.GetFile("input.txt")[0]
	input = input[15:]

	xyRaw := strings.Split(input, ",")
	yRange := strings.Split(xyRaw[1][3:], "..")

	yMin, _ := strconv.Atoi(yRange[0])

	vy := yMin
	if vy < 0 {
		vy *= -1
	}

	fmt.Println(vy * (vy - 1) / 2)
}

func Advent_17_2() {
	input := utils.GetFile("input.txt")[0]
	input = input[15:]

	xyRaw := strings.Split(input, ",")
	xRange := strings.Split(xyRaw[0], "..")
	yRange := strings.Split(xyRaw[1][3:], "..")

	xMin, _ := strconv.Atoi(xRange[0])
	xMax, _ := strconv.Atoi(xRange[1])

	yMin, _ := strconv.Atoi(yRange[0])
	yMax, _ := strconv.Atoi(yRange[1])

	dx := xMax - xMin
	if dx < 0 {
		dx *= -1
	}
	dx++

	dy := yMin - yMax
	if dy < 0 {
		dy *= -1
	}
	dy++

	fmt.Println(dx * dy * 2)
}
