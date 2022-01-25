package main

import (
	"bufio"
	"fmt"
	"os"

	advent2015_2 "github.com/AleksandrCherepanov/advent_of_code/advent_2015/2"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0, 1000)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println(advent2015_2.Advent2015_2_1(input))
	fmt.Println(advent2015_2.Advent2015_2_2(input))
}
