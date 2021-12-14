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
	lines, err := ReadFile("day11/test_input.txt")
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

	b.PrintBoard()
	fmt.Printf("\n\n")
	b = stepMap(b)
	b.PrintBoard()
	fmt.Printf("\n\n")
	b = stepMap(b)
	b.PrintBoard()
	fmt.Printf("\n\n")
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

func stepMap(b *Board) *Board {
	newBoard := NewBoard(b.XSize,b.YSize)
	for x := 0; x < b.XSize; x++ {
		for y := 0; y < b.YSize; y++ {
			newValue := b.Map[x][y] + 1
			newBoard.Map[x][y] = newValue
		}
	}
	flashNeighbors(b)
	for x := 0; x < b.XSize; x++ {
		for y := 0; y < b.YSize; y++ {
			if newBoard.Map[x][y] > 9 {
				newBoard.Map[x][y] = 0
			}
		}
	}
	return newBoard
}

func flashNeighbors(board *Board) {
	for x := 0; x < board.XSize; x ++ {
		for y := 0; y < board.YSize; y ++ {
			if board.Map[x][y] > 9 {
				neighbors := [][2]int{
					{x-1,y-1},
					{x-1,y},
					{x-1,y+1},
					{x,y-1},
					//{x,y}, ignore center
					{x,y+1},
					{x+1,y-1},
					{x+1,y},
					{x+1,y+1},
				}
				for _, n := range neighbors {
					if n[0] < 0 || n[1] < 0 || n[0] >= board.XSize || n[1] >= board.YSize {
						continue
					}
					board.Map[n[0]][n[1]] += 1
				}
			}
		}
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
