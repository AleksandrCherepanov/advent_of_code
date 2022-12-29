package advent_2015

func Sticks1(input []string) int {
	code := 0
	char := 0
	for _, s := range input {
		code += len(s)
		c := 0
		if len(s) == 2 {
			continue
		}

		runes := []rune(s)
		l := len(runes) - 1
		for i := 1; i < l; i++ {
			if string(runes[i]) == `\` {
				if i+1 < l && string(runes[i+1]) == `"` {
					c++
					i++
					continue
				}

				if i+1 < l && string(runes[i+1]) == `\` {
					c++
					i++
					continue
				}

				if i+1 < l && string(runes[i+1]) == `x` {
					c++
					i = i + 3
					continue
				}
			}

			c++
		}

		char += c
	}

	return code - char
}

func Sticks2(input []string) int {
	code := 0
	char := 0
	for _, s := range input {
		code += len(s)
		c := len(s) + 4
		runes := []rune(s)
		l := len(runes) - 1
		for i := 1; i < l; i++ {
			if string(runes[i]) == `"` {
				c++
				continue
			}

			if string(runes[i]) == `\` {
				c++
				continue
			}
		}

		char += c
	}

	return char - code
}
