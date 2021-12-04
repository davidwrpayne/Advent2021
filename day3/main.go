package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part2()
}

const MAXCOLUMN = 12
func part1() {
	lines, err := getLines("/Users/david.payne/personal/2021-advent-code/day3/input.txt")
	if err != nil {
		panic("error")
	}

	lineLength := MAXCOLUMN
	counts := make([]int, lineLength)
	for _, line := range lines {
		for j := 0; j < lineLength; j++ {
			if line[j] ==  '1' {
				counts[j] += 1
			}
		}
	}

	fmt.Printf("counts: %v ", counts)
	fmt.Print("\n")
	countOfLines := len(lines)
	gamma := uint(0)
	bitmask := ^(^uint(0) << lineLength) // not of 0 == 111111...1111   shift left length of line == 111...110000000,  not again 0000..00001111111

	for j:= 0; j < lineLength; j++{
		gamma = gamma << 1
		if counts[j] >= countOfLines / 2 {
			gamma = gamma + 1
		}
	}

	epsilon := ^gamma
	fmt.Printf("gamma %b \n", gamma)
	fmt.Printf("gamma decimal %d \n", gamma)
	fmt.Printf("epsilon %b \n", epsilon&bitmask)
	fmt.Printf("epsilon decimal %d \n", epsilon&bitmask)

	result := gamma * (epsilon&bitmask)
	fmt.Printf("multiplication %d \n", result)
}




func part2(){
	lines, err := getLines("/Users/david.payne/personal/2021-advent-code/day3/input.txt")
	if err != nil {
		panic("error")

	}

	values := make([]uint,len(lines))
	for i, line := range lines {
		values[i] = convertToByteSlice(line)
	}

	most := filterListColumn(true, values)
	fmt.Printf("most common: %0b\n", most)
	fmt.Printf("most common: %d\n", most)
	least := filterListColumn(false, values)
	fmt.Printf("least common: %0b\n", least)
	fmt.Printf("least common: %d\n", least)

	7928162
}

func filterListColumn(mostCommon bool, list []uint)uint {
	currentColumn := 0
	currentList := list
	for ; len(currentList) > 1;{
		var finderBit uint
		if mostCommon {
			finderBit = findMostCommonBit(currentColumn, currentList)
		} else {
			finderBit = findLeastCommonBit(currentColumn, currentList)
		}
		currentList = filterList(currentColumn, finderBit, currentList)
		currentColumn++
	}

	return currentList[0]
}

func filterList(column int, match uint, list []uint)[]uint {
	result := []uint{}
	shiftDistance := (MAXCOLUMN - column - 1) // correct. creates a bitmask at
	bitmask := uint(1) << shiftDistance
	match = match << shiftDistance
	for _, line := range list {
		if (line & bitmask) == match {
			result = append(result, line)
		}
	}
	return result
}


func findLeastCommonBit(column int, lines []uint) uint {
	if uint(1) == findMostCommonBit(column,lines) {
		return uint(0)
	} else {
		return uint(1)
	}
}

// findMostCommonBit finds the most common bit from the left either a 0 or 1
func findMostCommonBit(column int, lines []uint) uint {
	bitmask := uint(1) << (MAXCOLUMN - column - 1)
	countOfZeros := 0
	for i := 0; i < len(lines); i++ {
		if bitmask & lines[i] == uint(0) {
			countOfZeros++
		}
	}
	if  countOfZeros > len(lines) / 2 {
		return uint(0)
	} else {
		return uint(1)
	}
}


func convertToByteSlice(line string)uint {
	value := uint(0)

	for j := 0; j < len(line); j++ {
		value = value << 1
		if line[j] ==  '1' {
			value = value + 1
		}
	}
	return value
}


func getLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	result := []string{}
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}