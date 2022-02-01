package advent_2021

import (
	"fmt"
	"sort"

	"github.com/AleksandrCherepanov/advent_of_code/field"
	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

type priorityElement struct {
	priority int
	element  field.CellInt
}

func (pq *priorityQueue) Add(element field.CellInt, priority int) {
	priorityElement := priorityElement{
		priority: priority,
		element:  element,
	}

	pq.elements = append(pq.elements, priorityElement)
}

func (pq *priorityQueue) Get() field.CellInt {
	sort.Slice(pq.elements, func(i int, j int) bool {
		return pq.elements[i].priority < pq.elements[j].priority
	})

	element := pq.elements[0]
	pq.elements = pq.elements[1:]
	return element.element
}

func (pq *priorityQueue) Empty() bool {
	return len(pq.elements) == 0
}

type priorityQueue struct {
	elements []priorityElement
}

func search(f field.FieldInt, start, end field.CellInt) int {
	queue := &priorityQueue{
		elements: make([]priorityElement, 0, f.GetSize()),
	}
	queue.Add(start, 0)
	came_from := make(map[string]field.CellInt, 0)
	came_from[start.GetIndex()] = field.CellInt{}
	cost_from_start := make(map[string]int, 0)
	cost_from_start[start.GetIndex()] = 0

	for !queue.Empty() {
		current := queue.Get()
		if current.Equals(end) {
			break
		}

		adjacent := make([]field.CellInt, 0, 4)
		top, err := f.Top(current)
		if err == nil {
			adjacent = append(adjacent, top)
		}

		right, err := f.Right(current)
		if err == nil {
			adjacent = append(adjacent, right)
		}

		bottom, err := f.Bottom(current)
		if err == nil {
			adjacent = append(adjacent, bottom)
		}

		left, err := f.Left(current)
		if err == nil {
			adjacent = append(adjacent, left)
		}

		for _, newCell := range adjacent {
			newCost := cost_from_start[current.GetIndex()] + current.GetValue() + newCell.GetValue()
			if _, ok := cost_from_start[newCell.GetIndex()]; !ok || newCost < cost_from_start[newCell.GetIndex()] {
				cost_from_start[newCell.GetIndex()] = newCost
				queue.Add(newCell, newCost)
				came_from[newCell.GetIndex()] = current
			}
		}
	}

	return cost_from_start[end.GetIndex()]
}

func Advent_15_1() {
	input := utils.GetFieldInt(utils.GetFile())
	f := field.NewFieldInt(input)
	result := search(f, f.GetCell(0, 0), f.GetCell(len(f[0])-1, len(f)-1))
	fmt.Println(result)
}
