package day15

import (
	"bufio"
	"os"
	"strings"
)

func FirstHalf() int {
	type coord struct {
		row int
		col int
	}

	const (
		robotId = "@"
		wallId  = "#"
		boxId   = "O"
		freeId  = "."

		up    = "^"
		right = ">"
		down  = "v"
		left  = "<"
	)

	directions := []string{up, right, down, left}
	stepMap := map[string]coord{up: {-1, 0}, right: {0, 1}, down: {1, 0}, left: {0, -1}}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	robot := coord{}
	wallMap := map[coord]bool{}
	boxMap := map[coord]bool{}
	freeMap := map[coord]bool{}
	moves := []string{}

	rowCount := 0
	colCount := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		substrings := strings.Split(text, "")
		for index, substring := range substrings {
			currCoord := coord{rowCount, index}
			if substring == robotId {
				robot = currCoord
				freeMap[currCoord] = false
			}
			if substring == wallId {
				wallMap[currCoord] = true
			}
			if substring == boxId {
				boxMap[currCoord] = true
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

	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		for _, substring := range substrings {
			moves = append(moves, substring)
		}
	}

	for _, move := range moves {
		for _, direction := range directions {
			if move != direction {
				continue
			}

			nextCoord := coord{robot.row + stepMap[move].row, robot.col + stepMap[move].col}
			if isWall := wallMap[nextCoord]; isWall {
				break
			}
			if isFree := freeMap[nextCoord]; isFree {
				freeMap[robot] = true
				freeMap[nextCoord] = false
				robot = nextCoord
				break
			}

			nextFreeCoord := nextCoord
			for {
				nextFreeCoord = coord{nextFreeCoord.row + stepMap[move].row, nextFreeCoord.col + stepMap[move].col}
				if isWall := wallMap[nextFreeCoord]; isWall {
					break
				}
				if isFree := freeMap[nextFreeCoord]; !isFree {
					continue
				}

				// swap next (now box) and next free
				boxMap[nextCoord] = false
				boxMap[nextFreeCoord] = true
				freeMap[nextFreeCoord] = false

				// swap robot and next (now free)
				freeMap[robot] = true
				freeMap[nextCoord] = false
				robot = nextCoord

				break
			}
		}
	}

	sum := 0
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			currCoord := coord{row, col}
			if isBox := boxMap[currCoord]; isBox {
				sum += 100*row + col
			}
		}
	}

	return sum
}

// TODO:
func SecondHalf() int {
	return 0
}
