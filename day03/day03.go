package day03

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func FirstHalf() int {
	// file, _ := os.Open("sample1.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	operationsRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	digitsRegex := regexp.MustCompile(`\d+`)

	operations := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		operations = append(operations, operationsRegex.FindAllString(text, -1)...)
	}

	sum := 0
	for _, operation := range operations {
		matches := digitsRegex.FindAllString(operation, 2)
		leftValue, _ := strconv.Atoi(matches[0])
		rightValue, _ := strconv.Atoi(matches[1])

		sum += leftValue * rightValue
	}

	return sum
}

func SecondHalf() int {
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	operationsRegex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	digitsRegex := regexp.MustCompile(`\d+`)
	doRegex := regexp.MustCompile(`do\(\)`)
	dontRegex := regexp.MustCompile(`don't\(\)`)

	operations := []string{}
	enabled := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		instruction := ""
		currentIndex := 0

		for {
			// do(), store enabled instruction
			if enabled {
				matchIndices := dontRegex.FindStringIndex(text[currentIndex:])
				if matchIndices != nil {
					instruction += text[currentIndex : currentIndex+matchIndices[0]]
					currentIndex += matchIndices[1]
					enabled = false
					continue
				}

				instruction += text[currentIndex:]
				break
			}

			// don't(), update current index
			matchIndices := doRegex.FindStringIndex(text[currentIndex:])
			if matchIndices != nil {
				currentIndex += matchIndices[1]
				enabled = true
				continue
			}

			break
		}

		operations = append(operations, operationsRegex.FindAllString(instruction, -1)...)
	}

	sum := 0
	for _, operation := range operations {
		matches := digitsRegex.FindAllString(operation, 2)
		leftValue, _ := strconv.Atoi(matches[0])
		rightValue, _ := strconv.Atoi(matches[1])

		sum += leftValue * rightValue
	}

	return sum
}
