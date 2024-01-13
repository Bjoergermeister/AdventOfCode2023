package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) ([]byte, []int) {
	parts := strings.Split(line, " ")

	var pattern []int
	for _, numberAsString := range strings.Split(parts[1], ",") {
		number, _ := strconv.Atoi(numberAsString)
		pattern = append(pattern, number)
	}

	return []byte(parts[0]), pattern
}

func Puzzle1(lines []string) int {
	total := 0
	for _, line := range lines {
		characters, pattern := parseLine(line)

		numberOfPossibleLines := 1
		for i := 0; i < len(characters); i++ {
			if characters[i] == '?' {
				numberOfPossibleLines *= 2
			}
		}

		characters = append(characters, '.')

		validVariants := 0
		for i := 0; i < numberOfPossibleLines; i++ {
			count := 0
			mask := 1
			patternIndex := 0

			matches := true
			for _, character := range characters {
				// If the character is a ?, replace it based on the current value of i
				if character == '?' {
					if i&mask > 0 {
						character = '#'
					} else {
						character = '.'
					}
					mask *= 2
				}

				if character == '#' {
					count++
					continue
				}

				if count > 0 {
					if patternIndex >= len(pattern) || pattern[patternIndex] != count {
						matches = false
						break
					}
					patternIndex++
					count = 0
				}
			}

			if matches && patternIndex == len(pattern) {
				validVariants++
			}
		}
		total += validVariants
	}

	return total
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
}
