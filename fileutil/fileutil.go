package fileutil

import (
	"bufio"
	"errors"
	"log"
	"os"
)

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

