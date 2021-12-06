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
	part1()

}

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x,
		y,
	}
}
func part1() {
	var lines [][2]Point

	lines = readFile("./day5/input.txt")
	locationMap := NewMap()
	for _, vector := range lines {
		locationMap.graphLine(vector[0].x, vector[0].y, vector[1].x, vector[1].y)
	}
	//_ = lines
	//locationMap.printMap()
	fmt.Printf("score %d \n", locationMap.score())
}

func readFile(s string) [][2]Point {
	vectors := [][2]Point{}
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
		panic("Couldn't read file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		vectorLine := scanner.Text()
		//fmt.Printf("Reading vetor: %v \n", vectorLine)
		vector := readVector(vectorLine)
		//fmt.Printf("Vector: %v \n", vector)
		vectors = append(vectors, vector)
	}

	if err := scanner.Err(); err != nil {
		panic("scanner error")
	}
	return vectors
}

func readVector(text string) [2]Point {
	re := regexp.MustCompile("(\\d+),(\\d+)\\s->\\s(\\d+),(\\d+)")
	matches := re.FindStringSubmatch(text)

	x1, _ := strconv.Atoi(matches[1])
	y1, _ := strconv.Atoi(matches[2])
	x2, _ := strconv.Atoi(matches[3])
	y2, _ := strconv.Atoi(matches[4])

	p1 := NewPoint(x1, y1)
	p2 := NewPoint(x2, y2)
	result := [2]Point{p1, p2}
	return result
}

const BOARD_SIZE_X = 1000
const BOARD_SIZE_Y = 1000

func NewMap() *LocationMap {

	columns := make([][]int, BOARD_SIZE_X)
	for i := range columns {
		columns[i] = make([]int, BOARD_SIZE_Y)
	}
	return &LocationMap{
		pixels: columns,
	}
}

type LocationMap struct {
	pixels [][]int // x,y
}

func (l *LocationMap) graphLine(x1, y1, x2, y2 int) {
	if x1 == x2 {
		l.verticalLine(x1, y1, y2)
	} else {

		slope := (y1 - y2) / (x1 - x2)
		intercept := y1 - (slope * x1)

		startingX := min(x1, x2)
		endingX := max(x1, x2)

		for x := startingX; x <= endingX; x++ {
			y := slope*x + intercept
			l.pixels[x][y] += 1
		}


	}
}

func (l *LocationMap) verticalLine(x, y1, y2 int) {
	startingY := min(y1, y2)
	endingY := max(y1, y2)
	for i := startingY; i <= endingY; i++ {
		l.pixels[x][i] += 1
	}
}

func (l *LocationMap) printMap() {
	for y := 0; y < BOARD_SIZE_Y; y++ {
		for x := 0; x < BOARD_SIZE_X; x++ {
			value := l.pixels[x][y]
			var printString string
			if value == 0 {
				printString = ". "
			} else {
				printString = fmt.Sprintf("%d ",value)
			}

			fmt.Print(printString)
		}
		fmt.Println()
	}
}

func (l *LocationMap) score() int {
	score := 0
	for x := 0; x < BOARD_SIZE_X; x++ {
		for y := 0; y < BOARD_SIZE_Y; y++ {
			if l.pixels[x][y] > 1 {
				score += 1
			}
		}
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
