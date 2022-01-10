package advent_2021

import (
	"fmt"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_6(days int) {
	input := utils.GetFile()
	timers := utils.ConvertSliceStringToInt(strings.Split(input[0], ","))

	maxTime := 8
	initialTime := 6
	capacity := maxTime + 1

	fishes := make([]int, capacity)
	for _, timer := range timers {
		fishes[timer]++
	}

	countZero := 0
	for i := 0; i < days; i++ {
		for j := 1; j < len(fishes); j++ {
			if fishes[j] == 0 {
				if j == 1 {
					fishes[j-1] = 0
				}
				continue
			}
			fishes[j-1] = fishes[j]
			fishes[j] = 0
		}

		if countZero != 0 {
			fishes[initialTime] += countZero
			fishes[maxTime] += countZero
			countZero = 0
		}

		countZero = fishes[0]
	}

	sum := utils.SumSliceInt(fishes)
	fmt.Println(sum)
}
