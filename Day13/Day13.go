package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var isPuzzle2 bool = false

func getPatternSize(patternDimension float64, divider float64) int {
	if divider <= patternDimension/2 {
		return int(math.Floor(divider))
	} else {
		return int(math.Ceil(patternDimension - divider))
	}
}

func calculateLineDifference(line1 string, line2 string) int {
	difference := 0
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			difference++
		}
	}
	return difference
}

func checkHorizontalPattern(lines []string, start int, end int) int {
	patternHeight := float64(end-start) + 1
	for divider := 1.5; divider < patternHeight; divider++ {
		patternSize := getPatternSize(patternHeight, divider)

		difference := 0
		for j := 1; j <= patternSize; j++ {
			top := lines[start+int(math.Ceil(divider))-j-1]
			bottom := lines[start+int(math.Floor(divider))+j-1]

			difference += calculateLineDifference(top, bottom)
		}

		if (difference == 0 && isPuzzle2 == false) || (difference == 1 && isPuzzle2) {
			return int(math.Floor(divider)) * 100
		}

	}

	return 0
}

func checkVerticalPattern(lines []string, start int, end int) int {
	patternWidth := float64(len(lines[start]))

	for divider := 1.5; divider < patternWidth; divider++ {
		patternSize := getPatternSize(patternWidth, divider)

		difference := 0
		for j := 1; j <= patternSize; j++ {
			difference += compareVerticalLines(lines, start, end, divider, j)
		}

		if (difference == 0 && isPuzzle2 == false) || (difference == 1 && isPuzzle2) {
			return int(math.Floor(divider))
		}
	}

	return 0
}

func compareVerticalLines(lines []string, start int, end int, divider float64, index int) int {
	difference := 0
	for i := start; i <= end; i++ {
		lowerIndex := int(math.Ceil(divider)) - index - 1
		upperIndex := int(math.Floor(divider)) + index - 1
		left := lines[i][lowerIndex]
		right := lines[i][upperIndex]
		if left != right {
			difference++
		}
	}
	return difference
}

func run(lines []string) int {
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

	fmt.Printf("Puzzle 1: %d\n", run(lines))
	isPuzzle2 = true
	fmt.Printf("Puzzle 2: %d\n", run(lines))
}
