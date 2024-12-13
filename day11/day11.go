package day11

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FirstHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	substrings := strings.Fields(text)
	stones := []int{}

	for _, substring := range substrings {
		stone, _ := strconv.Atoi(substring)
		stones = append(stones, stone)
	}

	for range 25 {
		newStones := []int{}

		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)

				continue
			}

			if stoneString := strconv.Itoa(stone); len(stoneString)%2 == 0 {
				left, _ := strconv.Atoi(stoneString[:len(stoneString)/2])
				right, _ := strconv.Atoi(stoneString[len(stoneString)/2:])
				newStones = append(newStones, []int{left, right}...)

				continue
			}

			newStones = append(newStones, stone*2024)
		}

		stones = append([]int{}, newStones...)
	}

	return len(stones)
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	file.Close()

	countMap := map[int]int{}

	substrings := strings.Fields(text)
	for _, substring := range substrings {
		value, _ := strconv.Atoi(substring)
		countMap[value]++
	}

	for i := 0; i < 75; i++ {
		nextCountMap := map[int]int{}

		for key, value := range countMap {
			if key == 0 {
				nextCountMap[1] += value
				continue
			}

			if stringValue := strconv.Itoa(key); len(stringValue)%2 == 0 {
				left, _ := strconv.Atoi(stringValue[:len(stringValue)/2])
				right, _ := strconv.Atoi(stringValue[len(stringValue)/2:])
				nextCountMap[left] += value
				nextCountMap[right] += value
				continue
			}

			nextCountMap[key*2024] += value
		}

		countMap = nextCountMap
	}

	sum := 0
	for _, value := range countMap {
		sum += value
	}

	return sum
}
