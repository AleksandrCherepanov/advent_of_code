package advent_2022

import (
	"sort"
	"strconv"
)

func Calorie1(input []string) int {
	max := 0
	current := 0
	for _, v := range input {
		if v == "" {
			if max < current {
				max = current
			}
			current = 0
			continue
		}

		i, _ := strconv.Atoi(v)
		current += i
	}

	if max < current {
		max = current
		current = 0
	}

	return max
}

func Calorie2(input []string) int {
	elfs := make([]int, 0)
	current := 0
	for _, v := range input {
		if v == "" {
			elfs = append(elfs, current)
			current = 0
			continue
		}

		i, _ := strconv.Atoi(v)
		current += i
	}

	if current != 0 {
		elfs = append(elfs, current)
	}

	sort.Ints(elfs)

	sum := 0
	for i := len(elfs) - 1; i > len(elfs)-4; i-- {
		sum += elfs[i]
	}

	return sum
}
