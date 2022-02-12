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

	yMinAbs := yMin * -1
	count := 0
	for y := yMin; y < yMinAbs; y++ {
		for x := 1; x <= xMax; x++ {
			vx := x
			vy := y
			xp := 0
			yp := 0
			for t := 0; t < 2*yMinAbs+1; t++ {
				xp += vx
				yp += vy

				vx -= 1
				if vx < 0 {
					vx = 0
				}

				vy -= 1

				if xMin <= xp && xp <= xMax && yp <= yMax && yp >= yMin {
					count++
					break
				}

				if xp > xMax && yp < yMin {
					break
				}

			}
		}
	}

	fmt.Println(count)
}
