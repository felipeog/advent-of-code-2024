package day06

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func FirstHalf() int {
	const (
		up    = "up"
		right = "right"
		down  = "down"
		left  = "left"
	)
	directions := []string{up, right, down, left}
	stepMap := map[string][]int{
		up:    {-1, 0}, // up
		right: {0, 1},  // right
		down:  {1, 0},  // down
		left:  {0, -1}, // left
	}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	position := []int{0, 0}
	direction := up

	positionsMap := make(map[string]bool)
	obstructionsMap := make(map[string]bool)

	rowCount := 0
	colCount := 0

	getPosition := func(row, col int) string {
		return fmt.Sprintf("%d,%d", row, col)
	}

	getNextDirection := func(direction string) string {
		index := 0
		for i, d := range directions {
			if d == direction {
				index = i
				break
			}
		}

		newIndex := index + 1
		if newIndex > len(directions)-1 {
			return up
		}
		return directions[newIndex]
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.SplitN(text, "", -1)

		for index, substring := range substrings {
			if substring == "#" {
				obstructionsMap[getPosition(rowCount, index)] = true
			}
			if substring == "^" {
				position = []int{rowCount, index}
				positionsMap[getPosition(rowCount, index)] = true
			}
		}

		rowCount++
		colCount = int(math.Max(float64(colCount), float64(len(substrings))))
	}

	for {
		stop := false
		for _, d := range directions {
			if d != direction {
				continue
			}

			step := stepMap[d]
			newRow := position[0] + step[0]
			newCol := position[1] + step[1]

			inside := newRow >= 0 && newRow < rowCount && newCol >= 0 && newCol < colCount
			if !inside {
				stop = true
				break
			}

			key := getPosition(newRow, newCol)
			if _, exists := obstructionsMap[key]; exists {
				direction = getNextDirection(direction)
				break
			}

			positionsMap[key] = true
			position[0] = newRow
			position[1] = newCol
		}

		if stop {
			break
		}
	}

	return len(positionsMap)
}

// TODO:
func SecondHalf() int {
	return 0
}
