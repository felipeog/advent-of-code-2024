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

	scanner := bufio.NewScanner(file)

	operationsRegex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	digitsRegex, _ := regexp.Compile(`\d+`)

	operations := []string{}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		operations = append(operations, operationsRegex.FindAllString(line, -1)...)
	}

	for _, operation := range operations {
		digits := digitsRegex.FindAllString(operation, 2)
		leftDigit, _ := strconv.Atoi(digits[0])
		rightDigit, _ := strconv.Atoi(digits[1])

		sum = sum + leftDigit*rightDigit
	}

	return sum
}

func SecondHalf() int {
	// file, _ := os.Open("sample2.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	operationsRegex, _ := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	digitsRegex, _ := regexp.Compile(`\d+`)
	doRegex, _ := regexp.Compile(`do\(\)`)
	dontRegex, _ := regexp.Compile(`don't\(\)`)

	enabled := true
	operations := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		instruction := ""
		currentIndex := 0

		for {
			if enabled {
				matchIndices := dontRegex.FindStringIndex(line[currentIndex:])
				if matchIndices != nil {
					instruction = instruction + line[currentIndex:currentIndex+matchIndices[0]]
					currentIndex = currentIndex + matchIndices[1]
					enabled = false
					continue
				} else {
					instruction = instruction + line[currentIndex:]
				}
			} else {
				matchIndices := doRegex.FindStringIndex(line[currentIndex:])
				if matchIndices != nil {
					currentIndex = currentIndex + matchIndices[1]
					enabled = true
					continue
				}
			}

			break
		}

		operations = append(operations, operationsRegex.FindAllString(instruction, -1)...)
	}

	sum := 0
	for _, operation := range operations {
		digits := digitsRegex.FindAllString(operation, 2)
		leftDigit, _ := strconv.Atoi(digits[0])
		rightDigit, _ := strconv.Atoi(digits[1])

		sum = sum + leftDigit*rightDigit
	}

	return sum
}
