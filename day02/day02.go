package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FirstHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", -1)

		numList := make([]int, len(parts))
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			numList[i] = num
		}

		if numList[0]-numList[1] < 0 {
			for leftIndex, rightIndex := 0, len(numList)-1; leftIndex < rightIndex; leftIndex, rightIndex = leftIndex+1, rightIndex-1 {
				numList[leftIndex], numList[rightIndex] = numList[rightIndex], numList[leftIndex]
			}
		}

		safe := true
		for i := range len(numList) - 1 {
			difference := numList[i] - numList[i+1]
			if difference < 1 || difference > 3 {
				safe = false
				break
			}
		}
		if safe {
			safeCount = safeCount + 1
		}
	}

	return safeCount
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", -1)

		numList := make([]int, len(parts))
		for i, part := range parts {
			num, _ := strconv.Atoi(part)
			numList[i] = num
		}

		reverse := func(list []int) []int {
			for leftIndex, rightIndex := 0, len(list)-1; leftIndex < rightIndex; leftIndex, rightIndex = leftIndex+1, rightIndex-1 {
				list[leftIndex], list[rightIndex] = list[rightIndex], list[leftIndex]
			}
			return list
		}

		if numList[0]-numList[1] < 0 {
			numList = reverse(numList)
		}

		isSafe := func(list []int) bool {
			safe := true
			for i := range len(list) - 1 {
				difference := list[i] - list[i+1]
				if difference < 1 || difference > 3 {
					safe = false
					break
				}
			}
			return safe
		}

		if isSafe(numList) {
			safeCount = safeCount + 1
		} else {
			safeIfRemoved := false
			for i := range len(numList) {
				numListCopy := make([]int, len(numList))
				copy(numListCopy, numList)

				removed := append(numListCopy[:i], numList[i+1:]...)
				if isSafe(removed) || isSafe(reverse(removed)) {
					safeIfRemoved = true
					break
				}
			}
			if safeIfRemoved {
				safeCount = safeCount + 1
			}
		}
	}

	return safeCount
}
