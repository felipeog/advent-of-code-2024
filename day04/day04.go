package day04

import (
	"bufio"
	"os"
	"strings"
)

func FirstHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	var matrix [][]string
	xmas := "XMAS"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.SplitN(text, "", -1)
		matrix = append(matrix, substrings)
	}

	count := 0
	for rowIndex, row := range matrix {
		for colIndex := range row {
			letter := matrix[rowIndex][colIndex]
			var word string

			if letter != "X" {
				continue
			}

			hasTop := rowIndex-3 >= 0
			hasRight := colIndex+3 < len(matrix[0])
			hasBottom := rowIndex+3 < len(matrix)
			hasLeft := colIndex-3 >= 0

			// TODO: improve this, too many `if`s
			if hasTop {
				word = letter + matrix[rowIndex-1][colIndex] + matrix[rowIndex-2][colIndex] + matrix[rowIndex-3][colIndex]
				if word == xmas {
					count += 1
				}
			}

			if hasTop && hasRight {
				word = letter + matrix[rowIndex-1][colIndex+1] + matrix[rowIndex-2][colIndex+2] + matrix[rowIndex-3][colIndex+3]
				if word == xmas {
					count += 1
				}
			}

			if hasRight {
				word = letter + matrix[rowIndex][colIndex+1] + matrix[rowIndex][colIndex+2] + matrix[rowIndex][colIndex+3]
				if word == xmas {
					count += 1
				}
			}

			if hasBottom && hasRight {
				word = letter + matrix[rowIndex+1][colIndex+1] + matrix[rowIndex+2][colIndex+2] + matrix[rowIndex+3][colIndex+3]
				if word == xmas {
					count += 1
				}
			}

			if hasBottom {
				word = letter + matrix[rowIndex+1][colIndex] + matrix[rowIndex+2][colIndex] + matrix[rowIndex+3][colIndex]
				if word == xmas {
					count += 1
				}
			}

			if hasBottom && hasLeft {
				word = letter + matrix[rowIndex+1][colIndex-1] + matrix[rowIndex+2][colIndex-2] + matrix[rowIndex+3][colIndex-3]
				if word == xmas {
					count += 1
				}
			}

			if hasLeft {
				word = letter + matrix[rowIndex][colIndex-1] + matrix[rowIndex][colIndex-2] + matrix[rowIndex][colIndex-3]
				if word == xmas {
					count += 1
				}
			}

			if hasTop && hasLeft {
				word = letter + matrix[rowIndex-1][colIndex-1] + matrix[rowIndex-2][colIndex-2] + matrix[rowIndex-3][colIndex-3]
				if word == xmas {
					count += 1
				}
			}
		}
	}

	return count
}

func SecondHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	defer file.Close()

	var matrix [][]string
	mas := "MAS"
	sam := "SAM"

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		substrings := strings.SplitN(text, "", -1)
		matrix = append(matrix, substrings)
	}

	count := 0
	for rowIndex, row := range matrix {
		for colIndex := range row {
			letter := matrix[rowIndex][colIndex]
			var diagonal1 string
			var diagonal2 string

			if letter != "A" {
				continue
			}

			hasTop := rowIndex-1 >= 0
			hasRight := colIndex+1 < len(matrix[0])
			hasBottom := rowIndex+1 < len(matrix)
			hasLeft := colIndex-1 >= 0

			if hasTop && hasRight && hasBottom && hasLeft {
				diagonal1 = matrix[rowIndex-1][colIndex-1] + letter + matrix[rowIndex+1][colIndex+1]
				diagonal2 = matrix[rowIndex-1][colIndex+1] + letter + matrix[rowIndex+1][colIndex-1]

				// TODO: improve string matching
				if (diagonal1 == sam || diagonal1 == mas) && (diagonal2 == sam || diagonal2 == mas) {
					count += 1
				}
			}
		}
	}

	return count
}
