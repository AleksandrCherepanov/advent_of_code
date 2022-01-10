package advent_2021

import (
	"fmt"
	"sort"

	"github.com/AleksandrCherepanov/advent_of_code/tiles"
	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_9_1() {
	field := utils.GetFieldInt(utils.GetFile())

	min := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			t := tiles.NewTile(j, i)
			if isMinimumTile(field, t) {
				min += 1 + field[t.Y][t.X]
			}
		}
	}

	fmt.Println(min)
}

func Advent_9_2() {
	field := utils.GetFieldInt(utils.GetFile())

	minPoints := make([]tiles.Tile, 0)
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			t := tiles.NewTile(j, i)
			if isMinimumTile(field, t) {
				minPoints = append(minPoints, t)
			}
		}
	}

	basins := make([]int, 0)
	for _, tile := range minPoints {
		basins = append(basins, findBasin(field, tile))
	}

	sort.Ints(basins)
	result := 1
	for _, v := range basins[len(basins)-3:] {
		result *= v
	}

	fmt.Println(result)
}

func findBasin(field [][]int, point tiles.Tile) int {
	queue := []tiles.Tile{point}
	passed := make(map[string]bool, 0)
	count := 0

	for len(queue) > 0 {
		current := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if _, ok := passed[current.Key]; ok {
			continue
		}

		passed[current.Key] = true

		if needToAddQueue(field, current.TopTile(), passed) {
			queue = append(queue, current.TopTile())
		}

		if needToAddQueue(field, current.RightTile(), passed) {
			queue = append(queue, current.RightTile())
		}

		if needToAddQueue(field, current.BottomTile(), passed) {
			queue = append(queue, current.BottomTile())
		}

		if needToAddQueue(field, current.LeftTile(), passed) {
			queue = append(queue, current.LeftTile())
		}

		count++
	}

	return count
}

func needToAddQueue(field [][]int, t tiles.Tile, passed map[string]bool) bool {
	element := getElement(field, t)

	if element == 9 || element == -1 {
		return false
	}

	return true
}

func getElement(field [][]int, tile tiles.Tile) int {
	if tile.X < 0 || tile.X >= len(field[0]) {
		return -1
	}

	if tile.Y < 0 || tile.Y >= len(field) {
		return -1
	}

	return field[tile.Y][tile.X]
}

func isMinimumTile(field [][]int, t tiles.Tile) bool {
	up := getElement(field, t.TopTile())
	right := getElement(field, t.RightTile())
	down := getElement(field, t.BottomTile())
	left := getElement(field, t.LeftTile())

	current := field[t.Y][t.X]

	if current >= up && up != -1 {
		return false
	}

	if current >= right && right != -1 {
		return false
	}

	if current >= down && down != -1 {
		return false
	}

	if current >= left && left != -1 {
		return false
	}

	return true
}
