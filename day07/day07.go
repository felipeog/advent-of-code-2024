package day07

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FirstHalf() int {
	type entry struct {
		value   int
		numbers []int
	}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	entries := []entry{}
	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		valueAndNumbersSubstrings := strings.SplitN(text, ": ", 2)
		numbersSubstrings := strings.Fields(valueAndNumbersSubstrings[1])
		value, _ := strconv.Atoi(valueAndNumbersSubstrings[0])
		numbers := make([]int, len(numbersSubstrings))

		for index, substring := range numbersSubstrings {
			number, _ := strconv.Atoi(substring)
			numbers[index] = number
		}

		entries = append(entries, entry{value, numbers})
	}

	for _, entry := range entries {
		gaps := len(entry.numbers) - 1
		combinations := 1 << gaps
		finalResult := 0

		for i := 0; i < combinations; i++ {
			result := entry.numbers[0]

			for j := 0; j < gaps; j++ {
				if (i>>j)&1 == 1 {
					result *= entry.numbers[j+1]
					continue
				}

				result += entry.numbers[j+1]
			}

			if result == entry.value {
				finalResult = result
				break
			}
		}

		if finalResult > 0 {
			sum += finalResult
		}
	}

	return sum
}

// TODO:
func SecondHalf() int {
	return -1
}
