package main

import (
	"bufio"
	"fmt"
	"os"

	advent_2022 "github.com/AleksandrCherepanov/advent_of_code/advent_2022/5"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0, 3000)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	cargoes := [][]string{
		{"H", "B", "V", "W", "N", "M", "L", "P"},
		{"M", "Q", "H"},
		{"N", "D", "B", "G", "F", "Q", "M", "L"},
		{"Z", "T", "F", "Q", "M", "W", "G"},
		{"M", "T", "H", "P"},
		{"C", "B", "M", "J", "D", "H", "G", "T"},
		{"M", "N", "B", "F", "V", "R"},
		{"P", "L", "H", "M", "R", "G", "S"},
		{"P", "D", "B", "C", "N"},
	}

	fmt.Println(advent_2022.Supply1(input, cargoes))

	cargoes1 := [][]string{
		{"H", "B", "V", "W", "N", "M", "L", "P"},
		{"M", "Q", "H"},
		{"N", "D", "B", "G", "F", "Q", "M", "L"},
		{"Z", "T", "F", "Q", "M", "W", "G"},
		{"M", "T", "H", "P"},
		{"C", "B", "M", "J", "D", "H", "G", "T"},
		{"M", "N", "B", "F", "V", "R"},
		{"P", "L", "H", "M", "R", "G", "S"},
		{"P", "D", "B", "C", "N"},
	}

	fmt.Println(advent_2022.Supply2(input, cargoes1))
}
