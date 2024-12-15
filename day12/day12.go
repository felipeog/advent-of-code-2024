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

func SecondHalf() int {
	type coord struct {
		row int
		col int
	}

	type areaCoord struct {
		row     int
		col     int
		value   string
		visited bool
	}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	steps := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var traverse func([][]areaCoord, areaCoord, *int, map[coord]bool)
	traverse = func(matrix [][]areaCoord, start areaCoord, area *int, perimeterMap map[coord]bool) {
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
				continue
			}

			perimeterMap[coord{curr.row, curr.col}] = true
			next := matrix[newRow][newCol]
			if next.value != curr.value {
				continue
			}

			traverse(matrix, next, area, perimeterMap)
		}
	}

	countAxisSides := func(matrix [][]areaCoord, axis string, perimeterMap map[coord]bool, sides *int) {
		for col := 0; col < len(matrix[0]); col++ {
			for row := 0; row < len(matrix); row++ {
				_, curr := perimeterMap[coord{row, col}]
				var right bool
				var bottomRight bool
				var bottom bool
				var bottomLeft bool
				var left bool

				if axis == "vertical" {
					_, right = perimeterMap[coord{row, col + 1}]
					_, bottomRight = perimeterMap[coord{row + 1, col + 1}]
					_, bottom = perimeterMap[coord{row + 1, col}]
					_, bottomLeft = perimeterMap[coord{row + 1, col - 1}]
					_, left = perimeterMap[coord{row, col - 1}]
				}

				if axis == "horizontal" {
					_, right = perimeterMap[coord{row - 1, col}]
					_, bottomRight = perimeterMap[coord{row - 1, col + 1}]
					_, bottom = perimeterMap[coord{row, col + 1}]
					_, bottomLeft = perimeterMap[coord{row + 1, col + 1}]
					_, left = perimeterMap[coord{row + 1, col}]
				}

				if !curr {
					continue
				}

				// check zero
				if left && curr && right {
					continue
				}
				if !bottomLeft && bottom && !bottomRight {
					continue
				}
				if left && curr && !right && bottomLeft && bottom && !bottomRight {
					continue
				}
				if !left && curr && right && !bottomLeft && bottom && bottomRight {
					continue
				}

				// check one
				if left && curr && !right {
					if bottom && !bottomRight {
						continue
					}
					*sides++
					continue
				}

				if !left && curr && right {
					if bottom && !bottomLeft {
						continue
					}
					*sides++
					continue
				}
				if !left && curr && !right {
					if bottomLeft && bottom && !bottomRight {
						*sides++
						continue
					}
					if !bottomLeft && bottom && bottomRight {
						*sides++
						continue
					}
				}

				// check two
				if !left && curr && !right {
					*sides += 2
					continue
				}
			}
		}
	}

	matrix := [][]areaCoord{}
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		values := []areaCoord{}

		for index, substring := range substrings {
			values = append(values, areaCoord{len(matrix), index, substring, false})
		}

		matrix = append(matrix, values)
	}

	sum := 0
	for rowIndex, row := range matrix {
		for colIndex := range row {
			area := 0
			perimeterMap := map[coord]bool{}
			traverse(matrix, matrix[rowIndex][colIndex], &area, perimeterMap)

			vertical := 0
			horizontal := 0
			countAxisSides(matrix, "vertical", perimeterMap, &vertical)
			countAxisSides(matrix, "horizontal", perimeterMap, &horizontal)

			sides := vertical + horizontal
			sum += area * sides
		}
	}

	return sum
}
