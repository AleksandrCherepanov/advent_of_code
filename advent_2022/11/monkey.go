package advent_2022

import (
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	items map[string][]int
	op    *operation
}
type operation struct {
	div   int
	throw map[bool]string
}

func initState(input []string) map[string]*monkey {
	monkeys := make(map[string]*monkey, 0)
	key := ""
	var mky *monkey
	for _, s := range input {
		if strings.Contains(s, "Monkey") {
			mky = &monkey{items: make(map[string][]int)}
			key = getMonkeyKey(s)
			monkeys[key] = mky
			continue
		}

		if strings.Contains(s, "Starting") {
			items := make([]int, 0)
			s1 := strings.Trim(s, " ")
			numbers := strings.ReplaceAll(s1, "Starting items: ", "")
			for _, d := range strings.Split(numbers, ", ") {
				val, _ := strconv.Atoi(d)
				items = append(items, val)
			}

			monkeys[key].items[key] = items
		}

		if strings.Contains(s, "Test: divisible by ") {
			op := &operation{throw: make(map[bool]string)}
			mky.op = op
			s1 := strings.Trim(s, " ")
			number := strings.ReplaceAll(s1, "Test: divisible by ", "")
			mky.op.div, _ = strconv.Atoi(number)
		}

		if strings.Contains(s, "If true: throw to ") {
			s1 := strings.Trim(s, " ")
			m := strings.ReplaceAll(s1, "If true: throw to ", "")
			mky.op.throw[true] = m
		}

		if strings.Contains(s, "If false: throw to ") {
			s1 := strings.Trim(s, " ")
			m := strings.ReplaceAll(s1, "If false: throw to ", "")
			mky.op.throw[false] = m
		}
	}

	return monkeys
}

func exec(operand string, args ...int) int {
	if operand == "+" {
		return args[0] + args[1]
	}

	return args[0] * args[1]
}

func getMonkeyKey(s string) string {
	key := strings.ReplaceAll(s, ":", "")
	return strings.ToLower(key)
}

func Monkey1(input []string) int {
	state := initState(input)
	active := make(map[string]int, 0)
	for i := 0; i < 20; i++ {
		key := ""
		for _, s := range input {
			if strings.Contains(s, "Monkey") {
				key = getMonkeyKey(s)
			}

			if strings.Contains(s, "Operation") {
				for _, item := range state[key].items[key] {
					active[key]++
					s1 := strings.Trim(s, " ")
					expr := strings.ReplaceAll(s1, "Operation: new = ", "")
					parts := strings.Split(expr, " ")
					arg1 := 0
					arg2 := 0

					if parts[0] == "old" {
						arg1 = item
					} else {
						arg1, _ = strconv.Atoi(parts[0])
					}
					if parts[2] == "old" {
						arg2 = item
					} else {
						arg2, _ = strconv.Atoi(parts[2])
					}

					result := exec(parts[1], arg1, arg2)
					level := result / 3
					newMky := ""
					if level%state[key].op.div == 0 {
						newMky = state[key].op.throw[true]
					} else {
						newMky = state[key].op.throw[false]
					}

					state[newMky].items[newMky] = append(state[newMky].items[newMky], level)
				}
				state[key].items[key] = make([]int, 0)
			}
		}
	}

	max1 := 0
	max2 := 0

	for _, v := range active {
		if v > max1 {
			max2 = max1
			max1 = v
		}

		if v < max1 && v > max2 {
			max2 = v
		}
	}

	return max1 * max2
}

func printMky(state map[string]*monkey) {
	for k, v := range state {
		fmt.Println(k)
		fmt.Println(v.items)
		fmt.Println(v.op)
	}
}

func Monkey2(input []string) int {
	state := initState(input)
	active := make(map[string]int, 0)

	commonDiv := 1
	for _, m := range state {
		commonDiv *= m.op.div
	}

	for i := 0; i < 10000; i++ {
		key := ""
		for _, s := range input {
			if strings.Contains(s, "Monkey") {
				key = getMonkeyKey(s)
			}

			if strings.Contains(s, "Operation") {
				for _, item := range state[key].items[key] {
					active[key]++
					s1 := strings.Trim(s, " ")
					expr := strings.ReplaceAll(s1, "Operation: new = ", "")
					parts := strings.Split(expr, " ")
					arg1 := 0
					arg2 := 0

					if parts[0] == "old" {
						arg1 = item
					} else {
						arg1, _ = strconv.Atoi(parts[0])
					}
					if parts[2] == "old" {
						arg2 = item
					} else {
						arg2, _ = strconv.Atoi(parts[2])
					}

					result := exec(parts[1], arg1, arg2)
					level := result % commonDiv
					newMky := ""
					if level%state[key].op.div == 0 {
						newMky = state[key].op.throw[true]
					} else {
						newMky = state[key].op.throw[false]
					}

					state[newMky].items[newMky] = append(state[newMky].items[newMky], level)
				}
				state[key].items[key] = make([]int, 0)
			}
		}
	}

	max1 := 0
	max2 := 0

	for _, v := range active {
		if v > max1 {
			max2 = max1
			max1 = v
		}

		if v < max1 && v > max2 {
			max2 = v
		}
	}

	return max1 * max2
}
