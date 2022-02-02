package advent_2021

import (
	"fmt"

	"github.com/AleksandrCherepanov/advent_of_code/tiles"
	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func Advent_11_1() {
	field := utils.GetFieldInt(utils.GetFile("input.txt"))

	steps := 100
	count := 0

	for s := 0; s < steps; s++ {
		willFlash := make([]tiles.Tile, 0)
		flashed := make(map[string]bool, 0)
		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				field[i][j]++
				if field[i][j] > 9 {
					t := tiles.NewTile(j, i)
					willFlash = append(willFlash, t)
				}
			}
		}

		flash(willFlash, flashed, field)
		for key := range flashed {
			tile := tiles.NewTileByKey(key)
			field[tile.Y][tile.X] = 0
			count++
		}
	}

	fmt.Println(count)
}

func Advent_11_2() {
	field := utils.GetFieldInt(utils.GetFile("input.txt"))

	step := 0
	maxCount := len(field) * len(field[0])

	for {
		willFlash := make([]tiles.Tile, 0)
		flashed := make(map[string]bool, 0)
		count := 0
		for i := 0; i < len(field); i++ {
			for j := 0; j < len(field[i]); j++ {
				field[i][j]++
				if field[i][j] > 9 {
					t := tiles.NewTile(j, i)
					willFlash = append(willFlash, t)
				}
			}
		}

		flash(willFlash, flashed, field)
		for key := range flashed {
			tile := tiles.NewTileByKey(key)
			field[tile.Y][tile.X] = 0
			count++
		}
		step++
		if count == maxCount {
			break
		}
	}

	fmt.Println(step)
}

func flash(willFlash []tiles.Tile, flashed map[string]bool, field [][]int) {
	for len(willFlash) > 0 {
		oct := willFlash[0]
		if _, ok := flashed[oct.Key]; !ok {
			flashed[oct.Key] = true
			willFlash = increaseNeighbors(willFlash, field, oct)
		}
		willFlash = willFlash[1:]
	}
}

func increaseNeighbors(willFlash []tiles.Tile, field [][]int, t tiles.Tile) []tiles.Tile {
	willFlash = increaseNeighbor(willFlash, field, t.TopTile())
	willFlash = increaseNeighbor(willFlash, field, t.TopRightTile())
	willFlash = increaseNeighbor(willFlash, field, t.RightTile())
	willFlash = increaseNeighbor(willFlash, field, t.RightBottomTile())
	willFlash = increaseNeighbor(willFlash, field, t.BottomTile())
	willFlash = increaseNeighbor(willFlash, field, t.BottomLeftTile())
	willFlash = increaseNeighbor(willFlash, field, t.LeftTile())
	willFlash = increaseNeighbor(willFlash, field, t.LeftTopTile())

	return willFlash
}

func increaseNeighbor(willFlash []tiles.Tile, field [][]int, tile tiles.Tile) []tiles.Tile {
	if tile.Y >= 0 && tile.Y < len(field) && tile.X >= 0 && tile.X < len(field[tile.Y]) {
		field[tile.Y][tile.X]++
		if field[tile.Y][tile.X] > 9 {
			willFlash = append(willFlash, tile)
		}
	}

	return willFlash
}
