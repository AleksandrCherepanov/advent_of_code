package advent_2022

import (
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func (p point) down() point {
	return point{
		p.x,
		p.y + 1,
	}
}

func (p point) downLeft() point {
	return point{
		p.x - 1,
		p.y + 1,
	}
}

func (p point) downRight() point {
	return point{
		p.x + 1,
		p.y + 1,
	}
}

type segment struct {
	start point
	end   point
}

func (s segment) contains(p point) bool {
	if p.x == s.start.x && p.x == s.end.x && p.y >= s.start.y && p.y <= s.end.y {
		return true
	}
	if p.x == s.start.x && p.x == s.end.x && p.y >= s.end.y && p.y <= s.start.y {
		return true
	}
	if p.y == s.start.y && p.y == s.end.y && p.x >= s.start.x && p.x <= s.end.x {
		return true
	}
	if p.y == s.start.y && p.y == s.end.y && p.x >= s.end.x && p.x <= s.start.x {
		return true
	}

	return false
}

func canMove(segments []segment, sands map[int]bool, p point) bool {
	for _, seg := range segments {
		if seg.contains(p) {
			return false
		}
	}

	if _, ok := sands[p.x*1000+p.y]; ok {
		return false
	}

	return true
}

func Regolith1(input []string) int {
	lowest := 0
	segments := make([]segment, 0)
	for _, s := range input {
		parts := strings.Split(s, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			start := strings.Split(parts[i], ",")
			end := strings.Split(parts[i+1], ",")

			seg := segment{}

			seg.start.x, _ = strconv.Atoi(start[0])
			seg.start.y, _ = strconv.Atoi(start[1])

			seg.end.x, _ = strconv.Atoi(end[0])
			seg.end.y, _ = strconv.Atoi(end[1])

			if seg.start.y > lowest {
				lowest = seg.start.y
			}
			if seg.end.y > lowest {
				lowest = seg.end.y
			}
			segments = append(segments, seg)
		}
	}

	sands := make(map[int]bool, 0)
	sand := point{500, 0}

	for {
		if sand.y > lowest {
			break
		}

		downSand := sand.down()
		leftSand := sand.downLeft()
		rightSand := sand.downRight()

		cm := canMove(segments, sands, downSand)
		if cm {
			sand = downSand
			continue
		}

		cm = canMove(segments, sands, leftSand)
		if cm {
			sand = leftSand
			continue
		}

		cm = canMove(segments, sands, rightSand)
		if cm {
			sand = rightSand
			continue
		}

		sands[sand.x*1000+sand.y] = true
		sand = point{500, 0}
	}

	return len(sands)
}

func Regolith2(input []string) int {
	lowest := 0
	segments := make([]segment, 0)
	for _, s := range input {
		parts := strings.Split(s, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			start := strings.Split(parts[i], ",")
			end := strings.Split(parts[i+1], ",")

			seg := segment{}

			seg.start.x, _ = strconv.Atoi(start[0])
			seg.start.y, _ = strconv.Atoi(start[1])

			seg.end.x, _ = strconv.Atoi(end[0])
			seg.end.y, _ = strconv.Atoi(end[1])

			if seg.start.y > lowest {
				lowest = seg.start.y
			}
			if seg.end.y > lowest {
				lowest = seg.end.y
			}
			segments = append(segments, seg)
		}
	}

	lowest += 2
	sands := make(map[int]bool, 0)
	sand := point{500, 0}

	for {
		downSand := sand.down()
		leftSand := sand.downLeft()
		rightSand := sand.downRight()

		cm := canMove(segments, sands, downSand)
		if cm && downSand.y < lowest {
			sand = downSand
			continue
		}

		cm = canMove(segments, sands, leftSand)
		if cm && leftSand.y < lowest {
			sand = leftSand
			continue
		}

		cm = canMove(segments, sands, rightSand)
		if cm && rightSand.y < lowest {
			sand = rightSand
			continue
		}

		sands[sand.x*1000+sand.y] = true
		if sand.x == 500 && sand.y == 0 {
			break
		}
		sand = point{500, 0}
	}

	return len(sands)
}
