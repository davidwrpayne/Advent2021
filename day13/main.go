package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()
}

type Board struct {
	Points [][]bool
	PointList []*Point
}

type Point [2]int

func NewPoint(x,y int) *Point{
	return &Point{x,y}
}

func (p *Point)x()int {
	return p[0]
}
func (p *Point)y()int {
	return p[1]
}


func part1() {
	lines, err:= readFile("./day13/input.txt")
	if err != nil {
		panic(err)
	}

	pointStrings := []string{}
	commandStrings := []string{}
	readPoints := true
	for _, line := range lines {
		if readPoints && "" == strings.TrimSpace(line) {
			readPoints = false
			continue
		}
		if readPoints {
			pointStrings = append(pointStrings, line)
		} else {
			commandStrings = append(commandStrings, line)

		}
	}

	points := convertPointStrings(pointStrings)
	board := NewBoard(points)
	commands := convertCommandStrings(commandStrings)

	fmt.Printf("point Strings %v", points)
	fmt.Printf("command Strings %v", commands)

	result := runFold(board, commands, 1)
	fmt.Printf("number of points: %v\n", len(result.PointList))
}


func part2() {
	lines, err:= readFile("./day13/input.txt")
	if err != nil {
		panic(err)
	}

	pointStrings := []string{}
	commandStrings := []string{}
	readPoints := true
	for _, line := range lines {
		if readPoints && "" == strings.TrimSpace(line) {
			readPoints = false
			continue
		}
		if readPoints {
			pointStrings = append(pointStrings, line)
		} else {
			commandStrings = append(commandStrings, line)

		}
	}

	points := convertPointStrings(pointStrings)
	board := NewBoard(points)
	commands := convertCommandStrings(commandStrings)

	result := runFold(board, commands, len(commands))
	result.printBoard()
}


func (b *Board) printBoard() {
	fmt.Printf("\n\n")
	for x := len(b.Points) - 1; x >= 0; x-- {
		for y:=0; y < len(b.Points[0]); y++ {
			if b.Points[x][y] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func NewBoard(points []*Point) *Board {
	maxX := math.MinInt
	maxY := math.MinInt
	for _, p := range points {
		if maxX <= p.x() {
			maxX = p.x()
		}
		if maxY <= p.y() {
			maxY = p.y()
		}
	}
	pointArray := make([][]bool, maxX+1)
	b := &Board{}

	for i := 0; i <= maxX; i++ {
		pointArray[i] = make([]bool, maxY+1)
	}

	b.Points = pointArray
	b.PointList = points

	for _, p := range points {
		b.Points[p.x()][p.y()] = true
	}

	return b
}

func runFold(board *Board, commands []Command, i int) *Board {

	command := commands[0]
	remainingCommands := commands[1:]
	var newBoard *Board
	if command.axis == "x" {
		newBoard = foldVertical(board, command.number)
	} else {
		newBoard = foldHorizontal(board, command.number)
	}
	if i > 1 {
		return runFold(newBoard, remainingCommands, i-1)
	} else {
		return newBoard
	}
}


// folding vertical means line = x
// everything right of x becomes new point left of x
// everything left of x stays still in the new point list

func foldVertical(board *Board, number int) *Board {
	newPointList := []*Point{}
	for _, point := range board.PointList {
		if point.x() <= number { // point is left of line
			// point remains where it is
			newPointList = append(newPointList, point)
		} else {
			delta := float64(point.x() - number)
			newXLoc := number - int(math.Abs(delta))
			newPoint := NewPoint(newXLoc, point.y())

			if board.Points[newPoint.x()][newPoint.y()] == false { // point not recorded on board already
				board.Points[newPoint.x()][newPoint.y()] = true
				newPointList = append(newPointList, newPoint)
			} // otherwise point already recorded
		}
	}
	return NewBoard(newPointList)
}

func foldHorizontal(board *Board, number int) *Board{
	newPointList := []*Point{} // dont know the size of the slice yet
	for _, point := range board.PointList {
		if point.y() < number { // point is above the line
			// point remains where it is
			newPointList = append(newPointList, point)
		} else { // point is below the line
			delta := float64(point.y() - number)
			newYLoc := number - int(math.Abs(delta))
			newPoint := NewPoint(point.x(), newYLoc)

			if board.Points[newPoint.x()][newPoint.y()] == false { // point not recorded on board already
				board.Points[newPoint.x()][newPoint.y()] = true
				newPointList = append(newPointList, newPoint)
			} // otherwise point already recorded
		}
	}
	return NewBoard(newPointList)
}

type Command struct {
	axis string
	number int
}

func convertCommandStrings(commandStrings []string) []Command {
	commands := []Command{}
	for _, commandString := range commandStrings {
		foldLine :=strings.Split(strings.Split(commandString, "fold along ")[1], "=")
		axis := foldLine[0]
		number,_ := strconv.Atoi(foldLine[1])
		commands = append(commands , Command{axis: axis, number: number})
	}
	return commands
}

func convertPointStrings(pointStrings []string) []*Point {
	points := []*Point{}
	for _, pointString := range pointStrings {
		coordinates := strings.Split(pointString, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		points = append(points , NewPoint(x,y))
	}
	return points
}

func readFile(s string) ([]string, error) {
	lines := []string{}
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Couldn't read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//fmt.Printf("Reading vetor: %v \n", vectorLine)

	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("scanner error")
	}
	return lines, nil
}