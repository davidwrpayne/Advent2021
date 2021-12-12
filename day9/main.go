package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	part2()

}

func part2() {
	
}


func part1() {
	lines, err := readFile("day9/input.txt")
	if err != nil {
		panic("failed to read file lines")
	}
	//fmt.Printf("lines: %v", lines)
	var seaFloorMap [][]int // x,y formation
	seaFloorMap, err = convertLines(lines)
	if err != nil {
		panic("failed to map floor")
	}

	board := Board{
		XSize: len(seaFloorMap),
		YSize: len(seaFloorMap[0]),
		Map: seaFloorMap,
	}
	localMins := findLocalMins(board)
	totalRisk := 0
	for _, point := range localMins {
		totalRisk += calcRiskLevel(point,board)
	}
	fmt.Printf("Risk level%v\n", totalRisk)
}

type Point [2]int
type Board struct {
	XSize int
	YSize int
	Map   [][]int
}
func findLocalMins(b Board) []Point {

	localMins := []Point{}
	for x:=0; x <b.XSize ; x++ {
		for y := 0; y < b.YSize; y++ {
			value := b.Map[x][y]
			up := Point{x, y-1}
			down := Point{x, y+1}
			left := Point{x-1,y}
			right := Point{x+1,y}
			if GetValue(up, b) > value && GetValue(down, b) > value && GetValue(left, b) > value && GetValue(right, b) > value {
				localMins = append(localMins, Point{x,y})
			}
		}
	}

	return localMins

}

func calcRiskLevel(p Point, b Board) int {
	return GetValue(p,b) + 1
}

// returns MAX int size if p is outside input range
func GetValue(p Point, b Board) int {
	if p[0] < 0 || p[0] >= b.XSize || p[1] >= b.YSize || p[1] < 0 {
		return math.MaxInt
	} else {
		return b.Map[p[0]][p[1]]
	}
}


// convertLines returns the digits of the sea floor size
func convertLines(lines []string) ([][]int, error) {
	xSize := len(lines[0])
	ySize := len(lines)
	result := make([][]int,xSize)
	for y, line := range lines {
		for x := 0; x < xSize; x++ {
			if y == 0 {
				result[x] = make([]int, ySize)
			}
			num, err := strconv.Atoi(string(line[x]))
			if err != nil {
				return nil, err
			}
			result[x][y] = num
		}
	}
	return result, nil
}

func readFile(s string) ([]string, error){
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
		lines = append(lines,  scanner.Text())
		//fmt.Printf("Reading vetor: %v \n", vectorLine)

	}

	if err := scanner.Err(); err != nil {
		return nil,errors.New("scanner error")
	}
	return lines,nil
}