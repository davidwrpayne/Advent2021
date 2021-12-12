package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//part1()
	part2()

}

func part1() {

	lines := readFile("./day8/input.txt")


	countOfAll := 0
	for _, value := range lines {
		countOfAll += solveNumbersPart1(value)
	}
	fmt.Printf("answer of total %d \n", countOfAll)
}

func solveNumbersPart1(line [2]string) int{
	displayed := [4]bool{}
	substrings := strings.Fields(line[1])// only looking at outputs
	countOfAll := 0

	for _, value := range substrings {
		if len(value) == 2 && !displayed[0] {
			fmt.Printf("1 output value %s\n", value)
			displayed[0] = true
		}
		if len(value) == 3 && !displayed[1] {
			fmt.Printf("4 output value %s\n", value)
			displayed[1] = true
		}
		if len(value) == 4 && !displayed[2] {
			fmt.Printf("7 output value %s\n", value)
			displayed[2] = true
		}
		if len(value) == 7 && !displayed[3] {
			fmt.Printf("8 output value %s\n", value)
			displayed[3] = true
		}

		if len(value) == 2 || len(value) == 3 || len(value) == 4 || len(value) == 7 {
			countOfAll += 1
		}
	}

	return countOfAll
}

func part2() {
	lines := readFile("./day8/input.txt")

	sum := 0
	for _, value := range lines {
		input := value[0]
		output := value[1]
		translations := NewTranslationMemoization()

		for _, token := range strings.Fields(input) {
			sortedToken := sortStringAlphabetically(token)
			translations.RemainingTokens = append(translations.RemainingTokens, sortedToken)
		}

		translations.ApplyRules()

		fmt.Printf("translations object %v\n", translations)

		numberOutput :=  []string{}
		for _, token := range strings.Fields(output) {
			sortedToken := sortStringAlphabetically(token)
			fmt.Printf("output token: %s \n", sortedToken)
			number, err := translations.TranslateToken(sortedToken)
			if err != nil {
				panic("failure to translate")
			}
			fmt.Printf("translation is: %d \n", number)
			numberOutput = append(numberOutput, fmt.Sprintf("%d", number))
		}
		value, err := strconv.Atoi(strings.Join(numberOutput[:], ""))
		if err != nil {
			panic("coudn't convert to number")
		}
		sum += value
	}
	fmt.Printf("total is %d", sum)


}




func NewTranslationMemoization() *TranslationMemoization {
	return &TranslationMemoization{
		RemainingTokens: []string{},
		LookupNumberToToken: map[int]string{},
		LookupTokenToNumber: map[string]int{},
	}
}

func (t *TranslationMemoization) TranslateToken(token string)(int, error) {
	if len(t.RemainingTokens) != 0 {
		return 0, errors.New("Still remaining tokens un translated")
	}
	return t.LookupTokenToNumber[token], nil
}

func (t *TranslationMemoization) ApplyRules() {
	t.TokenLengthIs(2,1)
	t.TokenLengthIs(4,4)
	t.TokenLengthIs(3,7)
	t.TokenLengthIs(7,8)

	t.SetIsSubSetOfToken(SetSubtract(t.GetSet(8), t.GetSet(7)),6)
	t.SetIsSubSetOfToken(t.GetSet(4), 9)
	t.TokenLengthIs(6, 0)
	t.SetIsSubSetOfToken(SetIntersection(t.GetSet(9), t.GetSet(6)), 5)
	t.SetIsSubSetOfToken(t.GetSet(1), 3)
	t.TokenIsLastTokenRemaining(2)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func removeStringSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func SetSubtract(a, b map[rune]bool) map[rune]bool {
	subtractedSet := map[rune]bool{}
	for r,v := range a {
		if v {
			if b[r] != true {
				subtractedSet[r] = true
			}
		}
	}
	return subtractedSet
}

func SetIntersection(a, b map[rune]bool) map[rune]bool {
	intersection := map[rune]bool{}
	if len(a) > len(b) {
		a, b = b, a // better to iterate over a shorter set
	}
	for k,_ := range a {
		if b[k] {
			intersection[k] = true
		}
	}
	return intersection
}

func SetAddition(a ...map[rune]bool) map[rune]bool {
	additionSet := map[rune]bool{}
	for _, set := range a {
		for r,v := range set {
			if v {
				additionSet[r] = true
			}
		}
	}
	return additionSet
}

// if a is fully contained in b
func SetIsSubset(a, b map[rune]bool) bool {
	for r, v := range a {
		if v {
			if b[r] != true {
				return false
			}
		}
	}
	return true
}
func ToSet(token string)map[rune]bool {
	tokenSet := map[rune]bool {}
	for _, r := range []rune(token) {
		tokenSet[r] = true
	}
	return tokenSet
}

func (t *TranslationMemoization) TokenLengthIs(length int, id int) {
	for index, value := range t.RemainingTokens {
		if len(value) == length {
			t.RemainingTokens = removeStringSlice(t.RemainingTokens,index)
			t.LookupNumberToToken[id] = value
			t.LookupTokenToNumber[value] = id
			return
		}
	}
}

func (t *TranslationMemoization) FindSix() {
	searchSet := SetSubtract(ToSet(t.LookupNumberToToken[8]), ToSet(t.LookupNumberToToken[7]))
	t.SetIsSubSetOfToken(searchSet, 6)
}

func (t *TranslationMemoization) TokenIsSubsetOfSet(set map[rune]bool, numberOfToken int) {
	for index, remainingToken := range t.RemainingTokens {
		if SetIsSubset(ToSet(remainingToken), set) {
			t.LookupTokenToNumber[remainingToken] = numberOfToken
			t.LookupNumberToToken[numberOfToken] = remainingToken
			t.RemainingTokens = removeStringSlice(t.RemainingTokens, index)
			return
		}
	}
}

func (t *TranslationMemoization) SetIsSubSetOfToken(searchSet map[rune]bool, id int) {
	for index, remainingToken := range t.RemainingTokens {
		if SetIsSubset(searchSet, ToSet(remainingToken)) {
			t.LookupTokenToNumber[remainingToken] = id
			t.LookupNumberToToken[id] = remainingToken
			t.RemainingTokens = removeStringSlice(t.RemainingTokens, index)
			return
		}
	}
}

func (t *TranslationMemoization) GetSet(i int) map[rune]bool {
	return ToSet(t.LookupNumberToToken[i])
}

func (t *TranslationMemoization) TokenIsLastTokenRemaining(id int) {
	if len(t.RemainingTokens) == 1 {
		lastToken := t.RemainingTokens[0]
		t.LookupNumberToToken[id] = lastToken
		t.LookupTokenToNumber[lastToken] = id
		t.RemainingTokens = removeStringSlice(t.RemainingTokens,0)
	} else {
		panic("invalid number of tokens left")
	}

}




type TranslationMemoization struct {
	RemainingTokens []string
	LookupNumberToToken map[int]string
	LookupTokenToNumber map[string]int
}


func readFile(s string) [][2]string {
	inputOutputLines := [][2]string{}
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
		result := strings.Split(vectorLine, "|")
		inputOutputLines = append(inputOutputLines, [2]string{result[0], result[1]})
	}

	if err := scanner.Err(); err != nil {
		panic("scanner error")
	}
	return inputOutputLines
}




// taken from https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
type SortableRunes []rune
func (r SortableRunes) Len() int           { return len(r) }
func (r SortableRunes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r SortableRunes) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func sortStringAlphabetically(token string) string {
	var runes SortableRunes = StringToRuneSlice(token)
	sort.Sort(runes)
	return string(runes)
}
