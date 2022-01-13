package advent_2021

import (
	"fmt"
	"strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

type PathCounter struct {
	graph        map[string][]string
	pathCount    int
	visitedNodes map[string]int
	rules        map[string]int
	paths        map[string]bool
	path         string
}

func (pc *PathCounter) init() {
	pc.visitedNodes = make(map[string]int, 0)
	pc.rules = map[string]int{
		"start": 1,
		"end":   1,
	}
	pc.path = ""
}

func (pc *PathCounter) count(src, dst string) {
	if !isBigCave(src) {
		pc.visitedNodes[src]++
	}

	if src == dst {
		pc.pathCount++
	} else {
		for _, n := range pc.graph[src] {
			if pc.visitedNodes[n] == 0 {
				pc.count(n, dst)
			}
		}
	}

	if !isBigCave(src) {
		pc.visitedNodes[src]--
	}
}

func (pc *PathCounter) countWithSingleCaves(src, dst string) {
	if !isBigCave(src) {
		pc.visitedNodes[src]++
	}

	pc.path += src
	if src == dst {
		pc.paths[pc.path] = true
	} else {
		for _, n := range pc.graph[src] {
			limit, ok := pc.rules[n]
			if !ok {
				limit = 1
			}
			if pc.visitedNodes[n] < limit || isBigCave(n) {
				pc.countWithSingleCaves(n, dst)
			}
		}
	}

	if !isBigCave(src) {
		pc.visitedNodes[src]--
	}
	pc.path = strings.TrimRight(pc.path, src)
}

func Advent_12_1() {
	input := utils.GetFile()

	graph := make(map[string][]string, len(input))
	for _, line := range input {
		pair := strings.Split(line, "-")
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	pathCounter := &PathCounter{
		graph:        graph,
		visitedNodes: make(map[string]int, 0),
	}

	pathCounter.count("start", "end")

	fmt.Println(pathCounter.pathCount)
}

func isBigCave(cave string) bool {
	return strings.ToUpper(cave) == cave
}

func Advent_12_2() {
	input := utils.GetFile()

	graph := make(map[string][]string, len(input))
	for _, line := range input {
		pair := strings.Split(line, "-")
		graph[pair[0]] = append(graph[pair[0]], pair[1])
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	singleCaves := make(map[string]bool, 0)
	for n := range graph {
		if !isBigCave(n) && n != "start" && n != "end" {
			singleCaves[n] = false
		}
	}

	pathCounter := &PathCounter{
		graph: graph,
		paths: make(map[string]bool),
	}

	pathCounter.init()
	for cave := range singleCaves {
		pathCounter.rules[cave] = 2
		pathCounter.countWithSingleCaves("start", "end")
		pathCounter.init()
	}

	fmt.Println(len(pathCounter.paths))
}
