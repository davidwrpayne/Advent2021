package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
}

func part1() {
	lines, err := ReadFile("day11/input.txt")
	if err != nil {
		panic("failed to read input file")
	}

	xSize := len(lines[0])
	ySize := len(lines)
	b := NewBoard(xSize, ySize)

	for y, line := range lines {
		for x, rune := range line{
			number, err := strconv.Atoi(string(rune))
			if err != nil {
				panic ("failed to convert")
			}
			b.Map[x][y] = number
		}
	}

	//totalFlashes := 0
	var i int
	for i = 1; i <= 300; i++ {
		flashes := b.stepMap()
		if b.YSize * b.XSize <= flashes {
			fmt.Printf("all flashed on %d", i)
		}
	}
	//fmt.Printf("after %d stesp %d flashes \n", i, totalFlashes)
}

type Board struct {
	Map [][]int
	XSize int
	YSize int
}

func NewBoard(xSize, ySize int) *Board {
	energyMap := make([][]int, xSize) // x,y coordinates
	for x := 0; x < xSize; x++ {
		energyMap[x] = make([]int,ySize)
	}
	return &Board{
		energyMap,
		xSize,
		ySize,
	}
}


func (b *Board)PrintBoard() {
	for y:= 0; y< b.YSize; y++ {
		for x:= 0; x <b.XSize; x++ {
			fmt.Printf("%d",b.Map[x][y])
		}
		fmt.Printf("\n")
	}
}

func (b *Board)stepMap() int {

	flashing := []*Point{}
	for x := 0; x < b.XSize; x++ {
		for y := 0; y < b.YSize; y++ {
			b.Map[x][y] += 1
			if b.Map[x][y] == 10 {
				flashing = append(flashing, NewPoint(x,y))
			}
		}
	}
	for len(flashing) > 0 {
		element := len(flashing) - 1
		flasher := flashing[element]
		flashing = flashing[:element]

		neighbors := getNeighbors(flasher)
		for _, n := range neighbors {
			if n.x() < 0 || n.y() < 0 || n.x() >= b.XSize || n.y() >= b.YSize {
				continue
			}
			b.Map[n.x()][n.y()] += 1
			if b.Map[n.x()][n.y()] == 10 {
				flashing = append(flashing, NewPoint(n.x(), n.y()))
			}
		}
	}
	countOfFlashes := 0
	for x := 0; x < b.XSize; x++ {
		for y := 0; y < b.YSize; y++ {
			if b.Map[x][y] >= 10 {
				b.Map[x][y] = 0
				countOfFlashes += 1
			}
		}
	}
	return countOfFlashes
}

func NewPoint(x,y int) *Point {
	return &Point{x,y}
}

type Point [2]int

func (p *Point)x()int {
	return p[0]
}

func (p *Point)y()int {
	return p[1]
}

func getNeighbors(p *Point) []*Point {
	return []*Point{
		NewPoint(p.x() - 1, p.y() - 1),
		NewPoint(p.x() - 1, p.y()),
		NewPoint(p.x() - 1, p.y() + 1),
		NewPoint(p.x(), p.y() - 1),
		NewPoint(p.x(), p.y() + 1),
		NewPoint(p.x() + 1, p.y() - 1),
		NewPoint(p.x() + 1, p.y()),
		NewPoint(p.x() + 1, p.y() + 1),
	}
}

func ReadFile(s string) ([]string, error) {
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
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("scanner error")
	}
	return lines, nil
}
