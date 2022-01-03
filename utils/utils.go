package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFile() []string {
	fileData, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(fileData), "\n")
	lastIndex := len(lines) - 1
	if lines[lastIndex] == "" {
		lines = lines[:lastIndex]
	}
	return lines
}

func ConvertSliceStringToInt(stringSlice []string) []int {
	result := make([]int, 0, len(stringSlice))

	for _, v := range stringSlice {
		if v == "" {
			continue
		}

		intValue, err := strconv.Atoi(v)
		if err == nil {
			result = append(result, intValue)
			continue
		}
		fmt.Println(err)
	}

	return result
}

func SplitSliceStringToStringIntSlice(stringSlice []string) ([]string, []int) {
	resultString := make([]string, 0, len(stringSlice))
	resultInt := make([]int, 0, len(stringSlice))

	for _, v := range stringSlice {
		if v == "" {
			continue
		}
		pair := strings.Split(v, " ")
		intValue, err := strconv.Atoi(pair[1])
		if err == nil {
			resultString = append(resultString, pair[0])
			resultInt = append(resultInt, intValue)
			continue
		}
		fmt.Println(err)
	}

	return resultString, resultInt
}

func ReadFileByLine() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func SumSliceInt(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}

	return sum
}
