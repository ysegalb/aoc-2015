package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsNice(input string) bool {
	vowels := 0
	doubleLetter := false
	badStrings := [][]byte{{'a', 'b'}, {'c', 'd'}, {'p', 'q'}, {'x', 'y'}}
	for i := 0; i < len(input); i++ {
		if input[i] == 'a' || input[i] == 'e' || input[i] == 'i' || input[i] == 'o' || input[i] == 'u' {
			vowels++
		}
		if i > 0 && input[i] == input[i-1] {
			doubleLetter = true
		}
		for _, s := range badStrings {
			if i+1 < len(input) && input[i] == s[0] && input[i+1] == s[1] {
				return false
			}
		}
	}
	return vowels >= 3 && doubleLetter
}

func IsNice2(input string) bool {
	pairExists := false
	repeatWithGap := false

	for i := 0; i < len(input); i++ {
		if i+2 < len(input) && input[i] == input[i+2] {
			repeatWithGap = true
		}
		if i+1 < len(input) {
			pair := input[i : i+2]
			for j := i + 2; j < len(input)-1; j++ {
				if pair == input[j:j+2] {
					pairExists = true
					break
				}
			}
		}
	}

	return pairExists && repeatWithGap
}

func main() {
	nice1 := 0
	nice2 := 0
	for _, line := range readLines("./5/puzzle.txt") {
		if IsNice(line) {
			nice1++
		}
		if IsNice2(line) {
			nice2++
		}
	}
	fmt.Printf("Result: %d nice1 strings and %d nice2 strings", nice1, nice2)
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
