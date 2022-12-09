package advent_2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	name  string
	prev  *Node
	nodes map[string]*Node
}

func addNode(root *Node, name string) {
	root.nodes[name] = &Node{name: name, prev: root, nodes: make(map[string]*Node, 0)}
}

func Space1(input []string) int {
	root := &Node{name: "/", prev: nil, nodes: make(map[string]*Node, 0)}
	start := root
	for i := 2; i < len(input); i++ {
		c := input[i]
		if strings.Index(c, "$") < 0 {
			name := ""
			if strings.Index(c, "dir ") == 0 {
				name = strings.Split(c, " ")[1]
			} else {
				name = strings.Split(c, " ")[0]
			}
			addNode(root, name)
		} else {
			if strings.Index(c, "$ ls") == 0 {
				continue
			}
			if strings.Index(c, "$ cd ") == 0 {
				arg := strings.Split(c, " ")[2]
				if arg == ".." {
					root = root.prev
				} else {
					root = root.nodes[arg]
				}
			}
		}
	}

	sumTree(start)
	return total
}

func printTree(root *Node) {
	fmt.Println(root.name, root.nodes)
	for _, v := range root.nodes {
		printTree(v)
	}
}

var total int = 0

func sumTree(root *Node) int {
	if len(root.nodes) == 0 {
		val, _ := strconv.Atoi(root.name)
		return val
	}

	sum := 0
	for _, v := range root.nodes {
		sum += sumTree(v)
	}

	if sum <= 100000 {
		total += sum
	}

	return sum
}

var groupedTree map[string]int = make(map[string]int, 0)

func groupTree(root *Node) int {
	if len(root.nodes) == 0 {
		val, _ := strconv.Atoi(root.name)
		return val
	}

	sum := 0
	for _, v := range root.nodes {
		sum += groupTree(v)
	}

	groupedTree[root.name] = sum
	return sum
}

func Space2(input []string) int {
	root := &Node{name: "/", prev: nil, nodes: make(map[string]*Node, 0)}
	start := root
	for i := 2; i < len(input); i++ {
		c := input[i]
		if strings.Index(c, "$") < 0 {
			name := ""
			if strings.Index(c, "dir ") == 0 {
				name = strings.Split(c, " ")[1]
			} else {
				name = strings.Split(c, " ")[0]
			}
			addNode(root, name)
		} else {
			if strings.Index(c, "$ ls") == 0 {
				continue
			}
			if strings.Index(c, "$ cd ") == 0 {
				arg := strings.Split(c, " ")[2]
				if arg == ".." {
					root = root.prev
				} else {
					root = root.nodes[arg]
				}
			}
		}
	}

	groupTree(start)
	totalSpace := 70000000
	need := 30000000
	usedSpace := groupedTree["/"]

	freeSpace := totalSpace - usedSpace
	size := need - freeSpace

	min := usedSpace
	for _, space := range groupedTree {
		if space >= size && space < min {
			min = space
		}
	}

	return min
}
