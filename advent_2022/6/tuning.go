package advent_2022

func Tuning1(input string) int {
	result := 0
	runes := []rune(input)
	dic := make(map[rune]int, 0)
	c := 0
	for i := 0; i < len(runes); i++ {
		if c == 4 {
			result = i
			break
		}

		if _, ok := dic[runes[i]]; ok {
			c = 0
			i = dic[runes[i]] + 1
			dic = make(map[rune]int, 0)
		}

		dic[runes[i]] = i
		c++
	}

	return result
}

func Tuning2(input string) int {
	result := 0
	runes := []rune(input)
	dic := make(map[rune]int, 0)
	c := 0
	for i := 0; i < len(runes); i++ {
		if c == 14 {
			result = i
			break
		}

		if _, ok := dic[runes[i]]; ok {
			c = 0
			i = dic[runes[i]] + 1
			dic = make(map[rune]int, 0)
		}

		dic[runes[i]] = i
		c++
	}

	return result
}
