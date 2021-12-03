package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part2()
}



func part1() {
	lines, err := getLines("/Users/david.payne/personal/2021-advent-code/day2/input.txt")
	if err != nil {
		panic("error")
	}

	myLocation := NewLocation()
	for i := 0; i < len(lines); i++ {
		tokens := regexp.MustCompile("\\s").Split(lines[i], 2)
	 	count, err := strconv.Atoi(tokens[1])
	 	if err != nil {
	 		panic(fmt.Sprintf("couldn't convert line: %d", i))
		}
		operation := tokens[0]

		myLocation.processOperation(operation, count)
	}
	fmt.Printf("this is location %v", myLocation)

}

func part2(){
	lines, err := getLines("/Users/david.payne/personal/2021-advent-code/day2/input.txt")
	if err != nil {
		panic("error")

	}

	newSub := NewSubmarine()
	for i := 0; i < len(lines); i++ {
		tokens := regexp.MustCompile("\\s").Split(lines[i], 2)
		amount, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(fmt.Sprintf("couldn't convert line: %d", i))
		}
		operation := tokens[0]

		newSub.processOperation(operation, amount)
	}
	fmt.Printf("this is sub aim: %v depth: %v and horizontal pos: %v",newSub.Aim, newSub.Y, newSub.X)
	fmt.Printf("final answer %v", newSub.X * newSub.Y)
}



type Submarine struct {
	X int
	Y int
	Aim int
}
func NewSubmarine() *Submarine {
	return &Submarine{
		X: 0,
		Y: 0,
		Aim: 0,
	}
}

func (s *Submarine) processOperation(command string, amount int) {
	switch command {
	case "forward":
		s.X += amount
		s.Y += s.Aim * amount
	case "down":
		s.Aim += amount
	case "up":
		s.Aim -= amount
	default:
		panic(fmt.Sprintf("un handled operation %s",command))
	}
}

type Location struct {
	Depth int
	ForwardDistance int
}
func NewLocation() *Location {
	return &Location{
		Depth:  0,
		ForwardDistance: 0,
	}
}

func (l *Location) processOperation(operation string, distance int) {
	switch operation {
		case "forward":
			l.ForwardDistance += distance
		case "down":
			l.Depth += distance
		case "up":
			l.Depth -= distance
		default:
			panic(fmt.Sprintf("un handled operation %s",operation))
	}
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