package advent_2021

import (
	"fmt"
	"strconv"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func parse(packages string, sum *int) (string, int) {
	version := packages[0:3]
	dv, _ := strconv.ParseInt(version, 2, 0)
	*(sum) += int(dv)
	ptype := packages[3:6]
	newPackages := packages[6:]
	numbers := make([]int, 0)
	if ptype == "100" {
		needBreak := false
		number := ""
		for len([]rune(newPackages)) >= 5 && !needBreak {
			if newPackages[:1] == "0" {
				needBreak = true
			}

			number += string(newPackages[1:5])
			newPackages = newPackages[5:]
		}
		num, _ := strconv.ParseInt(number, 2, 0)
		return newPackages, int(num)
	} else {
		if newPackages[:1] == "1" {
			length, _ := strconv.ParseInt(newPackages[1:12], 2, 0)
			newPackages = newPackages[12:]
			num := 0
			for i := 0; i < int(length); i++ {
				newPackages, num = parse(newPackages, sum)
				numbers = append(numbers, num)
			}
		} else {
			length, _ := strconv.ParseInt(newPackages[1:16], 2, 0)
			newPackages = newPackages[16:]
			plength := len(newPackages)
			value, _ := strconv.ParseInt(newPackages, 2, 0)
			num := 0
			for int(length) >= 5 && value != 0 {
				newPackages, num = parse(newPackages, sum)
				length = int64(int(length) - (plength - len(newPackages)))
				plength = len(newPackages)
				value, _ = strconv.ParseInt(newPackages, 2, 0)
				numbers = append(numbers, num)
			}
		}
	}

	result := 0
	switch ptype {
	case "000":
		for _, v := range numbers {
			result += v
		}
	case "001":
		result = 1
		for _, v := range numbers {
			result *= v
		}
	case "010":
		result = numbers[0]
		for i := 1; i < len(numbers); i++ {
			if numbers[i] < result {
				result = numbers[i]
			}
		}
	case "011":
		result = numbers[0]
		for i := 1; i < len(numbers); i++ {
			if numbers[i] > result {
				result = numbers[i]
			}
		}
	case "101":
		if numbers[0] > numbers[1] {
			result = 1
		}
	case "110":
		if numbers[0] < numbers[1] {
			result = 1
		}
	case "111":
		if numbers[0] == numbers[1] {
			result = 1
		}
	}

	return newPackages, result
}

var hexToBinTable = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func Advent_16_1() {
	input := utils.GetFile("input.txt")[0]

	binary := ""
	for _, ch := range input {
		if bin, ok := hexToBinTable[string(ch)]; ok {
			binary += bin
		}
	}

	sum := 0
	_, r := parse(binary, &sum)
	fmt.Println(sum, r)
}
