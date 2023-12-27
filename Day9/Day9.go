package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allZero(list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] != 0 {
			return false
		}
	}

	return true
}

func parseSequence(line string) [][]int {
	sequences := make([][]int, 0, len(line))

	tokens := strings.Split(line, " ")
	var sequence []int
	for _, token := range tokens {
		value, _ := strconv.Atoi(token)
		sequence = append(sequence, value)
	}
	return append(sequences, sequence)
}

func calculateChangeSequences(sequences [][]int) [][]int {
	for {
		lastSequence := len(sequences) - 1
		var changes []int
		for i := 0; i < len(sequences[lastSequence])-1; i++ {
			changes = append(changes, sequences[lastSequence][i+1]-sequences[lastSequence][i])
		}

		sequences = append(sequences, changes)

		if allZero(changes) {
			break
		}
	}

	return sequences
}

func calculateChangeValue(sequences [][]int, findPrevious bool) int {
	lastChangeValue := 0
	for i := len(sequences) - 2; i >= 1; i-- {
		sequence := sequences[i]
		if findPrevious {
			lastChangeValue = sequence[0] - lastChangeValue
		} else {
			lastChangeValue = sequence[len(sequence)-1] + lastChangeValue
		}
	}

	if findPrevious {
		return sequences[0][0] - lastChangeValue
	} else {
		return sequences[0][len(sequences[0])-1] + lastChangeValue
	}
}

func run(lines []string, findPrevious bool) int {
	total := 0
	for _, line := range lines {
		var sequences [][]int

		sequences = parseSequence(line)
		sequences = calculateChangeSequences(sequences)
		total += calculateChangeValue(sequences, findPrevious)
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

	fmt.Printf("Puzzle 1: %d\n", run(lines, false))
	fmt.Printf("Puzzle 2: %d\n", run(lines, true))
}
