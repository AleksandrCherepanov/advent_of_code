package advent_2021

import (
	"fmt"
	"sort"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_10() {
	input := utils.GetFile()

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	closeBraces := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}

	openBraces := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	wrongBraces := make([]string, 0)
	unfinished := make([][]string, 0)

	for _, line := range input {
		stack := make([]string, 0)
		wrong := false
		for _, ch := range line {
			str := string(ch)
			if _, ok := openBraces[str]; ok {
				stack = append(stack, str)
				continue
			}

			if brace, ok := closeBraces[str]; ok {
				last := stack[len(stack)-1]
				if last == brace {
					stack = stack[:len(stack)-1]
				} else {
					wrongBraces = append(wrongBraces, str)
					wrong = true
					break
				}
			}
		}
		if !wrong && len(stack) > 0 {
			fmt.Println(stack)
			unfinished = append(unfinished, stack)
		}
	}

	sum1 := findCorruptedPoints(wrongBraces, points)
	fmt.Println(sum1)

	sum2 := findUnfinishedPoints(unfinished, openBraces)
	fmt.Println(sum2)
}

func findCorruptedPoints(slice []string, points map[string]int) int {
	sum := 0
	for _, brace := range slice {
		str := brace
		sum += points[str]
	}

	return sum
}

func findUnfinishedPoints(slice [][]string, points map[string]int) uint64 {
	sums := make([]uint64, 0, len(slice))
	for _, line := range slice {
		sum := uint64(0)
		for i := len(line) - 1; i >= 0; i-- {
			sum = sum * uint64(5)
			sum += uint64(points[line[i]])
		}
		sums = append(sums, sum)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] < sums[j]
	})
	fmt.Println(sums)
	return sums[len(sums)/2]
}
