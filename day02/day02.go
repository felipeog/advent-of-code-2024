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

		if numbers[0]-numbers[1] < 0 {
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
		for left, right := 0, len(list)-1; left < right; left, right = left+1, right-1 {
			list[left], list[right] = list[right], list[left]
		}
		return list
	}

	isSafe := func(list []int) bool {
		safe := true
		for index := range len(list) - 1 {
			difference := list[index] - list[index+1]
			if difference < 1 || difference > 3 {
				safe = false
				break
			}
		}
		return safe
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

		if numbers[0]-numbers[1] < 0 {
			reverse(numbers)
		}

		if isSafe(numbers) {
			safeCount++
			continue
		}

		safeIfRemoved := false
		for index := range len(numbers) {
			numbersCopy := make([]int, len(numbers))
			copy(numbersCopy, numbers)

			removed := append(numbersCopy[:index], numbers[index+1:]...)
			if isSafe(removed) {
				safeIfRemoved = true
				break
			}

			reverse(removed)
			if isSafe(removed) {
				safeIfRemoved = true
				break
			}
		}

		if safeIfRemoved {
			safeCount++
		}
	}

	return safeCount
}
