package advent_2022

import (
	"strconv"
	"strings"
)

func Camp1(input []string) int {
	sum := 0
	for _, r := range input {
		ranges := strings.Split(r, ",")
		minMax1 := strings.Split(ranges[0], "-")
		minMax2 := strings.Split(ranges[1], "-")

		l1, _ := strconv.ParseInt(minMax1[0], 10, 0)
		r1, _ := strconv.ParseInt(minMax1[1], 10, 0)
		l2, _ := strconv.ParseInt(minMax2[0], 10, 0)
		r2, _ := strconv.ParseInt(minMax2[1], 10, 0)

		if l1 <= l2 && r1 >= r2 {
			sum++
			continue
		}

		if l1 >= l2 && r1 <= r2 {
			sum++
			continue
		}
	}
	return sum
}

func Camp2(input []string) int {
	sum := 0
	for _, r := range input {
		ranges := strings.Split(r, ",")
		minMax1 := strings.Split(ranges[0], "-")
		minMax2 := strings.Split(ranges[1], "-")

		l1, _ := strconv.ParseInt(minMax1[0], 10, 0)
		r1, _ := strconv.ParseInt(minMax1[1], 10, 0)
		l2, _ := strconv.ParseInt(minMax2[0], 10, 0)
		r2, _ := strconv.ParseInt(minMax2[1], 10, 0)

		if l1 >= l2 && l1 <= r2 {
			sum++
			continue
		}

		if r1 <= r2 && r1 >= l2 {
			sum++
			continue
		}

		if l2 >= l1 && l2 <= r1 {
			sum++
			continue
		}

		if r2 <= r1 && r2 >= l1 {
			sum++
			continue
		}
	}
	return sum
}
