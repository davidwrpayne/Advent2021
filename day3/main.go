package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1()
}



func part1() {
	lines, err := getLines("/Users/david.payne/personal/2021-advent-code/day3/input.txt")
	if err != nil {
		panic("error")
	}


	lineLength := 12
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




//func part2(){
//	lines, err := getLines("/Users/david.payne/personal/2021-advent-code/day2/input.txt")
//	if err != nil {
//		panic("error")
//
//	}
//
//
//}
//


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