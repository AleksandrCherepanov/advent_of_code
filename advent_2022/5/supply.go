package advent_2022

import (
	"fmt"
	"strconv"
	"strings"
)

func Supply1(steps []string, cargoes [][]string) string {
	result := ""

	for _, step := range steps {
		words := strings.Split(step, " ")
		amount, _ := strconv.ParseInt(words[1], 10, 0)
		from, _ := strconv.ParseInt(words[3], 10, 0)
		to, _ := strconv.ParseInt(words[5], 10, 0)

		var el string
		for i := 0; i < int(amount); i++ {
			cargoes[from-1], el = pop(cargoes[from-1])
			cargoes[to-1] = push(cargoes[to-1], el)
		}
	}

	for _, cargoe := range cargoes {
		result += cargoe[len(cargoe)-1]
	}

	return result
}

func pop(stack []string) ([]string, string) {
	element := stack[len(stack)-1]
	newStack := stack[:len(stack)-1]

	return newStack, element
}

func push(stack []string, element string) []string {
	stack = append(stack, element)
	return stack
}

func Supply2(steps []string, cargoes [][]string) string {
	result := ""

	for _, step := range steps {
		fmt.Println(cargoes)
		words := strings.Split(step, " ")
		amount, _ := strconv.ParseInt(words[1], 10, 0)
		from, _ := strconv.ParseInt(words[3], 10, 0)
		to, _ := strconv.ParseInt(words[5], 10, 0)

		var el string
		els := make([]string, amount)
		j := amount
		for i := 0; i < int(amount); i++ {
			cargoes[from-1], el = pop(cargoes[from-1])
			els[j-1] = el
			j--
		}

		cargoes[to-1] = append(cargoes[to-1], els...)
	}

	for _, cargoe := range cargoes {
		result += cargoe[len(cargoe)-1]
	}

	return result
}
