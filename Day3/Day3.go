package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func checkSurrounding(lines []string, lineIndex int, start int, end int) bool {
	xMin := max(start-1, 0)
	yMin := max(lineIndex-1, 0)
	xMax := min(end+1, len(lines[lineIndex])-1)
	yMax := min(lineIndex+1, len(lines)-1)

	for y := yMin; y <= yMax; y++ {
		for x := xMin; x <= xMax; x++ {
			if lines[y][x] != '.' && isDigit(lines[y][x]) == false {
				return true
			}
		}
	}
	return false
}

func Puzzle1(lines []string) int {
	total := 0
	for lineIndex, line := range lines {
		previousPositionIsNumber := false
		numberStart := -1
		numberEnd := -1
		for i := 0; i < len(line); i++ {

			// If the current position is a digit and the previous is not, we found the start of a number
			if isDigit(line[i]) && previousPositionIsNumber == false {
				numberStart = i
				previousPositionIsNumber = true
			}

			// If the current position is not a number, check if the previous was.
			// In that case, we found the end of the number
			if (isDigit(line[i]) == false || i == len(line)-1) && previousPositionIsNumber {
				if isDigit(line[i]) {
					numberEnd = i
				} else {
					numberEnd = i - 1
				}
				previousPositionIsNumber = false
			}

			// We found a complete number
			if numberStart != -1 && numberEnd != -1 {
				if checkSurrounding(lines, lineIndex, numberStart, numberEnd) {
					number, _ := strconv.Atoi(line[numberStart : numberEnd+1])
					total += number
				}
				numberStart = -1
				numberEnd = -1
			}
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
