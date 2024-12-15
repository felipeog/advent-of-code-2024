package day14

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func FirstHalf() int {
	const (
		// width   = 11
		// height  = 7
		width   = 101
		height  = 103
		seconds = 100
	)

	digitsRegex := regexp.MustCompile(`-?\d+`)

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var px, py, vx, vy int
	var q1, q2, q3, q4 int

	for scanner.Scan() {
		matches := digitsRegex.FindAllString(scanner.Text(), 4)
		px, _ = strconv.Atoi(matches[0])
		py, _ = strconv.Atoi(matches[1])
		vx, _ = strconv.Atoi(matches[2])
		vy, _ = strconv.Atoi(matches[3])

		px = (px + vx*seconds) % width
		if px < 0 {
			px += width
		}

		py = (py + vy*seconds) % height
		if py < 0 {
			py += height
		}

		if px < width/2 && py < height/2 {
			q1++
		}
		if px < width/2 && py > height/2 {
			q2++
		}
		if px > width/2 && py > height/2 {
			q3++
		}
		if px > width/2 && py < height/2 {
			q4++
		}
	}

	return q1 * q2 * q3 * q4
}

// TODO:
func SecondHalf() int {
	return -1
}
