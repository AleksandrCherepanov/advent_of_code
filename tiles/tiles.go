package tiles

import (
	"strconv"
	"strings"
)

type Tile struct {
	X   int
	Y   int
	Key string
}

func NewTile(x, y int) Tile {
	tile := Tile{
		X: x,
		Y: y,
	}

	tile.updateTileKey()
	return tile
}

func NewTileByKey(key string) Tile {
	tile := Tile{
		Key: key,
	}

	tile.splitTileKey()
	return tile
}

func (tile *Tile) updateTileKey() {
	tile.Key = strconv.Itoa(tile.X) + ":" + strconv.Itoa(tile.Y)
}

func (tile *Tile) splitTileKey() {
	pair := strings.Split(tile.Key, ":")
	tile.X, _ = strconv.Atoi(pair[0])
	tile.Y, _ = strconv.Atoi(pair[1])
}

func (tile *Tile) TopTile() Tile {
	return NewTile(tile.X, tile.Y-1)
}

func (tile *Tile) TopRightTile() Tile {
	return NewTile(tile.X+1, tile.Y-1)
}

func (tile *Tile) RightTile() Tile {
	return NewTile(tile.X+1, tile.Y)
}

func (tile *Tile) RightBottomTile() Tile {
	return NewTile(tile.X+1, tile.Y+1)
}

func (tile *Tile) BottomTile() Tile {
	return NewTile(tile.X, tile.Y+1)
}

func (tile *Tile) BottomLeftTile() Tile {
	return NewTile(tile.X-1, tile.Y+1)
}

func (tile *Tile) LeftTile() Tile {
	return NewTile(tile.X-1, tile.Y)
}

func (tile *Tile) LeftTopTile() Tile {
	return NewTile(tile.X-1, tile.Y-1)
}
