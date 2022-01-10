package advent_2021

import (
	"fmt"
	"strconv"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_3_1() {
	input := utils.GetFile()
	rate := make([]int, len(input[0]))

	for i := 0; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		for j := 0; j < len(input[i]); j++ {
			if string(input[i][j]) == "1" {
				rate[j]++
			}
		}
	}

	gammaRate := ""
	for i := 0; i < len(rate); i++ {
		if len(input)-rate[i] < rate[i] {
			gammaRate += "1"
		} else {
			gammaRate += "0"
		}
	}

	gammaRateInt, err := strconv.ParseInt(gammaRate, 2, 0)
	if err != nil {
		fmt.Println(err)
	}

	epsilonRateInt := gammaRateInt ^ 31

	fmt.Println(gammaRateInt * epsilonRateInt)
}

func Advent_3_2() {
	input := utils.GetFile()
	input = input[:len(input)-1]

	oxgRate := calculateRate(input, mostBit)
	co2Rate := calculateRate(input, leastBit)

	oxg, err := strconv.ParseInt(oxgRate, 2, 0)
	if err != nil {
		fmt.Println(err)
	}

	co2, err := strconv.ParseInt(co2Rate, 2, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(oxg * co2)
}

func calculateRate(input []string, bitCalculator func(int) string) string {
	rate := append(make([]string, 0, len(input)), input...)
	pos := 0

	for len(rate) != 1 {
		bit := calculateBit(rate, pos, bitCalculator)
		tmp := make([]string, 0, len(rate))
		for _, v := range rate {
			if string(v[pos]) == bit {
				tmp = append(tmp, v)
			}
		}
		rate = tmp[:]
		pos++
	}

	return rate[0]
}

func calculateBit(input []string, pos int, determineBitFunc func(int) string) string {
	sum := 0
	for _, v := range input {
		if string(v[pos]) == "1" {
			sum++
		} else {
			sum--
		}
	}

	return determineBitFunc(sum)
}

func mostBit(sum int) string {
	if sum >= 0 {
		return "1"
	}

	return "0"

}

func leastBit(sum int) string {
	if sum >= 0 {
		return "0"
	}

	return "1"
}
