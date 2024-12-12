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

func SecondHalf() int {
	type block struct {
		id     int
		amount int
	}

	// file, _ := os.Open("sample.txt")
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	text := scanner.Text()
	blocks := []block{}
	id := 0
	sum := 0

	for index, currRune := range text {
		amount := int(currRune - '0')

		if amount <= 0 {
			continue
		}

		if index%2 == 0 {
			blocks = append(blocks, block{id, amount})
			id++
		} else {
			blocks = append(blocks, block{-1, amount})
		}
	}

	for index := id - 1; index >= 0; index-- {
		// get current index block
		currIndex := -1
		currBlock := block{}

		for blockIndex, block := range blocks {
			if block.id != index {
				continue
			}

			currIndex = blockIndex
			currBlock = block
			break
		}

		if currIndex < 0 {
			panic("currBlock not found")
		}

		// get first free block
		freeIndex := -1
		freeBlock := block{}

		for blockIndex := range currIndex {
			block := blocks[blockIndex]

			if block.id >= 0 || block.amount < currBlock.amount {
				continue
			}

			freeIndex = blockIndex
			freeBlock = block
			break
		}

		if freeIndex < 0 {
			continue
		}

		// handle current block deletion
		left := currIndex
		right := currIndex
		amount := currBlock.amount

		if currIndex-1 >= 0 && blocks[currIndex-1].id < 0 {
			left = currIndex - 1
			amount += blocks[left].amount
		}

		if currIndex+1 < len(blocks) && blocks[currIndex+1].id < 0 {
			right = currIndex + 1
			amount += blocks[right].amount
		}

		blocks = append(append(append([]block{}, blocks[:left]...), block{-1, amount}), blocks[right+1:]...)
		freeBlock = blocks[freeIndex]

		// handle current block insertion
		if freeBlock.amount > currBlock.amount {
			blocks = append(append(append([]block{}, blocks[:freeIndex]...), []block{currBlock, {-1, freeBlock.amount - currBlock.amount}}...), blocks[freeIndex+1:]...)
			continue
		}

		blocks = append(append(append([]block{}, blocks[:freeIndex]...), currBlock), blocks[freeIndex+1:]...)
	}

	// sum
	for blockIndex, i := 0, 0; blockIndex < len(blocks); blockIndex++ {
		block := blocks[blockIndex]

		for range block.amount {
			if block.id >= 0 {
				sum += i * block.id
			}

			i++
		}
	}

	return sum
}
