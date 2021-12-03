// Declaration of the main package
package main

// Importing packages
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Main function
func main() {

	err := part1()
	if err != nil {
		log.Fatal("failed part 1", err)
	}

	err = part2()
	if err != nil {
		log.Fatal("failed part 2", err)
	}
}

func convertLines(file *os.File) ([]int, error) {
	lines, err := getLines(file)
	if err != nil {
		log.Fatal("failed getting lines from file", err)
		return nil, err
	}
	convertedInts := make([]int, len(lines))
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal("error converting", err)
			return nil, err
		}
		convertedInts[i] = value
	}
	return convertedInts, nil
}

func getLines(file *os.File) ([]string, error) {
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

func part1() error {

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	convertedInts, err := convertLines(file)
	if err != nil {
		log.Fatal("failed to convert lines", err)
		return err
	}

	fmt.Printf("count of single sum increases: %d\n", countIncreases(convertedInts))
	return nil
}

func countIncreases(input []int) int {
	countOfIncrease := 0
	for i := 0; i < len(input); i++ {
		if i > 0 {
			if input[i-1] < input[i] {
				countOfIncrease++
			}
		}
	}
	return countOfIncrease
}


func part2() error {

	fileName := "input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, err := convertLines(file)
	if err != nil {
		return err
	}

	// perform a sliding window 3 value sum check increases

	sliding3ValueSum := slidingWindowSum(lines, 3)
	increases := countIncreases(sliding3ValueSum)
	fmt.Printf("3 value sliding window sum increases: %d \n", increases)
	return nil
}

func slidingWindowSum(input []int, n int) []int {
	var sliding3WindowSum []int
	valuesUnableToBeCounted := n-1
	sliding3WindowSum = make([]int, len(input) - valuesUnableToBeCounted )


	for i, j := valuesUnableToBeCounted, 0; i <= len(input) - 1; i, j = i+1,j+1 {
		slidingWindowSum := 0
		for k := 0; k < n; k++ {
			slidingWindowSum += input[i-k]
		}
		sliding3WindowSum[j] = slidingWindowSum
	}
	return sliding3WindowSum
}
