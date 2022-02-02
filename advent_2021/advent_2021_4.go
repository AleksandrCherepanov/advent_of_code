package advent_2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_4_1() {
	input := utils.GetFile("input.txt")
	numberString := strings.Split(strings.Trim(input[0], " "), ",")
	numbers := utils.ConvertSliceStringToInt(numberString)

	boardSize := 5
	boards := make([][]int, 0, (len(input)-2)/5)
	line := make([]int, 0, boardSize*boardSize)
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			boards = append(boards, line)
			line = make([]int, 0, boardSize*boardSize)
			continue
		}
		line = append(line, createLineOfBoard(input[i])...)
	}
	boards = append(boards, line)

	winColumn := make(map[string]int, 0)
	winLine := make(map[string]int, 0)

	columnFound := false
	lineFound := false
	winBoard := -1
	lastNumber := -1
	for _, number := range numbers {
		for boardIndex, board := range boards {
			for elementIndex, element := range board {
				if element == number {
					boards[boardIndex][elementIndex] = -1
					boardKey := strconv.Itoa(boardIndex)
					columnKey := boardKey + ":" + strconv.Itoa(elementIndex%5)
					lineKey := boardKey + ":" + strconv.Itoa(elementIndex/5)
					winColumn[columnKey]++
					if winColumn[columnKey] == 5 {
						winBoard = boardIndex
						lastNumber = number
						columnFound = true
						break
					}
					winLine[lineKey]++
					if winLine[lineKey] == 5 {
						winBoard = boardIndex
						lastNumber = number
						lineFound = true
						break
					}
				}
			}
			if lineFound || columnFound {
				break
			}
		}
		if lineFound || columnFound {
			break
		}
	}

	sum := 0
	for _, element := range boards[winBoard] {
		if element != -1 {
			sum += element
		}
	}

	fmt.Println(sum * lastNumber)
}

func Advent_4_2() {
	input := utils.GetFile("input.txt")
	numberString := strings.Split(strings.Trim(input[0], " "), ",")
	numbers := utils.ConvertSliceStringToInt(numberString)

	boardSize := 5
	boards := make([][]int, 0, (len(input)-2)/5)
	line := make([]int, 0, boardSize*boardSize)
	for i := 2; i < len(input); i++ {
		if input[i] == "" {
			boards = append(boards, line)
			line = make([]int, 0, boardSize*boardSize)
			continue
		}
		line = append(line, createLineOfBoard(input[i])...)
	}
	boards = append(boards, line)

	winColumn := make(map[string]int, 0)
	winLine := make(map[string]int, 0)

	columnFound := false
	lineFound := false
	winBoardMap := make(map[int]bool, 0)
	winBoard := make([]int, 0, len(boards))
	lastNumber := -1
	for _, number := range numbers {
		for boardIndex, board := range boards {
			if _, ok := winBoardMap[boardIndex]; ok {
				continue
			}
			for elementIndex, element := range board {
				if element == number {
					boards[boardIndex][elementIndex] = -1
					boardKey := strconv.Itoa(boardIndex)
					columnKey := boardKey + ":" + strconv.Itoa(elementIndex%5)
					lineKey := boardKey + ":" + strconv.Itoa(elementIndex/5)
					winColumn[columnKey]++
					if winColumn[columnKey] == 5 {
						lastNumber = number
						columnFound = true
						break
					}
					winLine[lineKey]++
					if winLine[lineKey] == 5 {
						lastNumber = number
						lineFound = true
						break
					}
				}
			}
			if lineFound || columnFound {
				lineFound = false
				columnFound = false
				winBoardMap[boardIndex] = true
				winBoard = append(winBoard, boardIndex)
			}
		}
	}

	sum := 0
	for _, element := range boards[winBoard[len(winBoard)-1]] {
		if element != -1 {
			sum += element
		}
	}

	fmt.Println(sum * lastNumber)
}

func createLineOfBoard(input string) []int {
	stringLine := strings.ReplaceAll(input, "  ", " ")
	stringLine = strings.Trim(stringLine, " ")
	sliceString := strings.Split(stringLine, " ")
	return utils.ConvertSliceStringToInt(sliceString)
}
