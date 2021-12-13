package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)


func main() {
	//part1()
	part2()

}

func part2() {

	lines, err := ReadFile("day10/input.txt")
	if err != nil {
		panic("failed to read input file")
	}

	totals := []int{}
	for _,line := range lines {
		score := 0
		chunks := strings.Fields(line)
		for _,c := range chunks {
			stack, err := getIncompleteChunks(c)
			if err != nil {
				continue // skip scoring chunk as corrupt or completed
			}

			// while the stack still has items pop and score them

			for len(stack) > 0 {
				 n := len(stack)-1
				 element := stack[n]
				 stack = stack[:n]

				 closingBracket := getAutoCloseSuggestion(element)
				 bracketScore := getPart2Score(closingBracket)
				 score = score * 5
				 score += bracketScore
			}

		}
		fmt.Printf(fmt.Sprintf("line score %d\n", score))
		if score > 0 {
			totals = append(totals, score)
		}

	}
	sort.Ints(totals)

	fmt.Printf(fmt.Sprintf("total score %v\n", totals))
	fmt.Printf(fmt.Sprintf("middle score %d\n", totals[len(totals)/2]))
}

func part1() {

	lines, err := ReadFile("day10/input.txt")
	if err != nil {
		panic("failed to read input file")
	}

	totals := 0
	for _,line := range lines {
		totalLineScore := 0
		chunks := strings.Fields(line)
		for _,c := range chunks {
			score, _ := scoreChunk(c)
			if score > 0 {
				totalLineScore += score
			}
		}
		fmt.Printf(fmt.Sprintf("line score %d\n", totalLineScore))
		totals += totalLineScore
	}
	fmt.Printf(fmt.Sprintf("total score %d\n", totals))
}

var runePairs = [][2]rune{
	{'(',')'},
	{'[',']'},
	{'<','>'},
	{'{','}'},
}

func closingRuneMatchesOpeningRune(openingRune string, s string) bool {
	switch openingRune {
	case "(":
		return s == ")"
	case "[":
		return s == "]"
	case "<":
		return s == ">"
	case "{":
		return s == "}"
	default:
		return false
	}
}


func getAutoCloseSuggestion(in string)string {
	switch in {
	case "(":
		return ")"
	case "[":
		return "]"
	case "<":
		return ">"
	case "{":
		return "}"
	default:
		panic("unexpected string input")
	}
}
func getPart2Score(in string)int {
	switch in {
	case ")":
		return 1
	case "]":
		return 2
	case ">":
		return 4
	case "}":
		return 3
	default:
		panic("unexepected input")
	}
}

func getPart1Score(in string)int {
	switch in {
	case ")":
		return 3
	case "]":
		return 57
	case ">":
		return 25137
	case "}":
		return 1197
	default:
		return 1
	}
}

// returns error if the chunk is corrupted
func getIncompleteChunks(chunk string) ([]string, error) { // returns the remaining stack or an error if its  not incomplete
	stack := []string{}
	for _, rune := range chunk {
		s := string(rune)
		if runeIsClosingRune(rune) {
			n := len(stack)-1
			lastOpeningRune := stack[n]
			stack = stack[:n]
			if !closingRuneMatchesOpeningRune(lastOpeningRune, s) { //corrupted
				return nil, errors.New("it was an corrupted chunk")
			}
			// does closing rune match last opening rune?
		} else if runeIsOpeningRune(rune) {
			stack = append(stack, s)
		}
		// current state
		// if the current rune is a closing rune and the last opening rune in the stack (i think the stack only tracks opening runes)
		// if those runes don't match then I think the line is corrupt.
		// record the failing rune
		// else if it matches pop the rune of the stack
		//
		// otherwise add the opening rune to the stack.
		// if we reach the end of the line without an empty stack we are corrupted.
		//if runeIsNotOpeningRune(rune) {
		//
		//}

		// also if theres a space consider that the end of a line
	}
	if len(stack) == 0 {
		return nil, errors.New("not an incomplete chunk")
	} else {
		return stack, nil
	}
}

func scoreChunk(chunk string) (int, string) {
	stack := []string{}
	for _, rune := range chunk {
		s := string(rune)
		if runeIsClosingRune(rune) {
			n := len(stack)-1
			lastOpeningRune := stack[n]
			stack = stack[:n]
			if !closingRuneMatchesOpeningRune(lastOpeningRune, s) {
				return getPart1Score(s),fmt.Sprintf("expected %s but chunk found %s", lastOpeningRune,s)
			}
			// does closing rune match last opening rune?
		} else if runeIsOpeningRune(rune) {
			stack = append(stack, s)
		}
		// current state
		// if the current rune is a closing rune and the last opening rune in the stack (i think the stack only tracks opening runes)
		// if those runes don't match then I think the line is corrupt.
		// record the failing rune
		// else if it matches pop the rune of the stack
		//
		// otherwise add the opening rune to the stack.
		// if we reach the end of the line without an empty stack we are corrupted.
		//if runeIsNotOpeningRune(rune) {
		//
		//}

		// also if theres a space consider that the end of a line
	}
	if len(stack) == 0 {
		return 0, ""
	} else {
		return -1, fmt.Sprintf("expected %s but chunk was incomplete", stack[len(stack)-1])
	}
}


func runeIsClosingRune(r rune) bool {
	for _, pair := range runePairs {
		if r == pair[1] {
			return true
		}
	}
	return false
}

func runeIsOpeningRune(r rune) bool {
	for _, pair := range runePairs {
		if r == pair[0] {
			return true
		}
	}
	return false
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
