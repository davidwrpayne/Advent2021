package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {

	lines, err := getLines("./day4/input.txt")
	if err != nil {
		panic(err)
	}
	//for _, line := range lines {
	//	fmt.Println(line)
	//}

	firstLineNumbers, lines := lines[0], lines[1:]
	numbersCalled := convertLine(firstLineNumbers, ",", 30000)
	boards := readBoards(lines)

	for _, number := range numbersCalled {
		for i, b := range boards {
			b.markNumber(number)
			if b.winningBoard() {
				fmt.Printf("winning board %v index %v", b.ScoredNumbers, i)
				fmt.Printf("score of winning %v", b.score(number))
				return
			}
		}
	}
}

//const REGEX = "\\s+" //multiple whitespaces
func readBoards(input []string) []*BingoBoard {
	bingoBoards := []*BingoBoard{}
	for ; len(input) >= BOARDSIZE+1; { // boardsize + 1 line of whitespace
		input = input[1:] // read whitespace
		boardLines := make([][]int, BOARDSIZE)
		for i := 0; i < BOARDSIZE; i++ {
			line := convertBoardLine(input[i])
			boardLines[i] = line
		}
		bingoBoards = append(bingoBoards, NewBingoBoard(boardLines))
		input = input[BOARDSIZE:] // remove board lines from input
	}
	return bingoBoards
}

const BOARDSIZE = 5

func convertLine(line string, regex string, size int) []int {
	result := make([]int, size)
	tokens := regexp.MustCompile(regex).Split(line, size)
	for i, token := range tokens {
		tokenValue, err := strconv.Atoi(token)
		if err != nil {
			fmt.Printf("token converting: %v \n", token)
			panic(fmt.Sprintf("couldn't convert line: %d", i))

		}
		result[i] = tokenValue
	}
	return result
}

func convertBoardLine(line string) []int {
	// each number is seperated by a space and each number is up to two digits,
	// if the number is a single digit the other digit is replace by a space
	s := strings.Fields(line)
	result := []int{}
	for _, field := range s {
		value, err := strconv.Atoi(field)
		if err != nil {
			panic("faield to convert")
		}
		result = append(result,value)
	}
	return result
}

func NewBingoBoard(numbers [][]int) *BingoBoard {
	return &BingoBoard{
		BoardNumbers: numbers,
		RowCounts:    [BOARDSIZE]int{},
		ColumnCounts: [BOARDSIZE]int{},
	}
}

type BingoBoard struct {
	BoardNumbers    [][]int        // row than column storage  stored [y],[x]
	ScoredLocations [][2]int       // array of [y,x] local coordinates that were called
	ScoredNumbers   []int          // array of numbers found on this board that were scored
	RowCounts       [BOARDSIZE]int //
	ColumnCounts    [BOARDSIZE]int // tracks how many marked numbers in each column there is if we reach 5 then we know board matched
}

// markNumber returns true if the board just won
func (b *BingoBoard) markNumber(n int) {
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			if n == b.BoardNumbers[i][j] {
				// found the number
				b.ScoredLocations = append(b.ScoredLocations, [2]int{i, j})
				b.ScoredNumbers = append(b.ScoredNumbers, n)
				b.ColumnCounts[j] += 1
				b.RowCounts[i] += 1
			}
		}
	}
}

func (b *BingoBoard) winningBoard() bool {
	for check := 0; check < BOARDSIZE; check++ {
		if b.ColumnCounts[check] == BOARDSIZE || b.RowCounts[check] == BOARDSIZE {
			return true
		}
	}
	return false
}

func (b *BingoBoard) score(lastCalledNumber int) int {
	sumOfAllNumbers := 0
	for y := 0; y < BOARDSIZE; y++ {
		for x := 0; x < BOARDSIZE; x++ {
			sumOfAllNumbers += b.BoardNumbers[y][x]
		}
	}
	score := sumOfAllNumbers
	for _, scoredNumber := range b.ScoredNumbers {
		score -= scoredNumber
	}
	if score < 0 {
		panic("Score was less than zero.. shouldn't be possible")
	}
	return score * lastCalledNumber
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
