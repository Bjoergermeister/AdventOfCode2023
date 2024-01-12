package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func getPatternSize(patternDimension float64, divider float64) int {
	if divider <= patternDimension/2 {
		return int(math.Floor(divider))
	} else {
		return int(math.Ceil(patternDimension - divider))
	}
}

func checkHorizontalPattern(lines []string, start int, end int) int {
	patternHeight := float64(end-start) + 1
	for divider := 1.5; divider < patternHeight; divider++ {
		patternSize := getPatternSize(patternHeight, divider)

		matches := true
		for j := 1; j <= patternSize; j++ {
			top := lines[start+int(math.Ceil(divider))-j-1]
			bottom := lines[start+int(math.Floor(divider))+j-1]

			if strings.EqualFold(top, bottom) == false {
				matches = false
				break
			}
		}

		if matches {
			return int(math.Floor(divider)) * 100
		}

	}

	return 0
}

func checkVerticalPattern(lines []string, start int, end int) int {
	patternWidth := float64(len(lines[start]))

	for divider := 1.5; divider < patternWidth; divider++ {
		patternSize := getPatternSize(patternWidth, divider)

		matches := true
		for j := 1; j <= patternSize; j++ {
			if compareVerticalLines(lines, start, end, divider, j) == false {
				matches = false
				break
			}
		}

		if matches {
			return int(math.Floor(divider))
		}
	}

	return 0
}

func compareVerticalLines(lines []string, start int, end int, divider float64, index int) bool {
	for i := start; i <= end; i++ {
		lowerIndex := int(math.Ceil(divider)) - index - 1
		upperIndex := int(math.Floor(divider)) + index - 1
		left := lines[i][lowerIndex]
		right := lines[i][upperIndex]
		if left != right {
			return false
		}
	}
	return true
}

func Puzzle1(lines []string) int {
	total := 0
	lineIndex := 0

	patternStart := 0
	patternEnd := 0
	for lineIndex < len(lines) {
		if len(lines[lineIndex]) == 0 || lineIndex == len(lines)-1 {
			if lineIndex == len(lines)-1 {
				patternEnd = lineIndex
			} else {
				patternEnd = lineIndex - 1
			}
			total += checkHorizontalPattern(lines, patternStart, patternEnd)
			total += checkVerticalPattern(lines, patternStart, patternEnd)
			patternStart = lineIndex + 1
		}

		lineIndex++
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
