package advent_2015_5

var unsupportedPair = map[string]bool{
	"ab": true,
	"cd": true,
	"pq": true,
	"xy": true,
}

var vowels = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

func Advent2015_5_1(input string) bool {
	countVowels := 0
	countLetter := 0
	prevLetter := ""

	if len(input) < 3 {
		return false
	}

	for _, letter := range input {
		sletter := string(letter)
		if _, ok := vowels[sletter]; ok {
			countVowels++
		}
		if prevLetter == "" {
			prevLetter = sletter
			continue
		}

		pair := prevLetter + sletter
		if _, ok := unsupportedPair[pair]; ok {
			return false
		}

		if prevLetter == sletter && countLetter == 0 {
			countLetter = 2
		}

		prevLetter = sletter
	}

	return countLetter == 2 && countVowels >= 3
}
