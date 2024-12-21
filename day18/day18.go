package day18

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func FirstHalf() int {
	type coord struct {
		row int
		col int
	}

	type path struct {
		cost       int
		visitedMap map[coord]int
	}

	type cost struct {
		moves int
	}

	const (
		up    = 0
		right = 1
		down  = 2
		left  = 3
	)

	stepMap := map[int]coord{up: {-1, 0}, right: {0, 1}, down: {1, 0}, left: {0, -1}}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	corruptedMap := map[coord]bool{}
	rowCount := 71
	colCount := 71
	start := coord{0, 0}
	end := coord{rowCount - 1, colCount - 1}

	for i := 0; i < 1024; i++ {
		scanner.Scan()
		text := scanner.Text()
		substrings := strings.Split(text, ",")

		row, _ := strconv.Atoi(substrings[1])
		col, _ := strconv.Atoi(substrings[0])
		currCoord := coord{row, col}

		corruptedMap[currCoord] = true
	}

	minMoves := math.MaxInt
	minCost := make(map[coord]cost)
	visitedMap := map[coord]int{}
	pathCost := cost{0}
	validPaths := []path{}

	var getPaths func(currCoord coord, direction int)
	getPaths = func(currCoord coord, direction int) {
		// base case
		if currCoord == end {
			if pathCost.moves < minMoves {
				minMoves = pathCost.moves
			}

			visitedMapCopy := make(map[coord]int)
			for key, value := range visitedMap {
				visitedMapCopy[key] = value
			}

			validPaths = append(validPaths, path{
				cost:       pathCost.moves,
				visitedMap: visitedMapCopy,
			})

			return
		}

		// skip if it isn't cheaper
		if pathCost.moves >= minMoves {
			return
		}

		from := direction - 1
		if from < 0 {
			from += 4
		}
		to := direction + 1
		if to > 3 {
			to -= 4
		}
		if from > to {
			from, to = to, from
		}

		nextDirections := []int{direction, from, to}
		for _, nextDirection := range nextDirections {
			step := stepMap[nextDirection]
			nextCoord := coord{currCoord.row + step.row, currCoord.col + step.col}

			// skip if it isn't cheaper
			if prevCost, exists := minCost[nextCoord]; exists && pathCost.moves >= prevCost.moves {
				continue
			}
			minCost[nextCoord] = cost{pathCost.moves}

			// skip if it's invalid
			if nextCoord.row < 0 || nextCoord.row >= rowCount || nextCoord.col < 0 || nextCoord.col >= colCount {
				continue
			}
			if _, exists := corruptedMap[nextCoord]; exists {
				continue
			}
			if _, exists := visitedMap[nextCoord]; exists {
				continue
			}

			// make move
			visitedMap[nextCoord] = nextDirection
			pathCost.moves++

			// recurse
			getPaths(nextCoord, nextDirection)

			// backtrack
			delete(visitedMap, nextCoord)
			pathCost.moves--
		}
	}

	getPaths(start, down)

	return minMoves
}

// TODO:
func SecondHalf() int {
	return -1
}
