package day16

import (
	"bufio"
	"os"
	"strings"
)

func FirstHalf() int {
	type path struct {
		moves     int
		rotations int
	}

	type coord struct {
		row int
		col int
	}

	const (
		startId = "S"
		endId   = "E"
		wallId  = "#"
		freeId  = "."

		up    = 0
		right = 1
		down  = 2
		left  = 3
	)

	directions := []int{up, right, down, left}
	stepMap := map[int]coord{up: {-1, 0}, right: {0, 1}, down: {1, 0}, left: {0, -1}}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	start := coord{}
	end := coord{}
	wallMap := map[coord]bool{}
	freeMap := map[coord]bool{}
	rowCount := 0
	colCount := 0

	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")

		for index, substring := range substrings {
			currCoord := coord{rowCount, index}
			if substring == startId {
				start = currCoord
				freeMap[currCoord] = false
			}
			if substring == endId {
				end = currCoord
				freeMap[currCoord] = true
			}
			if substring == wallId {
				wallMap[currCoord] = true
				freeMap[currCoord] = false
			}
			if substring == freeId {
				freeMap[currCoord] = true
			}
		}

		rowCount++
		if colCount <= 0 {
			colCount = len(substrings)
		}
	}

	minCost := make(map[coord]path)
	var getPaths func(currCoord coord, direction int, currPath path, validPaths *[]path, visitedMap map[coord]int)
	getPaths = func(currCoord coord, direction int, currPath path, validPaths *[]path, visitedMap map[coord]int) {
		// base case
		if currCoord == end {
			*validPaths = append(*validPaths, currPath)
			return
		}

		for _, d := range directions {
			step := stepMap[d]
			nextCoord := coord{currCoord.row + step.row, currCoord.col + step.col}

			// skip if it isn't cheaper
			prevCost, exists := minCost[nextCoord]
			if exists && (currPath.moves >= prevCost.moves && currPath.rotations >= prevCost.rotations) {
				continue
			}
			minCost[nextCoord] = path{currPath.moves, currPath.rotations}

			// skip if it's invalid
			if isFree, exists := freeMap[nextCoord]; !exists || !isFree {
				continue
			}
			if _, exists := visitedMap[nextCoord]; exists {
				continue
			}

			// make move
			visitedMap[nextCoord] = d
			currPath.moves++
			if d != direction {
				currPath.rotations++
			}

			// recurse
			getPaths(nextCoord, d, currPath, validPaths, visitedMap)

			// backtrack
			delete(visitedMap, nextCoord)
			currPath.moves--
			if d != direction {
				currPath.rotations--
			}
		}
	}

	validPaths := []path{}
	getPaths(start, right, path{}, &validPaths, map[coord]int{})

	score := -1
	for _, path := range validPaths {
		currentScore := 1000*path.rotations + path.moves

		if score < 0 || currentScore < score {
			score = currentScore
		}
	}

	return score
}

// TODO:
func SecondHalf() int {
	return -1
}
