package main

import (
	"bufio"
	"fmt"
	"os"

	advent_2022 "github.com/AleksandrCherepanov/advent_of_code/advent_2022/2"
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

	fmt.Println(advent_2022.Rps1(input))
	fmt.Println(advent_2022.Rps2(input))
}
