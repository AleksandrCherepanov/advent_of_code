package advent_2021

import (
	"fmt"
	"math"
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
	tmpList := list
	do := true
	for do {
		r := explode(tmpList)
		fmt.Println("E: ", r)
		r, ch := splitNumber(r)
		fmt.Println("R: ", r)
		tmpList = r
		if ch {
			continue
		}
		do = ch
	}

	return tmpList
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

func splitNumber(list string) (string, bool) {
	re, _ := regexp.Compile(`\d\d+`)

	tmpList := list

	bigNumberRange := re.FindStringIndex(tmpList)
	if bigNumberRange == nil {
		return tmpList, false
	}

	left := tmpList[:bigNumberRange[0]]
	right := tmpList[bigNumberRange[1]:]

	numberString := tmpList[bigNumberRange[0]:bigNumberRange[1]]
	numberInt, _ := strconv.Atoi(numberString)

	var half float64
	half = float64(numberInt) / float64(2)
	lowHalfInt := int(math.Floor(half))
	highHalfInt := int(math.Ceil(half))

	tmpList = left + "[" + strconv.Itoa(lowHalfInt) + "," + strconv.Itoa(highHalfInt) + "]" + right

	return tmpList, true
}

func Advent_18_1() {
	input := utils.GetFile("input.txt")
	fmt.Println(input)

	test := "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"
	r := reduce(test)
	fmt.Println(r)

	//test := []string{
	//	"[[[[[9,8],1],2],3],4]",
	//	"[7,[6,[5,[4,[3,2]]]]]",
	//	"[[6,[5,[4,[3,2]]]],1]",
	//	"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
	//}

	//expected := []string{
	//	"[[[[0,9],2],3],4]",
	//	"[7,[6,[5,[7,0]]]]",
	//	"[[6,[5,[7,0]]],3]",
	//	"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	//}

	//for i := 0; i < len(test); i++ {
	//	result := explode(test[i])
	//	if result != expected[i] {
	//		fmt.Println("actual: " + result)
	//		fmt.Println("expected: " + expected[i])
	//	}
	//}

	//test2 := "[[[[0,7],4],[15,[0,13]]],[1,1]]"

	//tr := true
	//for tr {
	//	r, ch := splitNumber(test2)
	//	fmt.Println(r)
	//	tr = ch
	//	test2 = r
	//}
	//for _, line := range input {
	//	fmt.Println(line)
	//	fmt.Println(needExplode(line))
	//}
}
