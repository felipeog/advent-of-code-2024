package day05

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

	rulesLeftMap := make(map[int][]int)
	rulesRightMap := make(map[int][]int)
	updates := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// stop scanning rules to scan updates
		text := scanner.Text()
		if text == "" {
			break
		}

		substrings := strings.SplitN(text, "|", 2)
		leftValue, _ := strconv.Atoi(substrings[0])
		rightValue, _ := strconv.Atoi(substrings[1])

		rulesLeftMap[leftValue] = append(rulesLeftMap[leftValue], rightValue)
		rulesRightMap[rightValue] = append(rulesRightMap[rightValue], leftValue)
	}
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, ",")
		update := []int{}
		for _, substring := range substrings {
			page, _ := strconv.Atoi(substring)
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	contains := func(slice []int, item int) bool {
		for _, value := range slice {
			if value == item {
				return true
			}
		}
		return false
	}

	sum := 0
	for _, update := range updates {
		valid := true
		for pageIndex, page := range update {
			// check values to the left
			if pageIndex > 0 {
				prevPages := update[:pageIndex-1]
				for _, prevPage := range prevPages {
					if contains(rulesLeftMap[page], prevPage) {
						valid = false
						break
					}
				}
			}
			// check values to the right
			if pageIndex < len(update)-1 {
				nextPages := update[pageIndex+1:]
				for _, nextPage := range nextPages {
					if contains(rulesRightMap[page], nextPage) {
						valid = false
						break
					}
				}
			}
		}

		if !valid {
			continue
		}

		middleValue := update[len(update)/2]
		sum += middleValue
	}

	return sum
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	rulesLeftMap := make(map[int][]int)
	rulesRightMap := make(map[int][]int)
	updates := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// stop scanning rules to scan updates
		text := scanner.Text()
		if text == "" {
			break
		}

		substrings := strings.SplitN(text, "|", 2)
		leftValue, _ := strconv.Atoi(substrings[0])
		rightValue, _ := strconv.Atoi(substrings[1])

		rulesLeftMap[leftValue] = append(rulesLeftMap[leftValue], rightValue)
		rulesRightMap[rightValue] = append(rulesRightMap[rightValue], leftValue)
	}
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, ",")
		update := []int{}
		for _, substring := range substrings {
			page, _ := strconv.Atoi(substring)
			update = append(update, page)
		}
		updates = append(updates, update)
	}

	contains := func(slice []int, item int) bool {
		for _, value := range slice {
			if value == item {
				return true
			}
		}
		return false
	}

	getInvalidIndex := func(update []int) int {
		index := -1
		for pageIndex, page := range update {
			// check values to the left
			if pageIndex > 0 {
				prevPages := update[:pageIndex-1]
				for _, prevPage := range prevPages {
					if contains(rulesLeftMap[page], prevPage) {
						index = pageIndex
						break
					}
				}
			}
			// check values to the right
			if pageIndex < len(update)-1 {
				nextPages := update[pageIndex+1:]
				for _, nextPage := range nextPages {
					if contains(rulesRightMap[page], nextPage) {
						index = pageIndex
						break
					}
				}
			}
		}
		return index
	}

	sum := 0
	for _, update := range updates {
		// if it's valid, skip
		if getInvalidIndex(update) < 0 {
			continue
		}

		// if it's not valid, try moving the invalid index
		result := append([]int{}, update...)
		for {
			invalidIndex := getInvalidIndex(result)
			if invalidIndex < 0 {
				break
			}

			page := result[invalidIndex]

			result = append(append([]int{}, result[:invalidIndex]...), result[invalidIndex+1:]...)
			for index := range len(result) {
				// left most, shift
				if index == 0 && contains(rulesLeftMap[page], result[index]) {
					result = append([]int{page}, result...)
					break
				}
				// right most, append
				if index == len(result)-1 && contains(rulesRightMap[page], result[index]) {
					result = append(result, page)
					break
				}
				// middle, splice
				if contains(rulesRightMap[page], result[index]) && contains(rulesLeftMap[page], result[index+1]) {
					result = append(append(append([]int{}, result[:index+1]...), page), result[index+1:]...)
					break
				}
			}
		}

		middleValue := result[len(result)/2]
		sum += middleValue
	}

	return sum
}
