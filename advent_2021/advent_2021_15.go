package advent_2021

import (
	"github.com/AleksandrCherepanov/advent_of_code/field"
	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_15_1() {
	input := utils.GetFieldInt(utils.GetFile())
	f := field.NewFieldInt(input)
	f.Print()
}
