package day08

import (
	"bufio"
	"os"
	"regexp"
)

func FirstHalf() int {
	type coord struct {
		row int
		col int
	}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	antennaRegex := regexp.MustCompile("[0-9a-zA-Z]")
	antennaMap := map[string][]coord{}
	antinodeMap := map[coord]bool{}
	rowCount := 0
	colCount := 0

	for scanner.Scan() {
		text := scanner.Text()
		positions := antennaRegex.FindAllStringIndex(text, -1)

		for _, position := range positions {
			antenna := text[position[0]:position[1]]
			antennaMap[antenna] = append(antennaMap[antenna], coord{rowCount, position[0]})
		}

		if rowCount == 0 {
			colCount = len(text)
		}

		rowCount++
	}

	for _, coords := range antennaMap {
		for i := range coords {
			from := coords[i]

			for j := range coords {
				if j == i {
					continue
				}

				to := coords[j]
				distance := coord{to.row - from.row, to.col - from.col}
				antinode := coord{to.row + distance.row, to.col + distance.col}

				if antinode.row < 0 || antinode.row >= rowCount || antinode.col < 0 || antinode.col >= colCount {
					continue
				}

				antinodeMap[antinode] = true
			}
		}
	}

	return len(antinodeMap)
}

// TODO:
func SecondHalf() int {
	return -1
}
