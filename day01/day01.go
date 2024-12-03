package day01

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FirstHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "   ", 2)

		leftValue, _ := strconv.Atoi(parts[0])
		rightValue, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	absInt := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	for index := range leftList {
		sum = sum + absInt(leftList[index]-rightList[index])
	}

	return sum
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "   ", 2)

		leftValue, _ := strconv.Atoi(parts[0])
		rightValue, _ := strconv.Atoi(parts[1])

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	rightListMap := make(map[int]int)
	for index := range rightList {
		mapKey := rightList[index]
		count, exists := rightListMap[mapKey]

		if exists {
			rightListMap[mapKey] = count + 1
		} else {
			rightListMap[mapKey] = 1
		}
	}

	for index := range leftList {
		leftItem := leftList[index]
		rightCount, exists := rightListMap[leftItem]

		if exists {
			sum = sum + leftItem*rightCount
		}
	}

	return sum
}
