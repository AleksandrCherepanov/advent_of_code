package advent_2022

import (
	"strings"
)

var figures = map[string]map[string]int{
	"X": {"A": 3, "B": 0, "C": 6},
	"Y": {"A": 6, "B": 3, "C": 0},
	"Z": {"A": 0, "B": 6, "C": 3},
}

var rules = map[string]map[string]int{
	"X": {"A": 3, "B": 1, "C": 2},
	"Y": {"A": 4, "B": 5, "C": 6},
	"Z": {"A": 8, "B": 9, "C": 7},
}

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func Rps1(input []string) int {
	sum := 0
	for _, s := range input {
		letters := strings.Split(s, " ")
		sum += scores[letters[1]] + figures[letters[1]][letters[0]]
	}

	return sum
}

func Rps2(input []string) int {
	sum := 0
	for _, s := range input {
		letters := strings.Split(s, " ")
		sum += rules[letters[1]][letters[0]]
	}

	return sum
}
