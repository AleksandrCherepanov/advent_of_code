package advent_2021

import (
	"fmt"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

var counter = map[string]int{
	"A": 0,
	"B": 0,
	"C": 0,
	"D": 0,
	"E": 0,
	"F": 0,
	"G": 0,
	"H": 0,
	"I": 0,
	"J": 0,
	"K": 0,
	"L": 0,
	"M": 0,
	"N": 0,
	"O": 0,
	"P": 0,
	"Q": 0,
	"R": 0,
	"S": 0,
	"T": 0,
	"U": 0,
	"V": 0,
	"W": 0,
	"X": 0,
	"Y": 0,
	"Z": 0,
}

func split(step int, pairs map[string]int, counter map[string]int, template map[string]string) map[string]int {
	i := 0
	for i < step {
		newPairs := make(map[string]int, 0)
		for pair, count := range pairs {
			element := template[pair]
			newPairs[pair[:1]+element] += count
			newPairs[element+pair[1:]] += count
			counter[element] += count
		}
		pairs = newPairs
		i++
	}

	return counter
}

func Advent_14_1() {
	input := utils.GetFile("input.txt")

	line := input[0]

	template := make(map[string]string, 0)
	for i := 2; i < len(input); i++ {
		pair := strings.Split(input[i], " -> ")
		template[pair[0]] = pair[1]
	}

	pairs := make(map[string]int)
	pair := ""
	for _, c := range line {
		counter[string(c)]++
		pair += string(c)
		if len([]rune(pair)) == 2 {
			pairs[pair]++
			pair = pair[1:]
		}
	}

	counter = split(40, pairs, counter, template)
	fmt.Println(counter)

	max := 0
	for _, count := range counter {
		if count > max {
			max = count
		}
	}

	min := max
	for _, count := range counter {
		if count != 0 && count < min {
			min = count
		}
	}

	fmt.Println(max - min)
}
