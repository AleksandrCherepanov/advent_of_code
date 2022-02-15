package advent_2021

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func p1Turn(step int, pos int) (int, int) {
	newPos := pos + (6 * step)
	score := newPos % 10
	if score == 0 {
		score = 10
	}
	return newPos, score
}

func p2Turn(step, offset int, pos int) (int, int) {
	newPos := pos + (15*step + 3*offset)
	score := newPos % 10
	if score == 0 {
		score = 10
	}
	return newPos, score
}

func Advent_21_1() {
	input := utils.GetFile("input.txt")

	p1Pos, _ := strconv.Atoi(strings.Split(input[0], ": ")[1])
	p2Pos, _ := strconv.Atoi(strings.Split(input[1], ": ")[1])

	fmt.Println(p1Pos, p2Pos)

	score1 := 0
	score2 := 0

	j := 1
	i := 0
	s := 1
	t := 0

	for score1 < 1000 && score2 < 1000 {
		if t%2 == 0 {
			nPos, score := p1Turn(j, p1Pos)
			j += 3
			p1Pos = nPos
			score1 += score
		} else {
			nPos, score := p2Turn(s, i, p2Pos)
			s++
			i++
			p2Pos = nPos
			score2 += score
		}
		t++
	}

	min := math.Min(float64(score1), float64(score2))
	fmt.Println(min * float64(t) * float64(3))

}
