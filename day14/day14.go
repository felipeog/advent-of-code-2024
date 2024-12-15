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

func SecondHalf() int {
	type coord struct {
		x int
		y int
	}

	type robot struct {
		px int
		py int
		vx int
		vy int
	}

	const (
		width  = 101
		height = 103
	)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	digitsRegex := regexp.MustCompile(`-?\d+`)
	robots := []robot{}
	px, py, vx, vy := 0, 0, 0, 0
	result := 0

	for scanner.Scan() {
		matches := digitsRegex.FindAllString(scanner.Text(), 4)

		px, _ = strconv.Atoi(matches[0])
		py, _ = strconv.Atoi(matches[1])
		vx, _ = strconv.Atoi(matches[2])
		vy, _ = strconv.Atoi(matches[3])

		robots = append(robots, robot{px, py, vx, vy})
	}

	for seconds := 0; seconds < width*height; seconds++ {
		positionMap := map[coord]int{}
		skip := false

		for _, robot := range robots {
			robot.px = (robot.px + robot.vx*seconds) % width
			robot.py = (robot.py + robot.vy*seconds) % height

			if robot.px < 0 {
				robot.px += width
			}

			if robot.py < 0 {
				robot.py += height
			}

			// skip when an overlap is found
			if _, exists := positionMap[coord{robot.px, robot.py}]; exists {
				skip = true
				break
			}

			positionMap[coord{robot.px, robot.py}]++
		}

		if skip {
			continue
		}

		result = seconds
	}

	return result
}
