package advent_2022

import (
	"strconv"
	"strings"
)

func isSeen(field [][]int, r int, c int) bool {
	el := field[r][c]
	direction := 0
	for row := r - 1; row >= 0; row-- {
		if field[row][c] >= el {
			direction++
			break
		}
	}

	for row := r + 1; row < len(field); row++ {
		if field[row][c] >= el {
			direction++
			break
		}
	}

	for col := c - 1; col >= 0; col-- {
		if field[r][col] >= el {
			direction++
			break
		}
	}

	for col := c + 1; col < len(field[0]); col++ {
		if field[r][col] >= el {
			direction++
			break
		}
	}

	return direction < 4
}

func score(field [][]int, r int, c int) int {
	el := field[r][c]

	d1 := 0
	for row := r - 1; row >= 0; row-- {
		d1++
		if field[row][c] >= el {
			break
		}
	}

	d2 := 0
	for row := r + 1; row < len(field); row++ {
		d2++
		if field[row][c] >= el {
			break
		}
	}

	d3 := 0
	for col := c - 1; col >= 0; col-- {
		d3++
		if field[r][col] >= el {
			break
		}
	}

	d4 := 0
	for col := c + 1; col < len(field[0]); col++ {
		d4++
		if field[r][col] >= el {
			break
		}
	}

	return d1 * d2 * d3 * d4
}

func stringToMatrix(input []string) [][]int {
	matrix := make([][]int, 0, len(input))

	for _, s := range input {
		row := strings.Split(s, "")

		rowInt := make([]int, 0, len(row))
		for _, r := range row {
			val, _ := strconv.Atoi(r)
			rowInt = append(rowInt, val)
		}

		matrix = append(matrix, rowInt)
	}

	return matrix
}

func Treetop1(input []string) int {
	field := stringToMatrix(input)
	sum := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if isSeen(field, i, j) {
				sum++
			}
		}
	}

	return sum
}

func Treetop2(input []string) int {
	field := stringToMatrix(input)
	max := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			s := score(field, i, j)
			if s > max {
				max = s
			}
		}
	}

	return max
}
