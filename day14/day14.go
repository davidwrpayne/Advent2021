package day14

import (
	"davidwrpayne/2021-advent-code/fileutil"
	"fmt"
	"math"
	"strings"
)

func Start() {
	lines, err := fileutil.ReadFile("./day14/input.txt")
	if err != nil {
		panic(err)
	}

	template := lines[0]
	insertianRules := lines[2:]
	fmt.Printf("initial code %v\n", template)
	fmt.Printf("expansion commands %v\n", insertianRules)

	result := processInsertion(template, insertianRules, 40)

	counts := map[string]int {}
	for _,c := range result {
		character := string(c)
		if _, found :=  counts[character]; !found { // not found
			counts[character] = 0
		}
		counts[character] += 1 // increment count of character
	}

	mostCommon := string(result[0])
	leastCommon := string(result[0])
	countOfMost := math.MinInt
	countOfLeast := math.MaxInt
	for k,v := range counts {
		if v > countOfMost {
			countOfMost = v
			mostCommon = k
		}
		if v < countOfLeast {
			countOfLeast = v
			leastCommon = k
		}
	}


	fmt.Printf("result: '%v' \n", result)
	fmt.Printf("countOfLeast: '%v' \n", countOfLeast)
	fmt.Printf("least common: '%v' \n", leastCommon)
	fmt.Printf("countOfMOst: '%v' \n", countOfMost)
	fmt.Printf("most common: '%v' \n", mostCommon)
}

func processInsertion(template string, rules []string, i int) string {
	newString := []string{}
	var lastCharacter string
	for i := 0; i <= len(template)-2; i++ {
		character1 := string(template[i])
		character2 := string(template[i+1])
		lastCharacter = character2
		insertion, ok := findRule(rules, character1, character2)
		if ok {
			newString = append(newString, character1, insertion)
		} else {
			newString = append(newString, character1)
		}
	}
	newString = append(newString, lastCharacter)

	if i <= 1 {
		return strings.Join(newString, "")
	} else {
		return processInsertion(strings.Join(newString, ""), rules, i-1)
	}

}

func findRule(rules []string, character1 string, character2 string) (string, bool) {

	for _, r := range rules {
		splits := strings.Split(r, " -> ")
		m1 := string(splits[0][0])
		m2 := string(splits[0][1])
		insertion := splits[1]
		if m1 == character1 && m2 == character2 {
			return insertion, true
		}
	}
	return "", false
}
