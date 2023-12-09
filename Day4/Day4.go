package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"math"
)

func parseStringToIntegerList(line string) []int {
	result := make([]int, 0)

	for _, numberAsString := range strings.Split(line, " ") {
		number, _ := strconv.Atoi(numberAsString)
		if number != 0 {
			result = append(result, number)
		}
	}

	return result
}

func countWinningNumbers(winningNumbers []int, myNumbers[] int) int {
	count := 0
	for _, winningNumber := range winningNumbers {
		if slices.Contains(myNumbers, winningNumber) {
			count++
		}
	}

	return count
}

func getWinningNumberCountForLine(line string) int {
	parts := strings.Split(line, "|")
	winningNumbers := parseStringToIntegerList(parts[0])
	myNumbers := parseStringToIntegerList(parts[1])
	return countWinningNumbers(winningNumbers, myNumbers)
}

func Puzzle1(lines []string) int {
	total := 0.0
	for _, line := range lines {
		winningNumberCount := getWinningNumberCountForLine(line)

		if winningNumberCount > 0 {
			total += math.Pow(2.0, float64(winningNumberCount - 1))
		}
	}

	return int(total)
}

func Puzzle2(lines []string) int {
	total := len(lines)
	cardCounts := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		cardCounts[i] = 1
	}

	for lineIndex, line := range lines {
		winningNumberCount := getWinningNumberCountForLine(line)
		
		if winningNumberCount > 0 {
			for i := lineIndex + 1; i < min(len(lines), lineIndex + 1 + winningNumberCount); i++ {
				cardCounts[i] += cardCounts[lineIndex]
				total += cardCounts[lineIndex]
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
		lines = append(lines, line[strings.Index(line, ":") + 1:])
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
	fmt.Printf("Puzzle 2: %d\n", Puzzle2(lines))
}