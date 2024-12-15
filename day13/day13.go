package day13

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func FirstHalf() int {
	type machine struct {
		ax     int
		ay     int
		bx     int
		by     int
		prizex int
		prizey int
	}

	digitsRegex := regexp.MustCompile(`\d+`)

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	machines := []machine{}
	lines := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		lines = append(lines, text)
		if len(lines) == 3 {
			aMatches := digitsRegex.FindAllString(lines[0], 2)
			bMatches := digitsRegex.FindAllString(lines[1], 2)
			prizeMatches := digitsRegex.FindAllString(lines[2], 2)

			ax, _ := strconv.Atoi(aMatches[0])
			ay, _ := strconv.Atoi(aMatches[1])
			bx, _ := strconv.Atoi(bMatches[0])
			by, _ := strconv.Atoi(bMatches[1])
			prizex, _ := strconv.Atoi(prizeMatches[0])
			prizey, _ := strconv.Atoi(prizeMatches[1])

			machines = append(machines, machine{ax, ay, bx, by, prizex, prizey})
			lines = []string{}
		}
	}

	sum := 0
	for _, machine := range machines {
		buttona := machine.by*machine.prizex - machine.bx*machine.prizey
		buttonb := -machine.ay*machine.prizex + machine.ax*machine.prizey
		determinant := (machine.ax * machine.by) - (machine.ay * machine.bx)

		if determinant == 0 || buttona%determinant != 0 || buttonb%determinant != 0 {
			continue
		}

		sum += 3*buttona/determinant + buttonb/determinant
	}

	return sum
}

func SecondHalf() int {
	type machine struct {
		ax     int
		ay     int
		bx     int
		by     int
		prizex int
		prizey int
	}

	digitsRegex := regexp.MustCompile(`\d+`)

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	machines := []machine{}
	lines := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		lines = append(lines, text)
		if len(lines) == 3 {
			aMatches := digitsRegex.FindAllString(lines[0], 2)
			bMatches := digitsRegex.FindAllString(lines[1], 2)
			prizeMatches := digitsRegex.FindAllString(lines[2], 2)

			ax, _ := strconv.Atoi(aMatches[0])
			ay, _ := strconv.Atoi(aMatches[1])
			bx, _ := strconv.Atoi(bMatches[0])
			by, _ := strconv.Atoi(bMatches[1])
			prizex, _ := strconv.Atoi(prizeMatches[0])
			prizey, _ := strconv.Atoi(prizeMatches[1])

			machines = append(machines, machine{
				ax, ay,
				bx, by,
				prizex + 10000000000000, prizey + 10000000000000,
			})
			lines = []string{}
		}
	}

	sum := 0
	for _, machine := range machines {
		buttona := machine.by*machine.prizex - machine.bx*machine.prizey
		buttonb := -machine.ay*machine.prizex + machine.ax*machine.prizey
		determinant := (machine.ax * machine.by) - (machine.ay * machine.bx)

		if determinant == 0 || buttona%determinant != 0 || buttonb%determinant != 0 {
			continue
		}

		sum += 3*buttona/determinant + buttonb/determinant
	}

	return sum
}
