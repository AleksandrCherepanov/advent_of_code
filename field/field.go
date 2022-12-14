package field

import (
	"errors"
	"fmt"
	"strconv"
)

type FieldInt [][]CellInt

type CellInt struct {
	x     int
	y     int
	value int
}

func (c CellInt) SetValue(value int) {
	c.value = value
}

func (c CellInt) GetValue() int {
	return c.value
}

func (c CellInt) GetX() int {
	return c.x
}

func (c CellInt) GetY() int {
	return c.y
}

func (c CellInt) GetIndex() string {
	return strconv.Itoa(c.GetX()*1000 + c.GetY())
}

func (c CellInt) Equals(cell CellInt) bool {
	return c.GetIndex() == cell.GetIndex()
}

func NewCellInt(x, y, value int) CellInt {
	return CellInt{
		x:     x,
		y:     y,
		value: value,
	}
}

func NewFieldInt(matrix [][]int) FieldInt {
	if len(matrix) == 0 {
		return make(FieldInt, 0)
	}

	field := make(FieldInt, 0, len(matrix))
	for y, row := range matrix {
		cellRow := make([]CellInt, 0, len(row))
		for x, col := range row {
			cellRow = append(cellRow, NewCellInt(x, y, col))
		}
		field = append(field, cellRow)
	}

	return field
}

func (field FieldInt) GetCell(x, y int) CellInt {
	return field[y][x]
}

func (field FieldInt) InBounds(x, y int) bool {
	if y < 0 || y >= len(field) {
		return false
	}

	if x < 0 || x >= len(field[y]) {
		return false
	}

	return true
}

func (field FieldInt) getAdjacent(x, y int) (CellInt, error) {
	if field.InBounds(x, y) {
		return field.GetCell(x, y), nil
	}

	return CellInt{}, errors.New("Out of range")
}

func (field FieldInt) Top(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX(), cell.GetY()-1)
}

func (field FieldInt) TopRight(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX()+1, cell.GetY()-1)
}

func (field FieldInt) Right(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX()+1, cell.GetY())
}

func (field FieldInt) RightBottom(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX()+1, cell.GetY()+1)
}

func (field FieldInt) Bottom(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX(), cell.GetY()+1)
}

func (field FieldInt) BottomLeft(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX()-1, cell.GetY()+1)
}

func (field FieldInt) Left(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX()-1, cell.GetY())
}

func (field FieldInt) LeftTop(cell CellInt) (CellInt, error) {
	return field.getAdjacent(cell.GetX()-1, cell.GetY()-1)
}

func (field FieldInt) Print() {
	for _, row := range field {
		for _, cell := range row {
			fmt.Print(cell.GetValue(), " ")
		}
		fmt.Println()
	}
}

func (field FieldInt) GetSize() int {
	if len(field) == 0 {
		return 0
	}
	return len(field) * len(field[0])
}
