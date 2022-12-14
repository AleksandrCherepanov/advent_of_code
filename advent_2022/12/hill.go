package advent_2022

import (
	"math"
	"sort"

	"github.com/AleksandrCherepanov/advent_of_code/field"
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

func canStep(next field.CellInt, current field.CellInt) bool {
	result := next.GetValue()-current.GetValue() <= 1
	return result
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

	step := 0
	for !queue.Empty() {
		step++
		current := queue.Get()
		if current.Equals(end) {
			break
		}

		adjacent := make([]field.CellInt, 0, 4)
		top, err := f.Top(current)
		if err == nil {
			if canStep(top, current) {
				adjacent = append(adjacent, top)
			}
		}

		right, err := f.Right(current)
		if err == nil {
			if canStep(right, current) {
				adjacent = append(adjacent, right)
			}
		}

		bottom, err := f.Bottom(current)
		if err == nil {
			if canStep(bottom, current) {
				adjacent = append(adjacent, bottom)
			}
		}

		left, err := f.Left(current)
		if err == nil {
			if canStep(left, current) {
				adjacent = append(adjacent, left)
			}
		}

		for _, newCell := range adjacent {
			newCost := cost_from_start[current.GetIndex()] + 1
			if _, ok := cost_from_start[newCell.GetIndex()]; !ok || newCost < cost_from_start[newCell.GetIndex()] {
				cost_from_start[newCell.GetIndex()] = newCost
				priority := newCost
				queue.Add(newCell, priority)
				came_from[newCell.GetIndex()] = current
			}
		}
	}

	return cost_from_start[end.GetIndex()]
}

func Hill1(input []string) int {
	inputInt := make([][]int, 0)
	start := []int{0, 0}
	end := []int{0, 0}
	for i, s := range input {
		line := make([]int, 0, len(s))
		for j, b := range s {
			v := int(b)
			if b == rune('S') {
				v = int(rune('a')) - 1
				start[0] = i
				start[1] = j
			}

			if b == rune('E') {
				v = int(rune('z')) + 1
				end[0] = i
				end[1] = j
			}
			line = append(line, v)
		}
		inputInt = append(inputInt, line)
	}

	f := field.NewFieldInt(inputInt)
	result := search(f, f.GetCell(start[1], start[0]), f.GetCell(end[1], end[0]))
	return result
}

func Hill2(input []string) int {
	inputInt := make([][]int, 0)
	start := make([][]int, 0)
	end := []int{0, 0}
	for i, s := range input {
		line := make([]int, 0, len(s))
		for j, b := range s {
			v := int(b)
			if b == rune('S') {
				v = int(rune('a')) - 1
				start = append(start, []int{i, j})
			}

			if b == rune('a') {
				v = int(rune('a'))
				start = append(start, []int{i, j})
			}

			if b == rune('E') {
				v = int(rune('z')) + 1
				end[0] = i
				end[1] = j
			}
			line = append(line, v)
		}
		inputInt = append(inputInt, line)
	}

	f := field.NewFieldInt(inputInt)
	min := math.MaxInt

	for _, st := range start {
		result := search(f, f.GetCell(st[1], st[0]), f.GetCell(end[1], end[0]))
		if result == 0 {
			continue
		}
		if result < min {
			min = result
		}
	}
	return min
}
