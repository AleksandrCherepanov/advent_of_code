package advent_2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_13_1() {
	input := utils.GetFile("input.txt")

	xSlice, ySlice, foldSlice, xyMap := parseInput(input)

	fold := foldSlice[0]

	if fold < 0 {
		ySlice = calculateNewPositions(ySlice, fold*-1)

	} else {
		xSlice = calculateNewPositions(xSlice, fold)
	}

	diff := 0
	for i := 0; i < len(xSlice); i++ {
		key := strconv.Itoa(xSlice[i]) + "," + strconv.Itoa(ySlice[i])
		xyMap[key]++
		if xyMap[key] > 1 {
			diff++
		}
	}

	fmt.Println(len(xSlice) - diff)

}

func Advent_13_2() {
	input := utils.GetFile("input.txt")

	xSlice, ySlice, foldSlice, _ := parseInput(input)

	for _, fold := range foldSlice {
		if fold < 0 {
			ySlice = calculateNewPositions(ySlice, fold*-1)
		} else {
			xSlice = calculateNewPositions(xSlice, fold)
		}
	}

	xMax := xSlice[0]
	yMax := ySlice[0]
	field := make(map[string]bool, 0)
	for i := 0; i < len(xSlice); i++ {
		if xSlice[i] > xMax {
			xMax = xSlice[i]
		}
		if ySlice[i] > yMax {
			yMax = ySlice[i]
		}
		key := strconv.Itoa(ySlice[i]) + "," + strconv.Itoa(xSlice[i])
		field[key] = true
	}

	for i := 0; i <= yMax; i++ {
		for j := 0; j <= xMax; j++ {
			key := strconv.Itoa(i) + "," + strconv.Itoa(j)
			if _, ok := field[key]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseInput(input []string) ([]int, []int, []int, map[string]int) {
	xSlice := make([]int, 0, len(input))
	ySlice := make([]int, 0, len(input))

	// x > 0, y < 0
	foldSlice := make([]int, 0)
	xyMap := make(map[string]int, 0)

	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.Contains(line, "fold") {
			var fold int
			if strings.Contains(line, "x=") {
				fold, _ = strconv.Atoi(strings.Split(line, "=")[1])
			} else {
				fold, _ = strconv.Atoi(strings.Split(line, "=")[1])
				fold *= -1
			}

			foldSlice = append(foldSlice, fold)
			continue
		}
		pair := strings.Split(line, ",")
		xyMap[line] = 0
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])

		xSlice = append(xSlice, x)
		ySlice = append(ySlice, y)
	}

	return xSlice, ySlice, foldSlice, xyMap
}

func calculateNewPositions(curentPositions []int, foldingLine int) []int {
	for i := 0; i < len(curentPositions); i++ {
		if curentPositions[i] > foldingLine {
			newY := curentPositions[i] - foldingLine
			curentPositions[i] = foldingLine - newY
		}
	}

	return curentPositions
}
