package advent_2015

import (
	"strings"
)

func addIfVowels(r rune) int {
	if strings.ContainsRune("aeiou", r) {
		return 1
	}

	return 0
}

func addIfDoubles(r1, r2 rune) int {
	if r1 == r2 {
		return 1
	}

	return 0
}

func isForbidden(w string) bool {
	return w == "ab" || w == "cd" || w == "pq" || w == "xy"
}

func Intern1(input []string) int {
	result := 0
	for _, s := range input {
		vowels := 0
		doubles := 0
		forbidden := false
		runes := []rune(s)
		window := string(runes[0])
		vowels += addIfVowels(runes[0])
		for i := 1; i < len(runes); i++ {
			window += string(runes[i])
			vowels += addIfVowels(runes[i])
			doubles += addIfDoubles(runes[i-1], runes[i])
			if isForbidden(window) {
				forbidden = true
				break
			}

			window = string(runes[i])
		}

		if !forbidden && doubles > 0 && vowels >= 3 {
			result++
		}
	}

	return result
}

func Intern2(input []string) int {
	result := 0
	for _, s := range input {
		repeats := 0
		pairs := 0

		runes := []rune(s)
		windowSize := 3
		window := ""
		pairMap := map[string]int{}
		pairWindow := ""

		for i := 0; i < len(runes); i++ {
			window += string(runes[i])
			pairWindow += string(runes[i])

			if len(pairWindow) == 2 {
				if v, ok := pairMap[pairWindow]; ok {
					if i-1 != v {
						pairs++
					}
				} else {
					pairMap[pairWindow] = i
				}
				pairWindow = string([]rune(pairWindow)[1:])
			}

			if len(window) == windowSize {
				if []rune(window)[0] == []rune(window)[2] {
					repeats++
				}
				window = string([]rune(window)[1:])
			}
		}

		if repeats > 0 && pairs > 0 {
			result++
		}
	}

	return result
}
