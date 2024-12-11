package day09

import (
	"bufio"
	"os"
)

func FirstHalf() int {
	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	blocks := []int{}
	id := 0
	sum := 0

	for index, currRune := range text {
		value := int(currRune - '0')

		for range value {
			if index%2 == 0 {
				blocks = append(blocks, id)
				continue
			}

			blocks = append(blocks, -1)
		}

		if index%2 == 0 {
			id++
		}
	}

	for index := 0; index < len(blocks); {
		currValue := blocks[index]
		lastValue := blocks[len(blocks)-1]

		if currValue >= 0 {
			sum += index * currValue
			index++
			continue
		}

		if lastValue < 0 {
			blocks = append([]int{}, blocks[:len(blocks)-1]...)
			continue
		}

		sum += index * lastValue
		blocks = append(append(append([]int{}, blocks[:index]...), lastValue), blocks[index+1:len(blocks)-1]...)
		index++
	}

	return sum
}

// TODO:
func SecondHalf() int {
	return -1
}
