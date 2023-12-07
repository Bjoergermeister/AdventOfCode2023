package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

func Puzzle1(lines []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "|")
		winningNumbers := parseStringToIntegerList(parts[0])
		myNumbers := parseStringToIntegerList(parts[1])

		cardResult := 0.5
		for _, winningNumber := range winningNumbers {
			if slices.Contains(myNumbers, winningNumber) {
				cardResult *= 2
			}
		}

		if cardResult >= 1 {
			total += int(cardResult)
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
		lines = append(lines, line[10:])
	}

	fmt.Printf("Puzzle 1: %d\n", Puzzle1(lines))
}
