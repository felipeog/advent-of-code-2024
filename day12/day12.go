package day12

import (
	"bufio"
	"os"
	"strings"
)

func FirstHalf() int {
	type coord struct {
		row     int
		col     int
		value   string
		visited bool
	}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	steps := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var traverse func([][]coord, coord, *int, *int)
	traverse = func(matrix [][]coord, start coord, area, perimeter *int) {
		if matrix[start.row][start.col].visited {
			return
		}
		matrix[start.row][start.col].visited = true

		curr := matrix[start.row][start.col]
		*area++

		for _, step := range steps {
			newRow := curr.row + step[0]
			newCol := curr.col + step[1]

			if newRow < 0 || newRow >= len(matrix) || newCol < 0 || newCol >= len(matrix[0]) {
				*perimeter++
				continue
			}

			next := matrix[newRow][newCol]
			if next.value != curr.value {
				*perimeter++
				continue
			}

			traverse(matrix, next, area, perimeter)
		}
	}

	matrix := [][]coord{}
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		values := []coord{}

		for index, substring := range substrings {
			values = append(values, coord{len(matrix), index, substring, false})
		}

		matrix = append(matrix, values)
	}

	sum := 0
	for rowIndex, row := range matrix {
		for colIndex := range row {
			area := 0
			perimeter := 0

			traverse(matrix, matrix[rowIndex][colIndex], &area, &perimeter)

			sum += area * perimeter
		}
	}

	return sum
}

// TODO:
func SecondHalf() int {
	return 0
}
