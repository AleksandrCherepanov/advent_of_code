package advent2015_2

import (
	"strconv"
	"strings"
)

func calcRibbon(sides []int) int {
	if len(sides) < 3 {
		return 0
	}

	result := 1
	for _, side := range sides {
		result *= side
	}

	return result
}

func calcSquareFeet(sizes string) int {
	sizesSlice := strings.Split(sizes, "x")
	if len(sizesSlice) < 3 {
		return 0
	}

	sizesInt := make([]int, len(sizesSlice))
	for i, sizeString := range sizesSlice {
		sizesInt[i], _ = strconv.Atoi(sizeString)
	}

	lw := sizesInt[0] * sizesInt[1]
	minSize := lw

	wh := sizesInt[1] * sizesInt[2]
	if wh < minSize {
		minSize = wh
	}

	hl := sizesInt[2] * sizesInt[0]
	if hl < minSize {
		minSize = hl
	}

	return (2*lw + 2*wh + 2*hl) + minSize
}

func calcFeetOfRibbon(sizes string) int {
	sizesSlice := strings.Split(sizes, "x")
	if len(sizesSlice) < 3 {
		return 0
	}

	sizesInt := make([]int, len(sizesSlice))
	for i, sizeString := range sizesSlice {
		sizesInt[i], _ = strconv.Atoi(sizeString)
	}

	lw := 2*sizesInt[0] + 2*sizesInt[1]
	minSize := lw

	wh := 2*sizesInt[1] + 2*sizesInt[2]
	if wh < minSize {
		minSize = wh
	}

	hl := 2*sizesInt[2] + 2*sizesInt[0]
	if hl < minSize {
		minSize = hl
	}

	return minSize + calcRibbon(sizesInt)
}

func Math1(input []string) int {
	sum := 0
	for _, sizes := range input {
		squareFeet := calcSquareFeet(sizes)
		sum += squareFeet
	}

	return sum
}

func Math2(input []string) int {
	sum := 0
	for _, sizes := range input {
		ribbon := calcFeetOfRibbon(sizes)
		sum += ribbon
	}

	return sum
}
