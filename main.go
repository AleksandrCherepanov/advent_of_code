package main

import (
	"bufio"
	"fmt"
	"os"

	advent "github.com/AleksandrCherepanov/advent_of_code/advent_2015/8"
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

	fmt.Println(advent.Sticks1(input))
	fmt.Println(advent.Sticks2(input))
}
