package main

import (
	"bufio"
	"fmt"
	"os"

	advent2015_1 "github.com/AleksandrCherepanov/advent_of_code/advent_2015"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		break
	}

	fmt.Println(advent2015_1.Advent2015_1_1(input))
	fmt.Println(advent2015_1.Advent2015_1_2(input))
}
