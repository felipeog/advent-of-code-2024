package day15

import (
	"bufio"
	"os"
	"slices"
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
		moves = append(moves, substrings...)
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

func SecondHalf() int {
	type coord struct {
		row int
		col int
	}

	const (
		robotId    = "@"
		wallId     = "#"
		boxId      = "O"
		boxLeftId  = "["
		boxRightId = "]"
		freeId     = "."

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
	boxMap := map[coord]string{}
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
		for index := 0; index < len(substrings); index++ {
			substring := substrings[index]

			currCoord := coord{rowCount, 2 * index}
			rightCoord := coord{rowCount, 2*index + 1}

			if substring == robotId {
				robot = currCoord
				freeMap[currCoord] = false
				freeMap[rightCoord] = true
			}
			if substring == wallId {
				wallMap[currCoord] = true
				wallMap[rightCoord] = true
			}
			if substring == boxId {
				boxMap[currCoord] = left
				boxMap[rightCoord] = right
				freeMap[currCoord] = false
				freeMap[rightCoord] = false
			}
			if substring == freeId {
				freeMap[currCoord] = true
				freeMap[rightCoord] = true
			}
		}

		rowCount++
		if colCount <= 0 {
			colCount = 2 * len(substrings)
		}
	}

	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.Split(text, "")
		moves = append(moves, substrings...)
	}

	var getBoxesToMove func(currCoord coord, direction string, wallHit *bool, boxesToMove map[coord]bool)
	getBoxesToMove = func(currCoord coord, direction string, wallHit *bool, boxesToMove map[coord]bool) {
		step := stepMap[direction]

		if isWall := wallMap[currCoord]; isWall {
			*wallHit = true
			return
		}
		if isFree := freeMap[currCoord]; isFree {
			return
		}

		if direction == left || direction == right {
			nextCoord := coord{currCoord.row + step.row, currCoord.col + step.col}
			boxesToMove[currCoord] = true
			getBoxesToMove(nextCoord, direction, wallHit, boxesToMove)
			return
		}
		if direction == up || direction == down {
			var currSideCoord = coord{}
			if boxMap[currCoord] == left {
				currSideCoord = coord{currCoord.row, currCoord.col + 1}
			} else {
				currSideCoord = coord{currCoord.row, currCoord.col - 1}
			}

			nextCoord := coord{currCoord.row + step.row, currCoord.col + step.col}
			nextSideCoord := coord{currSideCoord.row + step.row, currSideCoord.col + step.col}
			boxesToMove[currCoord] = true
			boxesToMove[currSideCoord] = true
			getBoxesToMove(nextCoord, direction, wallHit, boxesToMove)
			getBoxesToMove(nextSideCoord, direction, wallHit, boxesToMove)
			return
		}
	}

	for _, move := range moves {
		for _, direction := range directions {
			if move != direction {
				continue
			}
			step := stepMap[move]

			nextCoord := coord{robot.row + step.row, robot.col + step.col}
			if isWall := wallMap[nextCoord]; isWall {
				break
			}
			if isFree := freeMap[nextCoord]; isFree {
				freeMap[robot] = true
				freeMap[nextCoord] = false
				robot = nextCoord
				break
			}

			for {
				wallHit := false
				boxesToMove := map[coord]bool{}
				getBoxesToMove(nextCoord, direction, &wallHit, boxesToMove)

				if wallHit {
					break
				}

				// sort boxes from furthest to closest
				boxCoords := []coord{}
				for boxCoord := range boxesToMove {
					boxCoords = append(boxCoords, boxCoord)
				}
				slices.SortFunc(boxCoords, func(a, b coord) int {
					if step.row < 0 {
						return a.row - b.row
					}
					if step.row > 0 {
						return b.row - a.row
					}
					if step.col < 0 {
						return a.col - b.col
					}
					if step.col > 0 {
						return b.col - a.col
					}
					return 0
				})

				for _, boxCoord := range boxCoords {
					toCoord := coord{boxCoord.row + step.row, boxCoord.col + step.col}

					// swap box and to (now free)
					boxMap[toCoord] = boxMap[boxCoord]
					delete(boxMap, boxCoord)
					freeMap[boxCoord] = true
					freeMap[toCoord] = false
				}

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
			if side, exists := boxMap[currCoord]; exists && side == left {
				sum += 100*row + col
			}
		}
	}

	return sum
}
