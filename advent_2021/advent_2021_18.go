package advent_2021

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	_ "strings"

	"github.com/AleksandrCherepanov/advent_of_code/utils"
)

func add(list1, list2 string) string {
	return "[" + list1 + "," + list2 + "]"
}

func reduce(list string) string {
	return ""
}

func explode(list string) string {
	re, _ := regexp.Compile(`\[\d+,\d+\]`)
	leftRe, _ := regexp.Compile(`(\d+\])|(\d+,\[)|(\d+,$)`)
	rightRe, _ := regexp.Compile(`\d+`)

	tmpList := list
	for {
		nestedPairsRange := re.FindAllStringIndex(tmpList, -1)
		if nestedPairsRange == nil {
			break
		}

		changed := false
		for i := 0; i < len(nestedPairsRange); i++ {
			nestedPairRange := nestedPairsRange[i]
			s := tmpList[:nestedPairRange[1]]
			if strings.Count(s, "[")-strings.Count(s, "]") < 4 {
				continue
			}

			leftNested := strings.ReplaceAll(strings.Split(tmpList[nestedPairRange[0]:nestedPairRange[1]], ",")[0], "[", "")
			leftNestedNumber, _ := strconv.Atoi(leftNested)

			rightNested := strings.ReplaceAll(strings.Split(tmpList[nestedPairRange[0]:nestedPairRange[1]], ",")[1], "]", "")
			rightNestedNumber, _ := strconv.Atoi(rightNested)

			left := tmpList[:nestedPairRange[0]]
			leftNumbersRange := leftRe.FindAllStringIndex(left, -1)
			fmt.Println(left, leftNumbersRange)
			if leftNumbersRange != nil {
				leftNumberRange := leftNumbersRange[len(leftNumbersRange)-1]
				leftNumber, _ := strconv.Atoi(left[leftNumberRange[0] : leftNumberRange[1]-1])
				leftSum := leftNumber + leftNestedNumber
				left = left[:leftNumberRange[0]] + strconv.Itoa(leftSum) + left[leftNumberRange[1]-1:]
				changed = true
			}

			right := tmpList[nestedPairRange[1]:]
			rightNumberRange := rightRe.FindStringIndex(right)
			if rightNumberRange != nil {
				rightNumber, _ := strconv.Atoi(right[rightNumberRange[0]:rightNumberRange[1]])
				rightSum := rightNumber + rightNestedNumber
				right = right[:rightNumberRange[0]] + strconv.Itoa(rightSum) + right[rightNumberRange[1]:]
				changed = true
			}

			tmpList = left + "0" + right
			if changed {
				break
			}
		}

		if !changed {
			break
		}
	}

	return tmpList
}

func split_pair(list string) string {
	return ""
}

func needExplode(list string) bool {

	return true
}

func needSplit(list string) bool {
	return true
}

func Advent_18_1() {
	input := utils.GetFile("input.txt")
	fmt.Println(input)

	test := []string{
		"[[[[[9,8],1],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]",
		"[[6,[5,[4,[3,2]]]],1]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
	}

	expected := []string{
		"[[[[0,9],2],3],4]",
		"[7,[6,[5,[7,0]]]]",
		"[[6,[5,[7,0]]],3]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}

	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
		result := explode(test[i])
		if result != expected[i] {
			fmt.Println("actual: " + result)
			fmt.Println("expected: " + expected[i])
		}
	}

	//for _, line := range input {
	//	fmt.Println(line)
	//	fmt.Println(needExplode(line))
	//}
}
