package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    part2()
}

func part1() {
	fileName := "day6/input.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic("failed to read file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := scanner.Text()
	if scanner.Err() != nil {
		panic("failed to scan file")
	}

	initialAges := parseLine(firstLine)

	result := processDay(0, 80, initialAges, false)
    fmt.Printf("lattern fish count: %d", len(result))
}

const LANTERN_FISH_TIMER_RESET = 6
const LANTERN_FISH_MAX_TIMER = 9 // 0 through timer 8

func processDay(current, n int, initialAges []int, print bool) []int {
	nextAge := make([]int, len(initialAges))
	if current == n {
		return initialAges
	}
	for i := range initialAges {
		if initialAges[i] == 0 {
			nextAge[i] = LANTERN_FISH_TIMER_RESET
			nextAge = append(nextAge, LANTERN_FISH_MAX_TIMER - 1)
		} else {
			nextAge[i] = initialAges[i] - 1
		}
	}
	if print {
		fmt.Printf("after %d day: %v\n", current, nextAge)
	}
	return processDay(current+1, n, nextAge, print)
}

func parseLine(line string) []int {

	subStrings := strings.Split(line, ",")
	result := make([]int, len(subStrings))
	for i := range subStrings {
		value, err := strconv.Atoi(subStrings[i])
		if err != nil {
			panic("failed to convert digits in line")
		}
		result[i] = value
	}
	return result
}


func part2() {
    fileName := "day6/input.txt"
    file, err := os.Open(fileName)
    if err != nil {
        panic("failed to read file")
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Scan()
    firstLine := scanner.Text()
    if scanner.Err() != nil {
        panic("failed to scan file")
    }

    initialAges := parseLine(firstLine)
    // combine fish counts

    combineFishCounts := createCountsFromInitial(initialAges)
    result := processDay2(0, 256, combineFishCounts, false)
    sum_of_fish := 0
    for i := range result {
        sum_of_fish += result[i]
    }
    fmt.Printf("lattern fish count: %d", sum_of_fish)
}

func processDay2(current int, n int, thisAge LanternFishCounts, print bool) LanternFishCounts {
    nextAge := make([]int, LANTERN_FISH_MAX_TIMER)
    if current == n {
        return thisAge
    }
    for i := 0 ; i < LANTERN_FISH_MAX_TIMER - 1; i++ {
        nextAge[i] = thisAge[i+1]
    }
    nextAge[LANTERN_FISH_TIMER_RESET] += thisAge[0]
    nextAge[LANTERN_FISH_MAX_TIMER - 1 ] = thisAge[0]

    return processDay2(current+1, n, nextAge, print)
}


func createCountsFromInitial(ages []int) LanternFishCounts {
    counts := make([]int, LANTERN_FISH_MAX_TIMER )

    for i := range ages {
        counts[ages[i]] += 1
    }
    return counts
}


type LanternFishCounts []int