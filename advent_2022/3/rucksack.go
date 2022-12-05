package advent_2022

func Rucksack1(input []string) int {
	sum := 0
	for _, r := range input {
		runes := []rune(r)
		hash := make(map[rune]bool, len(r))
		middle := len(runes) / 2

		for i := 0; i < middle; i++ {
			hash[runes[i]] = true
		}

		for i := middle; i < len(runes); i++ {
			_, ok := hash[runes[i]]
			if ok {
				if runes[i] > 90 {
					sum += int(runes[i]) - 96
				} else {
					sum += int(runes[i]) - 38
				}
				break
			}
		}
	}
	return sum
}

func Rucksack2(input []string) int {
	sum := 0
	hash := make(map[rune]bool)
	hash2 := make(map[rune]bool)
	c := 0
	for _, r := range input {
		if c == 0 {
			for _, v := range r {
				hash[v] = true
			}
			c++
			continue
		}

		if c == 1 {
			for _, v := range r {
				_, ok := hash[v]
				if ok {
					hash2[v] = true
				}
			}
			c++
			continue
		}

		if c == 2 {
			for _, v := range r {
				_, ok := hash2[v]
				if ok {
					if v > 90 {
						sum += int(v) - 96
					} else {
						sum += int(v) - 38
					}
					break
				}
			}
			c = 0
			hash = make(map[rune]bool)
			hash2 = make(map[rune]bool)
		}
	}

	return sum
}
