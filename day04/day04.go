package day04

import (
	"bufio"
	"os"
	"strings"
)

func FirstHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	matrix := [][]string{}
	xmas := "XMAS"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		matrix = append(matrix, substrings)
	}

	directions := []struct {
		rowStep, colStep int
	}{
		{-1, 0},  // up
		{-1, 1},  // up right
		{0, 1},   // right
		{1, 1},   // down right
		{1, 0},   // down
		{1, -1},  // down left
		{0, -1},  // left
		{-1, -1}, // up left
	}

	count := 0
	for rowIndex, row := range matrix {
		for colIndex := range row {
			if matrix[rowIndex][colIndex] != "X" {
				continue
			}

			for _, direction := range directions {
				valid := true
				for letterIndex := range len(xmas) {
					newRowIndex := rowIndex + letterIndex*direction.rowStep
					newColIndex := colIndex + letterIndex*direction.colStep
					inside := newRowIndex >= 0 && newRowIndex < len(matrix) && newColIndex >= 0 && newColIndex < len(matrix[0])

					if !inside || matrix[newRowIndex][newColIndex] != string(xmas[letterIndex]) {
						valid = false
						break
					}
				}

				if valid {
					count++
				}
			}
		}
	}

	return count
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	matrix := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		matrix = append(matrix, substrings)
	}

	isDiagonalValid := func(diagonal string) bool {
		validValues := []string{"MAS", "SAM"}
		for _, value := range validValues {
			if diagonal == value {
				return true
			}
		}
		return false
	}

	count := 0
	for rowIndex, row := range matrix {
		for colIndex := range row {
			letter := matrix[rowIndex][colIndex]
			diagonal1 := ""
			diagonal2 := ""

			if letter != "A" {
				continue
			}

			inside := rowIndex-1 >= 0 && rowIndex+1 < len(matrix) && colIndex-1 >= 0 && colIndex+1 < len(matrix[0])
			if inside {
				diagonal1 = matrix[rowIndex-1][colIndex-1] + letter + matrix[rowIndex+1][colIndex+1]
				diagonal2 = matrix[rowIndex-1][colIndex+1] + letter + matrix[rowIndex+1][colIndex-1]

				if isDiagonalValid(diagonal1) && isDiagonalValid(diagonal2) {
					count++
				}
			}
		}
	}

	return count
}
