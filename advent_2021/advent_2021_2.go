package advent_2021

import (
	"fmt"
	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_2_1() {
	input := utils.GetFile("input.txt")
	commands, directions := utils.SplitSliceStringToStringIntSlice(input)

	x := 0
	d := 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == "forward" {
			x += directions[i]
			continue
		}

		if commands[i] == "down" {
			d += directions[i]
			continue
		}

		if commands[i] == "up" {
			d -= directions[i]
			continue
		}
	}

	fmt.Println(x * d)
}

func Advent_2_2() {
	input := utils.GetFile("input.txt")
	commands, directions := utils.SplitSliceStringToStringIntSlice(input)

	x := 0
	d := 0
	aim := 0
	for i := 0; i < len(commands); i++ {
		if commands[i] == "forward" {
			x += directions[i]
			d += aim * directions[i]
			continue
		}

		if commands[i] == "down" {
			aim += directions[i]
			continue
		}

		if commands[i] == "up" {
			aim -= directions[i]
			continue
		}
	}

	fmt.Println(x * d)
}
