package main

import (
	"bufio"
	"fmt"
	"os"
)

func Puzzle1(lines []string) int {
	width := len(lines[0])
	height := len(lines)

	total := 0
	for x := 0; x < width; x++ {
		top := 0
		bottom := 0

		for bottom < height {
			if lines[bottom][x] == 'O' {
				total += height - top
				top += 1
			}

			if lines[bottom][x] == '#' {
				top = bottom + 1
			}
			bottom++
		}
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
