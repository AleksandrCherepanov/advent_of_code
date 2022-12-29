package advent_2022

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type seg struct {
	start point
	end   point
}

var field [23][28]string

func abs(a, b int) int {
	d := a - b
	if d < 0 {
		return d * -1
	}

	return d
}

func processInput(input []string) ([]point, []point) {
	sensors := []point{}
	beacons := []point{}

	for _, line := range input {
		parts := strings.Split(line, ":")
		sensorIndex := strings.Index(parts[0], "x=")
		sensor := string([]rune(parts[0])[sensorIndex:])

		sensorCoords := strings.Split(sensor, ",")
		sxString := strings.Trim(strings.Split(sensorCoords[0], "=")[1], " ")
		syString := strings.Trim(strings.Split(sensorCoords[1], "=")[1], " ")

		spoint := point{}
		spoint.x, _ = strconv.Atoi(sxString)
		spoint.y, _ = strconv.Atoi(syString)
		sensors = append(sensors, spoint)

		beaconIndex := strings.Index(parts[1], "x=")
		beacon := string([]rune(parts[1])[beaconIndex:])

		beaconCoords := strings.Split(beacon, ",")
		bxString := strings.Trim(strings.Split(beaconCoords[0], "=")[1], " ")
		byString := strings.Trim(strings.Split(beaconCoords[1], "=")[1], " ")

		bpoint := point{}
		bpoint.x, _ = strconv.Atoi(bxString)
		bpoint.y, _ = strconv.Atoi(byString)
		beacons = append(beacons, bpoint)
	}

	return sensors, beacons
}

func Beacon1(Y int, input []string) int {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = "."
		}
	}

	sensors, beacons := processInput(input)

	minLeft := math.MaxInt
	maxRight := 0
	for i := 0; i < len(sensors); i++ {
		s := sensors[i]
		b := beacons[i]
		distance := abs(s.x, b.x) + abs(s.y, b.y)

		if s.y < Y && s.y+distance >= Y {
			r := Y - s.y
			amount := (distance*2 + 1) - r*2
			left := s.x - amount/2
			right := s.x + amount/2
			if left < minLeft {
				minLeft = left
			}

			if right > maxRight {
				maxRight = right
			}
		}

		if s.y > Y && s.y-distance <= Y {
			r := s.y - Y
			amount := (distance*2 + 1) - r*2
			left := s.x - amount/2
			right := s.x + amount/2

			if left < minLeft {
				minLeft = left
			}

			if right > maxRight {
				maxRight = right
			}
		}

		if s.y == Y {
			r := 0
			amount := (distance*2 + 1) - r*2
			left := s.x - amount/2
			right := s.x + amount/2

			if left < minLeft {
				minLeft = left
			}

			if right > maxRight {
				maxRight = right
			}
		}
	}

	return maxRight - minLeft
}

type segments []seg

func (s segments) Len() int {
	return len(s)
}

func (s segments) Less(a, b int) bool {
	return s[a].start.x < s[b].start.x
}

func (s segments) Swap(a, b int) {
	temp := s[a]
	s[a] = s[b]
	s[b] = temp
}

func Beacon2(Y int, input []string) int64 {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = "."
		}
	}

	sensors, beacons := processInput(input)

	dic := map[int]segments{}

	for i := 0; i < len(sensors); i++ {
		s := sensors[i]
		b := beacons[i]
		distance := abs(s.x, b.x) + abs(s.y, b.y)

		for j := 0; j < distance; j++ {
			key := s.y + j
			if key > Y {
				break
			}
			amount := (distance*2 + 1) - j*2
			left := s.x - amount/2
			if left < 0 {
				left = 0
			}
			right := s.x + amount/2
			if right > Y {
				right = Y
			}

			sg := seg{
				point{left, key},
				point{right, key},
			}
			dic[key] = append(dic[key], sg)
		}

		for j := 0; j < distance; j++ {
			key := s.y - j
			if key < 0 {
				break
			}
			amount := (distance*2 + 1) - j*2
			left := s.x - amount/2
			if left < 0 {
				left = 0
			}
			right := s.x + amount/2
			if right > Y {
				right = Y
			}

			sg := seg{
				point{left, key},
				point{right, key},
			}
			dic[key] = append(dic[key], sg)
		}
	}

	result := int64(0)
	for y, v := range dic {
		sort.Sort(v)

		c := 0
		index := 0

		for i := 1; i < len(v); i++ {
			if v[index].end.x >= v[i].start.x {
				max := 0
				if v[index].end.x > v[i].end.x {
					max = v[index].end.x
				} else {
					max = v[i].end.x
				}

				v[index].end.x = max
				c++
				continue
			}

			index++
			v[index] = v[i]
		}

		dic[y] = v[:len(v)-c]
		if len(dic[y]) == 2 {
			if dic[y][1].start.x-dic[y][0].end.x == 2 {
				result = int64((dic[y][1].start.x-1)*4000000) + int64(y)
			}
		}
	}

	return result
}
