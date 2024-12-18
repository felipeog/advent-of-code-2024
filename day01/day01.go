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

	leftList := []int{}
	rightList := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.SplitN(text, "   ", 2)

		leftValue, _ := strconv.Atoi(substrings[0])
		rightValue, _ := strconv.Atoi(substrings[1])

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

	sum := 0
	for index := range leftList {
		sum += absInt(leftList[index] - rightList[index])
	}

	return sum
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	leftList := []int{}
	rightList := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.SplitN(text, "   ", 2)

		leftValue, _ := strconv.Atoi(substrings[0])
		rightValue, _ := strconv.Atoi(substrings[1])

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	rightListFrequencyMap := make(map[int]int)
	for _, rightValue := range rightList {
		rightListFrequencyMap[rightValue]++
	}

	sum := 0
	for _, leftValue := range leftList {
		frequency, exists := rightListFrequencyMap[leftValue]
		if exists {
			sum += leftValue * frequency
		}
	}

	return sum
}
