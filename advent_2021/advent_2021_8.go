package advent_2021

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_8_1() {
	input := utils.GetFile()

	numbers := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}

	count := 0
	for _, line := range input {
		signal := 0
		runes := []rune(line)
		for i := len(runes) - 1; i >= 0; i-- {
			if string(runes[i]) == "|" {
				break
			}

			if string(runes[i]) == " " {
				if _, ok := numbers[signal]; ok {
					count++
				}
				signal = 0
				continue
			}
			signal++
		}
	}

	fmt.Println(count)
}

func Advent_8_2() {
	input := utils.GetFile()

	sum := 0
	for _, line := range input {
		displaysNumbers := strings.Split(line, " | ")
		displays := displaysNumbers[0]
		numbers := displaysNumbers[1]

		signals := defineSignals(displays)
		result := defineNumbers(numbers, signals)
		s, _ := strconv.Atoi(result)
		sum += s
	}

	fmt.Println(sum)
}

func defineSignals(input string) []string {
	tokens := strings.Split(input, " ")

	sort.Slice(tokens, func(i, j int) bool {
		return len([]rune(tokens[i])) < len([]rune(tokens[j]))
	})

	/*
			0
		3		1
			2
		6		4
			5
	*/
	signals := make([]string, 7)

	//0, 5
	source := tokens[9]
	for i := len(tokens) - 2; i > 2; i-- {
		source = stringIntersection(source, tokens[i])
	}

	signals[0] = stringIntersection(source, tokens[1])
	signals[5] = stringDiff(signals[0], source)

	//1, 4
	for i := 6; i <= 8; i++ {
		result := stringDiff(tokens[i], tokens[0])
		if result != "" {
			signals[1] = result
			signals[4] = stringDiff(result, tokens[0])
			break
		}
	}
	//2
	source = strings.Join(signals, "")
	for i := 3; i <= 5; i++ {
		result := stringIntersection(tokens[i], source)
		if len(result) == len(source) {
			signals[2] = stringDiff(source, tokens[i])
			break
		}
	}
	//3, 6
	source = strings.Join(signals, "")
	for i := 6; i <= 8; i++ {
		result := stringDiff(source, tokens[i])
		if len([]rune(result)) == 1 {
			signals[3] = result
			signals[6] = stringDiff(strings.Join(signals, ""), tokens[9])
			break
		}
	}

	result := make([]string, 10)
	for _, i := range []int{0, 1, 3, 4, 5, 6} {
		result[0] += signals[i]
	}

	for _, i := range []int{0, 1, 2, 5, 6} {
		result[2] += signals[i]
	}

	for _, i := range []int{0, 1, 2, 4, 5} {
		result[3] += signals[i]
	}

	for _, i := range []int{0, 2, 3, 4, 5} {
		result[5] += signals[i]
	}

	for _, i := range []int{0, 2, 3, 4, 5, 6} {
		result[6] += signals[i]
	}

	for _, i := range []int{0, 1, 2, 3, 4, 5} {
		result[9] += signals[i]
	}

	return result
}

func defineNumbers(input string, signals []string) string {
	numbers := strings.Split(input, " ")
	result := ""
	for _, number := range numbers {
		length := len([]rune(number))
		if length == 2 {
			result += "1"
			continue
		}

		if length == 3 {
			result += "7"
			continue
		}

		if length == 4 {
			result += "4"
			continue
		}

		if length == 7 {
			result += "8"
			continue
		}

		if length == 6 {
			intersection := stringIntersection(number, signals[9])
			if len(intersection) == len(number) {
				result += "9"
				continue
			}

			intersection = stringIntersection(number, signals[6])
			if len(intersection) == len(number) {
				result += "6"
				continue
			}

			intersection = stringIntersection(number, signals[0])
			if len(intersection) == len(number) {
				result += "0"
				continue
			}
		}

		if length == 5 {
			intersection := stringIntersection(number, signals[5])
			if len(intersection) == len(number) {
				result += "5"
				continue
			}

			intersection = stringIntersection(number, signals[3])
			if len(intersection) == len(number) {
				result += "3"
				continue
			}

			intersection = stringIntersection(number, signals[2])
			if len(intersection) == len(number) {
				result += "2"
				continue
			}
		}
	}

	return result
}

func stringDiff(source, target string) string {
	dict := make(map[string]bool)
	for _, s := range source {
		dict[string(s)] = true
	}

	result := ""
	for _, ch := range target {
		if _, ok := dict[string(ch)]; !ok {
			result += string(ch)
		}
	}

	return result
}

func stringIntersection(source, target string) string {
	dict := make(map[string]bool)
	for _, s := range source {
		dict[string(s)] = true
	}

	result := ""
	for _, ch := range target {
		if _, ok := dict[string(ch)]; ok {
			result += string(ch)
		}
	}

	return result
}
