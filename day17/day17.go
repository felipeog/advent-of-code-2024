package day17

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func FirstHalf() string {
	registerRegex := regexp.MustCompile(`Register`)
	programRegex := regexp.MustCompile(`Program`)
	digitsRegex := regexp.MustCompile(`\d+`)

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	register := []int{}
	program := []int{}
	output := []int{}
	stringOutput := []string{}
	pointer := 0

	for scanner.Scan() {
		text := scanner.Text()

		if registerRegex.MatchString(text) {
			match := digitsRegex.FindString(text)
			value, _ := strconv.Atoi(match)
			register = append(register, value)
		}

		if programRegex.MatchString(text) {
			matches := digitsRegex.FindAllString(text, -1)
			for _, match := range matches {
				value, _ := strconv.Atoi(match)
				program = append(program, value)
			}
		}
	}

	for {
		if pointer >= len(program)-1 {
			break
		}

		opcode := program[pointer]
		literalOperand := program[pointer+1]
		comboOperand := program[pointer+1]
		if comboOperand >= 4 && comboOperand <= 6 {
			comboOperand = register[comboOperand-4]
		}

		// adv
		if opcode == 0 {
			numerator := register[0]
			denominator := int(math.Pow(float64(2), float64(comboOperand)))
			register[0] = numerator / denominator
			pointer += 2
			continue
		}

		// bxl
		if opcode == 1 {
			register[1] = register[1] ^ literalOperand
			pointer += 2
			continue
		}

		// bst
		if opcode == 2 {
			register[1] = comboOperand % 8
			pointer += 2
			continue
		}

		// jnz
		if opcode == 3 {
			if register[0] <= 0 {
				pointer += 2
				continue
			}

			pointer = literalOperand
			continue
		}

		// bxc
		if opcode == 4 {
			register[1] = register[1] ^ register[2]
			pointer += 2
			continue
		}

		// out
		if opcode == 5 {
			output = append(output, comboOperand%8)
			pointer += 2
			continue
		}

		// bdv
		if opcode == 6 {
			numerator := register[0]
			denominator := int(math.Pow(float64(2), float64(comboOperand)))
			register[1] = numerator / denominator
			pointer += 2
			continue
		}

		// cdv
		if opcode == 7 {
			numerator := register[0]
			denominator := int(math.Pow(float64(2), float64(comboOperand)))
			register[2] = numerator / denominator
			pointer += 2
			continue
		}
	}

	for _, integer := range output {
		string := strconv.Itoa(integer)
		stringOutput = append(stringOutput, string)
	}

	return strings.Join(stringOutput, ",")
}

// TODO:
func SecondHalf() string {
	return "a"
}
