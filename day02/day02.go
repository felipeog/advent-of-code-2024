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

	safeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Fields(text)

		numbers := make([]int, len(substrings))
		for index, substring := range substrings {
			number, _ := strconv.Atoi(substring)
			numbers[index] = number
		}

		// reverse to always deal with decreasing numbers
		if numbers[0]-numbers[len(numbers)-1] < 0 {
			for left, right := 0, len(numbers)-1; left < right; left, right = left+1, right-1 {
				numbers[left], numbers[right] = numbers[right], numbers[left]
			}
		}

		safe := true
		for index := range len(numbers) - 1 {
			difference := numbers[index] - numbers[index+1]
			if difference < 1 || difference > 3 {
				safe = false
				break
			}
		}

		if safe {
			safeCount++
		}
	}

	return safeCount
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	reverse := func(list []int) []int {
		listCopy := make([]int, len(list))
		copy(listCopy, list)
		for left, right := 0, len(listCopy)-1; left < right; left, right = left+1, right-1 {
			listCopy[left], listCopy[right] = listCopy[right], listCopy[left]
		}
		return listCopy
	}

	getInvalidIndex := func(list []int) int {
		invalidIndex := -1
		for index := range len(list) - 1 {
			difference := list[index] - list[index+1]
			if difference < 1 || difference > 3 {
				invalidIndex = index
				break
			}
		}
		return invalidIndex
	}

	safeCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Fields(text)

		numbers := make([]int, len(substrings))
		for index, substring := range substrings {
			number, _ := strconv.Atoi(substring)
			numbers[index] = number
		}

		// reverse to always deal with decreasing numbers
		if numbers[0]-numbers[len(numbers)-1] < 0 {
			numbers = reverse(numbers)
		}

		// if it's valid, increment
		invalidIndex := getInvalidIndex(numbers)
		if invalidIndex < 0 {
			safeCount++
			continue
		}

		// if it's not valid, try removing the invalid index
		removed := append(append([]int{}, numbers[:invalidIndex]...), numbers[invalidIndex+1:]...)
		if getInvalidIndex(removed) < 0 {
			safeCount++
			continue
		}

		// if it's still not valid, try removing each index but the invalid index
		for index := range len(numbers) {
			if index == invalidIndex {
				continue
			}
			removed := append(append([]int{}, numbers[:index]...), numbers[index+1:]...)
			if getInvalidIndex(removed) < 0 {
				safeCount++
				break
			}
		}
	}

	return safeCount
}
