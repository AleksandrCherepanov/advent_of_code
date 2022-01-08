package advent_2021

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_7_1() {
	input := utils.GetFile()
	positions := utils.ConvertSliceStringToInt(strings.Split(input[0], ","))

	sort.Ints(positions)
	maxPosition := positions[len(positions)-1]

	minFuel := -1
	for i := maxPosition; i >= 0; i-- {
		sum := 0
		for _, v := range positions {
			diff := v - i
			if diff < 0 {
				diff *= -1
			}
			sum += diff
		}
		if minFuel == -1 {
			minFuel = sum
		} else {
			if sum < minFuel {
				minFuel = sum
			}
		}
	}

	fmt.Println(minFuel)
}

func Advent_7_2() {
	input := utils.GetFile()
	positions := utils.ConvertSliceStringToInt(strings.Split(input[0], ","))

	sort.Ints(positions)
	maxPosition := positions[len(positions)-1]

	minFuel := -1
	for i := maxPosition; i >= 0; i-- {
		sum := 0
		for _, v := range positions {
			diff := v - i
			if diff < 0 {
				diff *= -1
			}

			if diff%2 == 0 {
				sum += (diff / 2) * (1 + diff)
			} else {
				sum += ((1 + diff) / 2) * diff
			}
		}
		if minFuel == -1 {
			minFuel = sum
		} else {
			if sum < minFuel {
				minFuel = sum
			}
		}
	}

	fmt.Println(minFuel)
}
