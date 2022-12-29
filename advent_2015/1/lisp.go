package advent2015_1

func Lisp1(input string) int {
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

func Lisp2(input string) int {
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
