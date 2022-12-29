package advent_2022

import (
	"strconv"
)

var digits = map[string]bool{
	"0": true,
	"1": true,
	"2": true,
	"3": true,
	"4": true,
	"5": true,
	"6": true,
	"7": true,
	"8": true,
	"9": true,
}

type element struct {
	valInt   int
	valSlice []element
	isSlice  bool
	isFound  bool
}

func parseInt(startIndex int, r []rune) (int, element) {
	result := make([]rune, 0)
	i := startIndex
	for i < len(r) {
		if _, ok := digits[string(r[i])]; ok {
			result = append(result, r[i])
		} else {
			break
		}
		i++
	}

	intVal, _ := strconv.Atoi(string(result))
	return i - 1, element{valInt: intVal}
}

var pSum = 0

func parseSlice(startIndex int, r []rune) (int, element) {
	newEl := element{valSlice: make([]element, 0), isSlice: true}
	i := startIndex
	for i < len(r) {
		if string(r[i]) == "]" {
			break
		}

		if string(r[i]) == "[" {
			ind, el := parseSlice(i+1, r)
			newEl.valSlice = append(newEl.valSlice, el)

			i = ind
			i++

			continue
		}

		if string(r[i]) == "," {
			i++
			continue
		}

		ind, el := parseInt(i, r)
		i = ind
		newEl.valSlice = append(newEl.valSlice, el)

		i++
	}
	return i, newEl
}

func parse(s string) element {
	runes := []rune(s)

	var result element
	ind := 0
	for i := 0; i < len(runes); i++ {
		ind, result = parseSlice(i, runes)

		i = ind
	}

	return result
}

func compare(s1, s2 []element) int {
	i := 0
	result := 0
	for {
		if i >= len(s1) && i < len(s2) {
			return 1
		}
		if i < len(s1) && i >= len(s2) {
			return -1
		}
		if i >= len(s1) && i >= len(s2) {
			return result
		}
		if s1[i].isSlice == true && s2[i].isSlice == true {
			result = compare(s1[i].valSlice, s2[i].valSlice)
			if result == 0 {
				i++
				continue
			}

			return result
		}
		if s1[i].isSlice == false && s2[i].isSlice == false {
			if s1[i].valInt == s2[i].valInt {
				i++
				continue
			}
			if s1[i].valInt < s2[i].valInt {
				return 1
			} else {
				return -1
			}
		}
		if s1[i].isSlice == false && s2[i].isSlice == true {
			v := parse("[" + strconv.Itoa(s1[i].valInt) + "]")
			result = compare(v.valSlice[0].valSlice, s2[i].valSlice)
			if result == 0 {
				i++
				continue
			}
			return result
		}
		if s1[i].isSlice == true && s2[i].isSlice == false {
			v := parse("[" + strconv.Itoa(s2[i].valInt) + "]")
			result = compare(s1[i].valSlice, v.valSlice[0].valSlice)
			if result == 0 {
				i++
				continue
			}

			return result
		}
	}
}

func Distress1(input []string) int {
	var s1 element
	var s2 element

	i := 0
	j := 1
	sum := 0
	for i < len(input) {
		s1 = parse(input[i])
		s2 = parse(input[i+1])

		result := compare(s1.valSlice, s2.valSlice)
		if result > 0 {
			sum += j
		}
		j++
		i += 3
	}

	return sum
}

func Distress2(input []string) int {
	list := make([]element, 0, len(input))

	sum := 1
	for _, v := range input {
		if v == "" {
			continue
		}
		s1 := parse(v)
		list = append(list, s1)
	}

	sp1 := parse("[[2]]")
	sp1.isFound = true
	list = append(list, sp1)

	sp2 := parse("[[6]]")
	sp2.isFound = true
	list = append(list, sp2)

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-i-1; j++ {
			if compare(list[j].valSlice, list[j+1].valSlice) < 0 {
				t := list[j]
				list[j] = list[j+1]
				list[j+1] = t
			}
		}
	}

	for i, v := range list {
		if v.isFound {
			sum *= (i + 1)
		}
	}
	return sum
}
