package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func FirstHalf() int {
	type coord struct {
		row int
		col int
	}

	const (
		up    = "up"
		right = "right"
		down  = "down"
		left  = "left"
	)

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	directions := []string{up, right, down, left}
	stepMap := map[string]coord{
		up:    {-1, 0},
		right: {0, 1},
		down:  {1, 0},
		left:  {0, -1},
	}
	matrix := [][]int{}
	trailheads := []coord{}
	sum := 0

	var hike func([][]int, coord, map[coord]bool)
	hike = func(matrix [][]int, start coord, trailtailMap map[coord]bool) {
		if matrix[start.row][start.col] == 9 {
			trailtailMap[coord{start.row, start.col}] = true

			return
		}

		for _, d := range directions {
			step := stepMap[d]
			newRow := start.row + step.row
			newCol := start.col + step.col
			inside := newRow >= 0 && newRow < len(matrix) && newCol >= 0 && newCol < len(matrix[0])

			if !inside || matrix[newRow][newCol] != matrix[start.row][start.col]+1 {
				continue
			}

			hike(matrix, coord{newRow, newCol}, trailtailMap)
		}
	}

	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		values := []int{}

		for index, substring := range substrings {
			value, _ := strconv.Atoi(substring)
			values = append(values, value)

			if value == 0 {
				trailheads = append(trailheads, coord{len(matrix), index})
			}
		}

		matrix = append(matrix, values)
	}

	for _, trailhead := range trailheads {
		trailtailMap := map[coord]bool{}

		hike(matrix, trailhead, trailtailMap)

		sum += len(trailtailMap)
	}

	return sum
}

func SecondHalf() int {
	type coord struct {
		row int
		col int
	}

	const (
		up    = "up"
		right = "right"
		down  = "down"
		left  = "left"
	)

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	directions := []string{up, right, down, left}
	stepMap := map[string]coord{
		up:    {-1, 0},
		right: {0, 1},
		down:  {1, 0},
		left:  {0, -1},
	}
	matrix := [][]int{}
	trailheads := []coord{}
	sum := 0

	var hike func([][]int, coord, *int)
	hike = func(matrix [][]int, start coord, rating *int) {
		if matrix[start.row][start.col] == 9 {
			*rating++

			return
		}

		for _, d := range directions {
			step := stepMap[d]
			newRow := start.row + step.row
			newCol := start.col + step.col
			inside := newRow >= 0 && newRow < len(matrix) && newCol >= 0 && newCol < len(matrix[0])

			if !inside || matrix[newRow][newCol] != matrix[start.row][start.col]+1 {
				continue
			}

			hike(matrix, coord{newRow, newCol}, rating)
		}
	}

	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		values := []int{}

		for index, substring := range substrings {
			value, _ := strconv.Atoi(substring)
			values = append(values, value)

			if value == 0 {
				trailheads = append(trailheads, coord{len(matrix), index})
			}
		}

		matrix = append(matrix, values)
	}

	for _, trailhead := range trailheads {
		rating := 0

		hike(matrix, trailhead, &rating)

		sum += rating
	}

	return sum
}
