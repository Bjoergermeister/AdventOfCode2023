package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Puzzle1(sequences []string) int {
	total := 0
	for _, sequence := range sequences {
		currentValue := 0
		for _, character := range sequence {
			currentValue += int(character)
			currentValue *= 17
			currentValue %= 256
		}
		total += currentValue
	}
	return total
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	sequences := strings.Split(line, ",")

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(sequences))
}
