package advent_2022

import (
	"math"
	"strconv"
	"strings"
)

func distance(h []int, t []int) int {
	a := h[0] - t[0]
	b := h[1] - t[1]

	d := a*a + b*b
	r := math.Sqrt(float64(d))

	return int(r)
}

var unique map[int]bool

func calcCoord(h [][]int, idx1 int, idx2 int, delta int) int {
	headIdx := 0
	h[headIdx][idx1] += delta

	for i := 1; i < len(h); i++ {
		dis := distance(h[headIdx], h[i])

		f := true
		if dis > 1 {
			if h[headIdx][0] == h[i][0] && h[headIdx][1] > h[i][1] && f {
				f = false
				h[i][1]++
			}
			if h[headIdx][0] > h[i][0] && h[headIdx][1] > h[i][1] && f {
				f = false
				h[i][0]++
				h[i][1]++
			}
			if h[headIdx][0] > h[i][0] && h[headIdx][1] == h[i][1] && f {
				f = false
				h[i][0]++
			}
			if h[headIdx][0] > h[i][0] && h[headIdx][1] < h[i][1] && f {
				f = false
				h[i][0]++
				h[i][1]--
			}
			if h[headIdx][0] == h[i][0] && h[headIdx][1] < h[i][1] && f {
				f = false
				h[i][1]--
			}
			if h[headIdx][0] < h[i][0] && h[headIdx][1] < h[i][1] && f {
				f = false
				h[i][0]--
				h[i][1]--
			}
			if h[headIdx][0] < h[i][0] && h[headIdx][1] == h[i][1] && f {
				f = false
				h[i][0]--
			}
			if h[headIdx][0] < h[i][0] && h[headIdx][1] > h[i][1] && f {
				f = false
				h[i][0]--
				h[i][1]++
			}

			if i == len(h)-1 {
				key := 1000*h[i][0] + h[i][1]
				if _, ok := unique[key]; ok {
					return 0
				} else {
					unique[key] = true
					return 1
				}
			}
		}

		headIdx = i
	}

	return 0
}

func Rope1(input []string) int {
	h := [][]int{{0, 0}, {0, 0}}
	s := 1
	unique = make(map[int]bool)
	unique[0] = true

	for _, st := range input {
		parts := strings.Split(st, " ")
		d := parts[0]
		val, _ := strconv.Atoi(parts[1])

		for i := 0; i < val; i++ {
			if d == "R" {
				s += calcCoord(h, 0, 1, 1)
			}
			if d == "L" {
				s += calcCoord(h, 0, 1, -1)
			}

			if d == "U" {
				s += calcCoord(h, 1, 0, 1)
			}

			if d == "D" {
				s += calcCoord(h, 1, 0, -1)
			}
		}
	}

	return s
}

func Rope2(input []string) int {
	h := make([][]int, 0, 10)

	for i := 0; i < 10; i++ {
		el := []int{0, 0}
		h = append(h, el)
	}

	s := 1
	unique = make(map[int]bool)
	unique[0] = true

	for _, st := range input {
		parts := strings.Split(st, " ")
		d := parts[0]
		val, _ := strconv.Atoi(parts[1])

		for i := 0; i < val; i++ {
			if d == "R" {
				s += calcCoord(h, 0, 1, 1)
			}
			if d == "L" {
				s += calcCoord(h, 0, 1, -1)
			}

			if d == "U" {
				s += calcCoord(h, 1, 0, 1)
			}

			if d == "D" {
				s += calcCoord(h, 1, 0, -1)
			}
		}
	}

	return s
}
