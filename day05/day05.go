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
			if pageIndex > 0 {
				prevPages := update[:pageIndex-1]
				for _, prevPage := range prevPages {
					if contains(rulesLeftMap[page], prevPage) {
						valid = false
						break
					}
				}
			}

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

		if valid {
			middleValue := update[len(update)/2]
			sum += middleValue
		}
	}

	return sum
}

// TODO:
func SecondHalf() int {
	return 0
}
