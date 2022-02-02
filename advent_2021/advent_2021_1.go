package advent_2021

import (
	"fmt"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_1_1() {
	input := utils.GetFile("input.txt")
	intSlice := utils.ConvertSliceStringToInt(input)

	prevValue := intSlice[0]
	result := 0
	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] > prevValue {
			result++
		}
		prevValue = intSlice[i]
	}

	fmt.Println(result)
}

func Advent_1_2() {
	input := utils.GetFile("input.txt")
	intSlice := utils.ConvertSliceStringToInt(input)

	prevValue := intSlice[0] + intSlice[1] + intSlice[2]
	result := 0
	for i := 1; i < len(intSlice)-2; i++ {
		windowValue := intSlice[i] + intSlice[i+1] + intSlice[i+2]
		if windowValue > prevValue {
			result++
		}
		prevValue = windowValue
	}

	fmt.Println(result)
}
