package advent2015_1

func Advent2015_1_1(input string) int {
	floor := 0
	for _, inputRune := range input {
		inputChar := string(inputRune)
		if inputChar == "(" {
			floor++
		} else {
			floor--
		}
	}

	return floor
}

func Advent2015_1_2(input string) int {
	floor := 0
	for pos, inputRune := range input {
		inputChar := string(inputRune)
		if inputChar == "(" {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			return pos + 1
		}
	}

	return len([]rune(input))
}
